#!/usr/bin/env python3
import os
import re
import sys

proto_fields = []
class Method:
	def __init__(self, name):
		result = re.match("# \\(([0-9]+)\\) ([A-Za-z0-9_]+)", name)
		if result == None:
			self.id = None
			self.name = name.replace("# ", "")
		else:
			self.id = int(result[1])
			self.name = result[2]
		self.request = []
		self.response = []

class Type:
	def type_name(self):
		# return a name.
		pass

	def save(self, stream, name):
		# Return a source fragment.
		pass

	def load(self, stream):
		# Return a source fragment.
		pass

class Struct(Type):
	def __init__(self, name, base=None):
		self.fields = []
		self.name = name
		if base != None:
			self.base = delink(base)
			self.fields.append(('Base', self.base, None))
		else:
			self.base = None

	def type_name(self):
		global in_nex
		if not in_nex:
			return "NEX." + self.name
		return self.name

	def safe_name(self):
		return self.name

	def save(self, stream, name):
		return "{}.Struct_{}({})".format(stream, self.safe_name(), name)

	def load(self, stream):
		return "{}.Struct_{}()".format(stream, self.safe_name())

class IntType:
	def __init__(self, bits, signed=False, actual_name=None):
		self.bits = bits
		self.signed = signed

		if actual_name:
			self.actual_name = actual_name
		else:
			self.actual_name = None
	
	def safe_name(self):
		if self.actual_name != None:
			return self.actual_name
		return "{}int{}".format("" if self.signed else "u", self.bits)

	def type_name(self):
		global in_nex
		if self.actual_name != None:
			if not in_nex:
				return "NEX." + self.actual_name
			return self.actual_name
		return "{}int{}".format("" if self.signed else "u", self.bits)

	def save(self, stream, name):
		if self.actual_name:
			real_name = "{}int{}".format("" if self.signed else "u", self.bits) 
			name = real_name + "(" + name + ")"
		return "{}.{}Int{}{}({})".format(stream, "" if self.signed else "U", self.bits, "LE" if self.bits != 8 else "", name)

	def load(self, stream):
		a = "{}.{}Int{}{}()".format(stream, "" if self.signed else "U", self.bits, "LE" if self.bits != 8 else "")
		if self.actual_name:
			a = self.type_name() + "(" + a + ")"
		return a

class StringType:
	def __init__(self, actual_name):
		self.actual_name = actual_name

	def type_name(self):
		global in_nex
		if self.actual_name == "StationURL" and not in_nex:
			return "NEX."+self.actual_name
		return self.actual_name

	def safe_name(self):
		return self.actual_name

	def save(self, stream, name):
		if self.actual_name:
			name = "string(" + name + ")"
		return "{}.String({})".format(stream, name)

	def load(self, stream):
		a = "{}.String()".format(stream)
		if self.actual_name:
			a = self.type_name() + "(" + a + ")"
		return a

class BufferType:
	def __init__(self, len_bits):
		self.len = IntType(len_bits)
		if len_bits == 32:
			self.actual_name = "Buffer"
		else:
			self.actual_name = "QBuffer"

	def safe_name(self):
		return self.actual_name

	def type_name(self):
		global in_nex
		if not in_nex:
			return "NEX."+self.actual_name
		return self.actual_name

	def save(self, stream, name):
		return "{}.{}({})".format(stream, self.actual_name, name)

	def load(self, stream):
		return "{}.{}()".format(stream, self.actual_name)

class BoolType:
	def type_name(self):
		return "bool"

	def safe_name(self):
		return "bool"

	def save(self, stream, name):
		return "{}.Bool({})".format(stream, name)

	def load(self, stream):
		return "{}.Bool()".format(stream)

class FloatType:
	def __init__(self, bits):
		self.bits = bits

	def type_name(self):
		return "float{}".format(self.bits)

	def safe_name(self):
		return self.type_name()

	def save(self, stream, name):
		return "{}.Float{}LE({})".format(stream, self.bits, name)

	def load(self, stream):
		return "{}.Float{}LE()".format(stream, self.bits)

class ListType:
	def __init__(self, inner_type):
		self.inner_type = inner_type
		self.name = "List<{}>".format(inner_type.name)

	def safe_name(self):
		return "List_" + self.inner_type.safe_name()

	def type_name(self):
		return "[]{}".format(self.inner_type.type_name())

	def save(self, stream, name):
		return "{}.{}({})".format(stream, self.safe_name(), name)

	def load(self, stream):
		return "{}.{}()".format(stream, self.safe_name())

class MapType:
	def __init__(self, key_type, value_type):
		self.key_type = key_type
		self.value_type = value_type

	def safe_name(self):
		return "Map_" + self.key_type.safe_name() + "_" + self.value_type.safe_name()

	def type_name(self):
		return "map[{}]{}".format(self.key_type.type_name(), self.value_type.type_name())

	def save(self, stream, name):
		return "{}.{}({})".format(stream, self.safe_name(), name)

	def load(self, stream):
		return "{}.{}()".format(stream, self.safe_name())

class VariantType:
	def type_name(self):
		return "Variant"

	def safe_name(self):
		return "Variant"

	def save(self, stream, name):
		return "{}.Variant({})".format(stream, name)

	def load(self, stream):
		return "{}.Variant()".format(stream)

def extract_name(link):
	return re.match("\\[([A-Za-z0-9 _]+)\\]\\(([A-Za-z0-9#-_]+)\\)", link)[1]

def delink(link):
	return re.sub("\\[([A-Za-z0-9 _]+)\\](?:\\([A-Za-z0-9#-_]+\\))?", lambda a: a[1], link)

def fix_html(s):
	return s.replace("&gt;", ">").replace("&lt;", "<")

struct_funcs = {}
def reg_struct(struct_name, struct_info):
	types[struct_name] = Struct(struct_name)
	types[struct_name].fields = struct_info

types = {}
def reg_type(type_name, typ):
	typ.name = type_name
	types[type_name] = typ

reg_type('Sint8', IntType(8, True))
reg_type('Sint16', IntType(16, True))
reg_type('Sint32', IntType(32, True))
reg_type('Sint64', IntType(64, True))
reg_type('Int32', IntType(32, True))

reg_type('byte', IntType(8))
reg_type('Uint8', IntType(8))
reg_type('Uint16', IntType(16))
reg_type('Uint32', IntType(32))
reg_type('Uint64', IntType(64))

reg_type('PID', IntType(32, actual_name = "PID")) # may differ between 3ds/switch! TODO!
reg_type('Result', IntType(32, actual_name = "Result"))
reg_type('DateTime', IntType(64, actual_name = "DateTime"))

reg_type('Buffer', BufferType(32))
reg_type('qBuffer', BufferType(16))

reg_type('String', StringType("string"))
reg_type('StationURL', StringType("StationURL"))

reg_type('Bool', BoolType())

reg_type('Float', FloatType(32))
reg_type('Double', FloatType(64))

reg_type('Variant', VariantType())

Data_info = (
	('type_name', 'String'),
	('len_plus_four', 'Uint32'),
	('data', 'Buffer')
)
reg_struct('Data', Data_info)
Structure_info = (
)
reg_struct('Structure', Structure_info)

ConnectionData_info = (
	('UrlRegularProtocols', 'String'),
	('LstSpecialProtocols', 'List<byte>'),
	('UrlSpecialProtocols', 'String')
)
reg_struct('RVConnectionData', ConnectionData_info)

class StubType(Type):
	def __init__(self, name):
		self.name = name

	def safe_name(self):
		return "memes"

	def type_name(self):
		print("Stubbed type!! {}".format(repr(self.name)))
		return "memes"

	def save(self, a, b):
		print("stub save ", a, b)
		return ""

	def load(self, a):
		print("stub load ", a)
		return "nil"

list_types = set()
map_pairs = set()

def get_type(type_name):
	if type_name.startswith('List<'):
		if type_name[5:-1] == 'byte': # [hacks intensify]
			list_types.add('Uint8')
		elif type_name[5:10] == 'Data<' or type_name[5:10] == 'Data>':
			list_types.add('Data') # memes?
		else:
			list_types.add(type_name[5:-1])
		return ListType(get_type(type_name[5:-1]))
	elif type_name.startswith('Data<') or type_name == 'Data':
		#return DataHolder(get_type(type_name[5:-1]))
		return types['Data'] # todo specialized data
	elif type_name.startswith('Map<'):
		inner = type_name[4:-1]
		key, value = inner.split(",")
		map_pairs.add((key, value))
		return MapType(get_type(key), get_type(value))
	elif not type_name in types:
		return StubType(type_name)
	else:
		return types[type_name]

def upcase_list(l):
	return list(map(lambda s: s[0].upper() + s[1:], l))

def cleanup_name(name):
	if name.startswith("m_"):
		name = name[2:]
	return "".join(upcase_list(delink(name).split(" "))).replace("(","").replace(")","")

def dedup(pairs):
	fixed_pairs = []

	found_names = set()
	rename_count = {}
	for name, *others in pairs:
		if name in found_names:
			rename_count[name] = rename_count.get(name, 1) + 1
			fixed_pairs.append((name + str(rename_count[name]), *others))
		else:
			found_names.add(name)
			fixed_pairs.append((name, *others))

	return fixed_pairs

structs = set()
def types_pass(f):
	global struct_infos
	# First pass: get type info
	table = False
	skip_table = not is_types
	current_type = None
	types_header_found = is_types # If it's a types page, all of it is types.

	for l in f.readlines():
		l = l.strip()
		if not table and l.startswith('|'):
			if 'This structure inherits from' in l:
				base = delink(re.search("This structure inherits from ([^|]+) \|", l)[1])
				current_type.base = base # lol
				current_type.fields.append(('Base2', base, None))
				table = True
				skip_table = True
				print("We got some weird inheritance!", base)
				#exit()
				continue
			else:
				table = True
				if types_header_found:
					if 'Value' in l:
						print("Got weird table ", l)
						skip_table = True
				continue # Skip the table header..

		if table:
			if l == '': # End of table
				if not skip_table and types_header_found:
					current_type.fields = dedup(current_type.fields)
					types[current_type.name] = current_type
					structs.add(current_type.name)
				skip_table = False
				table = False
			else:
				row = list(map(lambda a: a.strip(), l[1:-1].split('|')))
				if set(row) == set(['---']):
					continue
				if skip_table:
					continue
				if types_header_found:
					field_name_safe = cleanup_name(row[1])
					type_name_safe = fix_html(delink(row[0])).replace(" ", "")

					current_type.fields.append((field_name_safe, type_name_safe))
		else:
			if l == '# Types':
				types_header_found = True
			elif l.startswith("## ") and types_header_found:
				struct_name = re.match("## ([^\\(]+)(?:\\((\S+)\\))?", l)
				current_type = Struct(struct_name[1].strip().replace(" ", ""), struct_name[2])

def methods_pass(f):
	global proto_info
	header = f.readline().strip()
	if header.startswith("## "):
		header = header[3:]
	if header.startswith("[["): # parse link
		end = header.find("]] > ")
		if end == None:
			print("?", name)
			return
		header = header[end+5:]
	# Second (more complicated) pass: get method info.
	# states
	CmdList = 0
	SearchingForMethod = 1
	MethodRequest = 2
	MethodResponse = 3

	cmd_list = []
	method_infos = None
	current_method = None

	state = CmdList
	
	cmd = False
	table = False
	skip_table = False

	for l in f.readlines():
		l = l.strip()
		if not table and l.startswith('|'):
			table = True
			continue # Skip the table header..

		if table:
			if l == '': # End of table
				if not skip_table: # Don't do state transitions if we skip a table!
					if state == CmdList:
						state = SearchingForMethod
					elif state == MethodRequest:
						state = MethodResponse
					elif state == MethodResponse:
						state = SearchingForMethod
						current_method.request = dedup(current_method.request)
						current_method.response = dedup(current_method.response)
						method_infos[current_method.id] = current_method
				table = False
			else: # Table row
				row = list(map(lambda a: a.strip(), l[1:-1].split('|')))
				if set(row) == set(['---']):
					continue

				if skip_table:
					continue

				if state == 0: # the cmd list is the first table
					cmd_list.append(row)
					method_infos = {}
				elif state == MethodRequest:
					name = cleanup_name(row[1])
					if name == '%retval%':
						name = 'returnValue'

					current_method.request.append([name, delink(fix_html(row[0])).replace(" ","")] + row[2:])
				elif state == MethodResponse:
					name = cleanup_name(row[1])
					if name == '%retval%':
						name = 'returnValue'
					current_method.response.append([name, delink(fix_html(row[0])).replace(" ","")] + row[2:])
		else:
			if l.startswith("# "):
				if state == SearchingForMethod:
					meth = Method(l)
					found = False
					for m in cmd_list:
						if delink(m[1]) == meth.name:
							found = True
					if not found:
						#print("Skipping non-method ", l)
						continue
					current_method = meth
					state = MethodRequest
				elif state == MethodResponse:
					# Maybe the method before is just missing info. That's fine.
					current_method.request = dedup(current_method.request)
					current_method.response = dedup(current_method.response)
					method_infos[current_method.id] = current_method

					current_method = Method(l)
					state = MethodRequest
			elif l.startswith("##"):
				if (state == MethodRequest and l != '## Request') or (state == MethodResponse and l != '## Response'):
					#print("Odd!", state, l)
					#	print((state == MethodRequest and l != '## Request'), (state == MethodResponse and l != '## Response'))
					skip_table = True
				else:
					skip_table = False
			elif l.startswith('This method does not take any request data') or l.startswith('This method does not take any parameters'):
				state = MethodResponse
			elif l == 'Same as response for the [Register](#1-register) method.': # Hack!
				current_method.response = method_infos[1].response
			elif l.startswith('This method does not return anything') or l.startswith("This method doesn't return anything"):
				state = SearchingForMethod
				current_method.request = dedup(current_method.request)
				method_infos[current_method.id] = current_method
			elif l.startswith("This method takes no parameters and doesn't return anything."):
				state = SearchingForMethod
				method_infos[current_method.id] = current_method
	if table:
		table = False

	#print("header", header)
	#print(cmd_list)
	proto_info[header] = method_infos

if not os.path.exists("NintendoClients.wiki"):
	print("Please run 'git clone https://github.com/Kinnay/NintendoClients.wiki.git'")
	exit()

if len(sys.argv) == 1:
	print("Usage: {} [output dir]".format(sys.argv[0]))
	exit()

proto_info = {}

blacklist = [
	'RMC-Protocol.md',
	'PRUDP-Protocol.md',
	'PIA-Protocol.md',
	'Mario-Kart-8-Protocol.md',
	'NEX-Common-Types.md',
	'PIA-Types.md',
	'NAT-Traversal-Protocol-(PIA).md'
]

types['Ticket'] = types['Buffer']

a = os.listdir("NintendoClients.wiki")
for name in a:
	if name in blacklist:
		continue
	is_proto =  re.search("Protocol(?:-[^.]+)?.md", name)
	is_types = re.search("Types(?:-[^.]+)?.md", name)

	if is_proto or is_types:
		with open("NintendoClients.wiki/"+name) as f:
			header = f.readline().strip()
			types_pass(f)
"""
print("=====================================================================")

prereq_types = []

def pull_prereqs(l):
	top = []
	for item in l[:]: # Copy the list!
		for f_name, f_type in struct_infos[item]:
			if f_type.startswith("List"):
				f_type = f_type[5:-1]
			if f_type in l:
				l.remove(f_type)
				top.append(f_type)
	return top

all_types = list(struct_infos)

prereq_chunks = [pull_prereqs(all_types)]
remaining = None
while remaining != []:
	remaining = pull_prereqs(prereq_chunks[len(prereq_chunks) - 1])
	if remaining != []:
		prereq_chunks.append(remaining)

for i in range(len(prereq_chunks) - 1, -1, -1):
	prereq_types += prereq_chunks[i]

for prereq_name in prereq_types:
	print(prereq_name)
	#reg_struct(prereq_name, struct_infos[prereq_name])
	pass
print("=====================================================================")

for type_name in struct_infos:
	if not type_name in prereq_types:
		pass
		#reg_struct(type_name, struct_infos[type_name])

print("=====================================================================")
"""

a = os.listdir("NintendoClients.wiki")
for name in a:
	if name in blacklist:
		continue
	is_proto = re.search("Protocol(?:-[^.]+)?.md", name)

	if is_proto:
		with open("NintendoClients.wiki/"+name) as f:
			methods_pass(f)

in_nex = True
out_file = open(sys.argv[1]+"/structs.go", 'w')
out_file.write("// This file is autogenerated.\n// I apologise in advance.\n")
out_file.write("package nex\n\n")
out_file.write("""
type PID uint32
type Result uint32
type Buffer []byte
type QBuffer []byte
type DateTime uint64
type StationURL string
""")

structs.add("RVConnectionData")
structs.add("Data")
structs.add("Structure")

for s in sorted(structs):
	struct = "type {} struct {{\n".format(s)
	for field in types[s].fields:
		field_name = field[0]
		field_type = get_type(field[1])

		struct += "    " + field_name + " " + field_type.type_name() + "\n"
	struct += "}\n\n"
	out_file.write(struct)

in_nex = False

thunks_file = open(sys.argv[1]+"/thunks.go", 'w')
thunks_file.write("package protocols\n")
thunks_file.write("""// This file is autogenerated.
// I apologise in advance.
import (
	NEX "github.com/Stary2001/nex-go"
	)
""")

for p in sorted(proto_info):
	methods = proto_info[p]
	proto_name = "_".join(p.split(" ")[:-1])
	for m in sorted(methods):
		request_params = [field[0] for field in methods[m].request]
		response_params = [field[0] for field in methods[m].response]

		method_text = "func {}_{}_Wrapper(req NEX.RMCRequest) (ret NEX.RMCResponse) {{\n".format(proto_name, methods[m].name)
		if len(request_params) != 0:
			method_text += "    stream := NEX.NewInputStream(req.Parameters)\n"

		for field in methods[m].request:
			field_name = field[0]
			field_type = get_type(field[1])
			method_text += "    {} := {}\n".format(field_name, field_type.load("stream"))

		if len(response_params) != 0:
			method_text += "    rmcResult, {} := {}_{}({})\n".format(",".join(response_params), proto_name, methods[m].name, ",".join(request_params))
		else:
			method_text += "    rmcResult := {}_{}({})\n".format(proto_name, methods[m].name, ",".join(request_params))

		method_text += "    responseStream := NEX.NewOutputStream()\n"
		for field in methods[m].response:
			field_name = field[0]
			field_type = get_type(field[1])
			method_text += "    " + field_type.save("responseStream", field_name) + "\n"


		method_text += """    ret = NEX.NewRMCResponse(int(req.ProtocolID & ^uint8(0x80)), req.CallID)
    if rmcResult == 0x00010001  {
    	ret.SetSuccess(req.MethodID, responseStream.Bytes())
    } else {
    	ret.SetError(rmcResult)
    }
    return
}
"""
		#print(method_text)

		thunks_file.write(method_text)

thunks_file.close()

stubs_file = open(sys.argv[1]+"/stubs.go", 'w')
stubs_file.write("// This file is autogenerated.\n// I apologise in advance.\n")
stubs_file.write("package protocols\n")
stubs_file.write("""import (
	NEX "github.com/Stary2001/nex-go"
	)\n""")

blacklist = ["Authentication (0x0A)", "Secure Connection (0x0B)", "Friends 3DS (0x65)", "Friends Wii U (0x66)"] # Update as new protocols are added.

for p in sorted(proto_info):
	if p in blacklist:
		stubs_file.write("/*")
	methods = proto_info[p]
	proto_name = "_".join(p.split(" ")[:-1])
	for m in sorted(methods):
		request_params = [field[0] + " " + get_type(field[1]).type_name() for field in methods[m].request]
		response_params = [field[0] + " " + get_type(field[1]).type_name() for field in methods[m].response]
		method_text = "func {}_{}({}) (rmcResult uint32, {}) {{\n".format(proto_name, methods[m].name, ", ".join(request_params), ", ".join(response_params))
		method_text += "    rmcResult = 0x80010002\n    return\n"
		method_text += "}\n"
		stubs_file.write(method_text)
	
	if p in blacklist:
		stubs_file.seek(stubs_file.tell() - 1) # This is real dumb
		stubs_file.write("*/\n")
stubs_file.close()

in_nex = True

nex_file = open(sys.argv[1]+"/nex_stream_mixin.go", 'w')
nex_file.write("// This file is autogenerated.\n// I apologise in advance.\n")
nex_file.write("package nex\n")

nex_file.write("import \"fmt\"\n")
nex_file.write("import \"reflect\"\n")

struct_in_method = """func (stream *InputStream) Struct(typeName string) interface{} {
	switch typeName {\n"""

struct_out_method = """func (stream *OutputStream) Struct(out interface{}) {
	switch out.(type) {\n"""

memes = ""

for s in sorted(structs):
	method_in_text = "func (stream *InputStream) Struct_{}() (in {}) {{\n".format(s,s)
	method_out_text = "func (stream *OutputStream) Struct_{}(out {}) {{\n".format(s,s)

	s_obj = types[s]
	for field in s_obj.fields:
		method_in_text += "    in.{} = {}\n".format(field[0], get_type(field[1]).load("stream"))
		method_out_text += "    " + get_type(field[1]).save("stream", "out.{}".format(field[0])) + "\n"

	method_in_text += "    return\n}\n"
	method_out_text += "    return\n}\n"


	memes += method_in_text
	memes += method_out_text

	struct_in_method += "    case \"{}\":\n        return stream.Struct_{}()\n".format(s,s)
	struct_out_method += "    case {}:\n        stream.Struct_{}(out.({}))\n".format(s,s,s)

struct_in_method += """
	default:
		fmt.Println("struct: invalid type", typeName)
		return nil
	}
}
"""
struct_out_method += """
	default:
	fmt.Println("struct: invalid type", reflect.TypeOf(out))
	}
}
"""
nex_file.write(struct_in_method)
nex_file.write(struct_out_method)


for t in sorted(list_types):
	type_name_safe = get_type(t).safe_name()
	real_type_name = get_type(t).type_name()
	method_in_text = "func (stream *InputStream) List_{}() []{} {{".format(type_name_safe, real_type_name)
	method_out_text = "func (stream *OutputStream) List_{}(out []{}) () {{\n".format(type_name_safe, real_type_name)

	method_in_text += """
    list_len := int(stream.UInt32LE())
    list := make([]{}, list_len)
    for i := 0; i < list_len; i++ {{
		list[i] = {}
	}}
	return list
}}\n""".format(real_type_name, get_type(t).load("stream"))

	method_out_text += """
	length := len(out)
    stream.UInt32LE(uint32(length))
    for _, item := range out {{
		{}
	}}
	return
}}\n""".format(get_type(t).save("stream", "item"))

	memes += method_in_text
	memes += method_out_text

for m in sorted(map_pairs):
	type_name_safe = get_type(m[0]).safe_name() + "_" + get_type(m[1]).safe_name()
	real_key_type_name = get_type(m[0]).type_name()
	real_value_type_name = get_type(m[1]).type_name()
	kv = (real_key_type_name, real_value_type_name)

	method_in_text = "func (stream *InputStream) Map_{}() map[{}]{} {{".format(type_name_safe, *kv)
	method_out_text = "func (stream *OutputStream) Map_{}(out map[{}]{}) () {{\n".format(type_name_safe, *kv)

	method_in_text += ("""
    map_len := int(stream.UInt32LE())
    m := make(map[{}]{})
    for i := 0; i < map_len; i++ {{
		key := {}
		value := {}
		m[key] = value
	}}
	return m
}}\n""").format(*kv, get_type(m[0]).load("stream"), get_type(m[1]).load("stream"))

	method_out_text += """
	length := len(out)
    stream.UInt32LE(uint32(length))
    for key, value := range out {{
    	{}
    	{}
	}}
	return
}}\n""".format(get_type(m[0]).save("stream", "key"), get_type(m[1]).save("stream", "value"))

	memes += method_in_text
	memes += method_out_text

nex_file.write(memes)
