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

	ResponsePacket.SetVersion(Packet.Version)
	ResponsePacket.SetSource(Packet.Destination)
	ResponsePacket.SetDestination(Packet.Source)
	ResponsePacket.SetType(NEX.Types["Data"])

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


	ResponsePacket.Flags |= NEX.Flags["HasSize"] | NEX.Flags["Reliable"] | NEX.Flags["NeedsAck"]
	ResponsePacket.SetPayload(response.Bytes())
	Client.Server.Send(Client, &ResponsePacket)

	Client.Server.Acknowledge(Packet)
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
		if Packet.HasFlag(NEX.Flags["NeedAck"]) && Packet.Type != NEX.Types["Connect"] && Packet.Type != NEX.Types["Data"] {
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

	mk8_settings := NEX.NewSettings()
	mk8_settings.PrudpVersion = 1
	mk8_settings.KerberosKeySize = 32
	mk8_settings.AccessKey = "25dbf96a"

	chat_settings := NEX.NewSettings()
	chat_settings.PrudpVersion = 1
	chat_settings.KerberosKeySize = 32
	chat_settings.AccessKey = "25dbf96a"

	fmt.Println("starting memes")

	AuthServer := createNexServer(settings)
	AuthServer.On("Connect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Acknowledge(Packet)
	})

	AuthServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	go AuthServer.Listen(":60900")

	SecureServer := createNexServer(settings)
	SecureServer.On("Connect", secureServerHandleConnect)
	SecureServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	go SecureServer.Listen(":60901")

	MK8AuthServer := createNexServer(mk8_settings)
	MK8AuthServer.On("Connect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Acknowledge(Packet)
	})

	MK8AuthServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	go MK8AuthServer.Listen(":60800")

	MK8SecureServer := createNexServer(mk8_settings)
	MK8SecureServer.On("Connect", secureServerHandleConnect)
	MK8SecureServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	go MK8SecureServer.Listen(":60801")

	ChatAuthServer := createNexServer(chat_settings)
	ChatAuthServer.On("Connect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Acknowledge(Packet)
	})

	ChatAuthServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	go ChatAuthServer.Listen(":60700")

	ChatSecureServer := createNexServer(chat_settings)
	ChatSecureServer.On("Connect", secureServerHandleConnect)
	ChatSecureServer.On("Disconnect",  func(Client *NEX.Client, Packet *NEX.Packet) {
		Client.Server.Kick(*Client)
	})

	ChatSecureServer.Listen(":60701")
}
