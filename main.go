package main

import (
	"fmt"
	NEXProtocols "./protocols"
	NEX "github.com/Stary2001/nex-go"
)

func main() {
	NEXProtocols.RegisterProtocols()

	settings := NEX.NewSettings()

	settings.PrudpVersion = 0
	settings.PrudpV0SignatureVersion = 1
	settings.KerberosKeySize = 16
	settings.AccessKey = "ridfebb9"

	fmt.Println("starting memes")

	Server := NEX.NewServer(settings)

	Server.On("Packet", func(Client *NEX.Client, Packet *NEX.Packet) {
		if Packet.HasFlag(NEX.Flags["NeedAck"]) {
			Server.Acknowledge(Packet)
		}
	})

	Server.On("Data", func(Client *NEX.Client, Packet *NEX.Packet) {
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
		Server.Send(Client, &ResponsePacket)
	})

	Server.Listen(":60000")
}