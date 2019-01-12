package protocols

import (
	NEX "github.com/Stary2001/nex-go"
	)

func makeTicket(data []byte, secure_key []byte) []byte {
	//secure_key + unk_u32 + struct.pack("I", len(ticket_data))
	stream := NEX.NewOutputStream()
	stream.Write(secure_key)
	stream.UInt32LE(1)
	stream.Buffer(data)
	return stream.Bytes()
}

func Authentication_Login(username string) (rmcResult uint32, result int, pid NEX.PID, ticket []byte, station NEX.RVConnectionData, name string) {
	rmcResult = 0x00010001
	result = 0x00010001
	pid = 1863211397
	password := "|+GF-i):/7Z87_:q"

	key := []byte(password)

	for i := 0; i < 65000+int(pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}
	Kerberos := NEX.NewKerberos(string(key))
	secureKey := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	memes := "\x10\x00\x00\x00\xdds,@\t\xde\x94r$\xa4\xaeB\xad\xf9\xb1\xca,\x00\x00\x00pI\xbd\x8f\xc0\xeb\xb1\x92\xb2]s1\xa9G\xb9\xfe\xe4\x960\xd7\x13\x9co\x97]\xb2E\xfb\x0c`\xae\xab\xeb(\x7f\xb5\xf8\x9aoQ\xa0-\xad\xe5"
	ticket = Kerberos.Encrypt(makeTicket([]byte(memes), []byte(secureKey)))

	station.UrlRegularProtocols = "prudps:/address=192.168.12.1;port=60901;CID=1;PID=2;sid=1;stream=10;type=2"	

	name = "branch:origin/feature/45925_FixAutoReconnect build:3_10_11_2006_0" // official server
	return
}

func Authentication_RequestTicket(user_pid NEX.PID, server_pid NEX.PID) (rmcResult uint32, result int, ticket []byte) {
	rmcResult = 0x00010001
	result = 0x00010001
	password := "|+GF-i):/7Z87_:q"

	key := []byte(password)

	for i := 0; i < 65000+int(user_pid)%1024; i++ {
		key = NEX.MD5Hash(key)
	}

	Kerberos := NEX.NewKerberos(string(key))
	secureKey := "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	memes := "\x10\x00\x00\x00\xdds,@\t\xde\x94r$\xa4\xaeB\xad\xf9\xb1\xca,\x00\x00\x00pI\xbd\x8f\xc0\xeb\xb1\x92\xb2]s1\xa9G\xb9\xfe\xe4\x960\xd7\x13\x9co\x97]\xb2E\xfb\x0c`\xae\xab\xeb(\x7f\xb5\xf8\x9aoQ\xa0-\xad\xe5"
	ticket = Kerberos.Encrypt(makeTicket([]byte(memes), []byte(secureKey)))
	return
}
