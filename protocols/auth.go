package protocols

import (
	NEX "github.com/Stary2001/nex-go"
	"strconv"
	)

func makeTicket(data []byte, secure_key []byte) []byte {
	//secure_key + unk_u32 + struct.pack("I", len(ticket_data))
	stream := NEX.NewOutputStream()
	stream.Write(secure_key)
	stream.UInt32LE(1)
	stream.Buffer(data)
	return stream.Bytes()
}

func Authentication_Login(Client *NEX.Client, username string) (rmcResult uint32, result NEX.Result, pid NEX.PID, ticket []byte, station NEX.RVConnectionData, name string) {
	rmcResult = 0x00010001
	result = 0x00010001

	// TODO: proper username -> pid mapping, proper password retrieval.
	//password := "|+GF-i):/7Z87_:q"
	password := "testing"

	if(username == "guest") {
		pid = 100
		password = "MMQea3n!fsik"
	} else {
		pid_int, _ := strconv.Atoi(username)
		pid = NEX.PID(pid_int)
	}

	key := []byte(password)

	for i := 0; i < 65000+int(pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}
	Kerberos := NEX.NewKerberos(string(key))
	secureKey := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	if(Client.Server.Settings.KerberosKeySize == 32) {
		println("32 byte key")
		secureKey = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	}
	// This just needs to not be empty.
	memes := "\x00"
	ticket = Kerberos.Encrypt(makeTicket([]byte(memes), []byte(secureKey)))

	station.UrlRegularProtocols = "prudps:/address=144.32.68.225;port="+strconv.FormatInt(int64(Client.Server.Settings.SecurePort), 10)+";CID=1;PID=2;sid=1;stream=10;type=2"	

	name = "branch:origin/feature/45925_FixAutoReconnect build:3_10_11_2006_0" // official server
	return
}

func Authentication_LoginEx(Client *NEX.Client, StrUserName string, OExtraData NEX.Data) (rmcResult uint32, returnValue NEX.Result, PidPrincipal NEX.PID, PbufResponse []byte, PConnectionData NEX.RVConnectionData, StrReturnMsg string) {
    rmcResult, returnValue, PidPrincipal, PbufResponse,PConnectionData,StrReturnMsg = Authentication_Login(Client, StrUserName)
    return
}

func Authentication_RequestTicket(Client *NEX.Client, user_pid NEX.PID, server_pid NEX.PID) (rmcResult uint32, result int, ticket []byte) {
	rmcResult = 0x00010001
	result = 0x00010001

	// TODO: proper password lookup
	//password := "|+GF-i):/7Z87_:q"
	password := "testing"
	if(user_pid == 100) {
		password = "MMQea3n!fsik"
	}

	key := []byte(password)

	for i := 0; i < 65000+int(user_pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}

	Kerberos := NEX.NewKerberos(string(key))
	secureKey := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	if(Client.Server.Settings.KerberosKeySize == 32) {
		secureKey = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
		println("32 byte key")
	}
	memes := "\x10\x00\x00\x00\xdds,@\t\xde\x94r$\xa4\xaeB\xad\xf9\xb1\xca,\x00\x00\x00pI\xbd\x8f\xc0\xeb\xb1\x92\xb2]s1\xa9G\xb9\xfe\xe4\x960\xd7\x13\x9co\x97]\xb2E\xfb\x0c`\xae\xab\xeb(\x7f\xb5\xf8\x9aoQ\xa0-\xad\xe5"
	ticket = Kerberos.Encrypt(makeTicket([]byte(memes), []byte(secureKey)))
	return
}

func Authentication_GetPID(Client *NEX.Client, StrUserName string) (rmcResult uint32, returnValue NEX.PID) {
    rmcResult = 0x80010002
    return
}

func Authentication_GetName(Client *NEX.Client, Id NEX.PID) (rmcResult uint32, returnValue string) {
    rmcResult = 0x80010002
    return
}

func Authentication_LoginWithContext(Client *NEX.Client, LoginData NEX.Data) (rmcResult uint32, returnValue NEX.Result, PidPrincipal NEX.PID, PbufResponse NEX.Buffer, PConnectionData NEX.RVConnectionData) {
    rmcResult = 0x80010002
    return
}
