package protocols

import (
	NEX "github.com/Stary2001/nex-go"
)

type NEXMethod func(NEX.RMCRequest) (NEX.RMCResponse)
type NEXProtocol struct
{
	ProtocolID uint8
	Methods map [uint32]NEXMethod
}

func RegisterProtocols() {
	Protocols = make(map[uint8]NEXProtocol)
	Protocols[0xa] = NEXProtocol { 0xa, make(map[uint32]NEXMethod) }
	Protocols[0xb] = NEXProtocol { 0xb, make(map[uint32]NEXMethod) }

	Protocols[0xa].Methods[1] = Authentication_Login_Wrapper
	Protocols[0xa].Methods[3] = Authentication_RequestTicket_Wrapper

	Protocols[0xb].Methods[1] = Secure_Connection_Register_Wrapper
	Protocols[0xb].Methods[4] = Secure_Connection_RegisterEx_Wrapper
}

var Protocols map[uint8]NEXProtocol