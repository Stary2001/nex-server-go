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
	Protocols[0xa] = NEXProtocol { 0xa, make(map[uint32]NEXMethod) } // Authentication
	Protocols[0xb] = NEXProtocol { 0xb, make(map[uint32]NEXMethod) } // Secure Connection
	Protocols[0x65] = NEXProtocol { 0x65, make(map[uint32]NEXMethod) } // Friends (3DS)

	Protocols[0xa].Methods[1] = Authentication_Login_Wrapper
	Protocols[0xa].Methods[2] = Authentication_LoginEx_Wrapper
	Protocols[0xa].Methods[3] = Authentication_RequestTicket_Wrapper
	Protocols[0xa].Methods[4] = Authentication_GetPID_Wrapper
	Protocols[0xa].Methods[5] = Authentication_LoginWithContext_Wrapper

	Protocols[0xb].Methods[1] = Secure_Connection_Register_Wrapper
	Protocols[0xb].Methods[2] = Secure_Connection_RequestConnectionData_Wrapper
	Protocols[0xb].Methods[3] = Secure_Connection_RequestUrls_Wrapper
	Protocols[0xb].Methods[4] = Secure_Connection_RegisterEx_Wrapper
	Protocols[0xb].Methods[5] = Secure_Connection_TestConnectivity_Wrapper
	Protocols[0xb].Methods[6] = Secure_Connection_UpdateURLs_Wrapper
	Protocols[0xb].Methods[7] = Secure_Connection_ReplaceURL_Wrapper
	Protocols[0xb].Methods[8] = Secure_Connection_SendReport_Wrapper

	Protocols[0x65].Methods[1] = Friends_3DS_UpdateProfile_Wrapper
	Protocols[0x65].Methods[2] = Friends_3DS_UpdateMii_Wrapper
	Protocols[0x65].Methods[3] = Friends_3DS_UpdateMiiList_Wrapper
	Protocols[0x65].Methods[4] = Friends_3DS_UpdatePlayedGames_Wrapper
	Protocols[0x65].Methods[5] = Friends_3DS_UpdatePreference_Wrapper
	Protocols[0x65].Methods[6] = Friends_3DS_GetFriendMii_Wrapper
	Protocols[0x65].Methods[7] = Friends_3DS_GetFriendMiiList_Wrapper
	//Protocols[0x65].Methods[8] = Friends_3DS_Unknown_Wrapper
	//Protocols[0x65].Methods[9] = Friends_3DS_Unknown_Wrapper
	Protocols[0x65].Methods[10] = Friends_3DS_GetFriendRelationships_Wrapper
	Protocols[0x65].Methods[11] = Friends_3DS_AddFriendByPrincipalID_Wrapper
	//Protocols[0x65].Methods[12] = Friends_3DS_Unknown_Wrapper
	//Protocols[0x65].Methods[13] = Friends_3DS_Unknown_Wrapper
	//Protocols[0x65].Methods[14] = Friends_3DS_Unknown_Wrapper
	Protocols[0x65].Methods[15] = Friends_3DS_GetAllFriends_Wrapper
	//Protocols[0x65].Methods[16] = Friends_3DS_Unknown_Wrapper
	Protocols[0x65].Methods[17] = Friends_3DS_SyncFriend_Wrapper
	Protocols[0x65].Methods[18] = Friends_3DS_UpdatePresence_Wrapper
	Protocols[0x65].Methods[19] = Friends_3DS_UpdateFavoriteGameKey_Wrapper
	Protocols[0x65].Methods[20] = Friends_3DS_UpdateComment_Wrapper
	//Protocols[0x65].Methods[21] = Friends_3DS_Unknown_Wrapper
	Protocols[0x65].Methods[22] = Friends_3DS_GetFriendPresence_Wrapper 
	//Protocols[0x65].Methods[23] = Friends_3DS_Unknown_Wrapper
	Protocols[0x65].Methods[24] = Friends_3DS_GetFriendPicture_Wrapper
	Protocols[0x65].Methods[25] = Friends_3DS_GetFriendPersistentInfo_Wrapper
	Protocols[0x65].Methods[26] = Friends_3DS_SendInvitation_Wrapper
}

var Protocols map[uint8]NEXProtocol