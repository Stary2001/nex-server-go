package main

import (
	"fmt"
	NEXProtocols "./protocols"
	NEX "github.com/Stary2001/nex-go"
)

func nexServerHandleData(Client *NEX.Client, Packet *NEX.Packet) {
	//stream := NEX.NewInputStream(Packet.RMCRequest.Parameters)
	var response NEX.RMCResponse
	//responseStream := NEX.NewOutputStream()

	ResponsePacket := NEX.NewPacket(Client)

	ResponsePacket.SetVersion(0)
	ResponsePacket.SetSource(Packet.Destination)
	ResponsePacket.SetDestination(Packet.Source)
	ResponsePacket.SetType(NEX.Types["Data"])

	pid := 1234567890          // account PID/NEX username (dummy)
	password := "nex_password" // account NEX password (dummy)

	key := []byte(password)

	for i := 0; i < 65000+pid%1024; i++ {
		key = NEX.MD5Hash(key)
	}

	//Kerberos := NEX.NewKerberos(string(key))
	proto_id := Packet.RMCRequest.ProtocolID & ^uint8(0x80)
	proto, ok := NEXProtocols.Protocols[proto_id]

	if ok {
		method, ok := proto.Methods[Packet.RMCRequest.MethodID]
		if ok {
			response = method(Packet.RMCRequest)
		} else {
			// response := failure
			panic(fmt.Sprintf("Unimplemented method %08x in protocol %02x!", Packet.RMCRequest.MethodID, proto_id))
		}
	} else {
		// response := failure
		panic(fmt.Sprintf("Unimplemented protocol id %02x!", proto_id))
	}

	ResponsePacket.SetPayload(response.Bytes())
	Client.Server.Send(Client, &ResponsePacket)
}

func secureServerHandleConnect(Client *NEX.Client, Packet *NEX.Packet) {
	stream := NEX.NewInputStream(Packet.Payload)
	/*ticketData*/ _ = stream.Buffer()
	checkData := stream.Buffer()

	secureKey := ""
	// TODO: use the ticket data to lookup a proper session key. We return 16/32 zeroes atm.
	if Client.Server.Settings.KerberosKeySize == 16 {
		secureKey = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	} else {
		secureKey = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	}

	Client.SetKey(secureKey)

	Kerberos := NEX.NewKerberos(secureKey)
	checkDataDecrypted := Kerberos.Decrypt(checkData)
	checkStream := NEX.NewInputStream(checkDataDecrypted)
	_ = NEX.PID(checkStream.UInt32LE())
	_ = checkStream.UInt32LE()
	check := checkStream.UInt32LE()

	
	outStream := NEX.NewOutputStream()
	outStream.UInt32LE(check + 1)
	outRealStream := NEX.NewOutputStream()
	outRealStream.Buffer(outStream.Bytes())

	Client.Server.AcknowledgeWithPayload(Packet, outRealStream.Bytes())
}

func createNexServer(settings NEX.Settings) *NEX.Server {
	Server := NEX.NewServer(settings)
	Server.On("Packet", func(Client *NEX.Client, Packet *NEX.Packet) {
		if Packet.HasFlag(NEX.Flags["NeedAck"]) && Packet.Type != NEX.Types["Connect"]{
			Server.Acknowledge(Packet)
		}
	})

	Server.On("Data", nexServerHandleData);
	return Server;
}

func main() {
	NEXProtocols.RegisterProtocols()

	settings := NEX.NewSettings()
	settings.PrudpVersion = 0
	settings.PrudpV0SignatureVersion = 1
	settings.KerberosKeySize = 16
	settings.AccessKey = "ridfebb9"

	fmt.Println("starting memes")

	AuthServer := createNexServer(settings)
	AuthServer.On("Connect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Acknowledge(Packet)
	})

	go AuthServer.Listen(":60000")

	SecureServer := createNexServer(settings)
	SecureServer.On("Connect", secureServerHandleConnect)
	SecureServer.Listen(":60001")
}