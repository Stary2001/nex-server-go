package protocols

import (
	NEX "github.com/Stary2001/nex-go"
	"encoding/json"
	"encoding/binary"
)

type NEXMethod func(NEX.RMCRequest) (NEX.RMCResponse)
type NEXProtocol struct
{
	ProtocolID uint8
	Methods map [uint32]NEXMethod
}

func Authentication_Login(username string) (rmcResult uint32, result int, pid NEX.PID, ticket []byte, station NEX.RVConnectionData, name string) {
	rmcResult = 0x00010001
	result = 0x00010001
	pid = 1234567890          // account PID/NEX username (dummy)
	password := "nex_password" // account NEX password (dummy)

	key := []byte(password)

	for i := 0; i < 65000+int(pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}
	Kerberos := NEX.NewKerberos(string(key))
	str := "test string test" // dummy data

	ticket = Kerberos.Encrypt([]byte(str))

	JSONBuffer := []byte(`{
		"stream":  "10",
		"type": "2",
		"PID": "2",
		"port": "60001",
		"address": "127.0.0.1",
		"sid": "1",
		"CID": "1"
	}`)

	var JSON map[string]string
	_ = json.Unmarshal(JSONBuffer, &JSON)

	station.UrlRegularProtocols = NEX.NewStationURL("prudps", JSON)

	name = "branch:origin/feature/45925_FixAutoReconnect build:3_10_11_2006_0" // official server
	return
}

func Authentication_RequestTicket(user_pid NEX.PID, server_pid NEX.PID) (rmcResult uint32, result int, ticket []byte) {
	rmcResult = 0x00010001
	result = 0x00010001
	password := "nex_password" // account NEX password (dummy)

	key := []byte(password)

	for i := 0; i < 65000+int(user_pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}
	Kerberos := NEX.NewKerberos(string(key))
	str := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x04" // dummy data
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, uint32(user_pid))
	ticket = Kerberos.Encrypt(append([]byte(str), data...))
	return
}

func RegisterProtocols() {
	Protocols = make(map[uint8]NEXProtocol)
	Protocols[0xa] = NEXProtocol { 0xa, make(map[uint32]NEXMethod) }
	Protocols[0xb] = NEXProtocol { 0xb, make(map[uint32]NEXMethod) }

	Protocols[0xa].Methods[1] = Authentication_Login_Wrapper
	Protocols[0xa].Methods[3] = Authentication_RequestTicket_Wrapper

	//Protocols[0xb].Methods[1] = Secure_Connection_Register_Wrapper
	Protocols[0xb].Methods[4] = Secure_Connection_RegisterEx_Wrapper
}

var Protocols map[uint8]NEXProtocol