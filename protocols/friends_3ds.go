package protocols

import (
    NEX "github.com/Stary2001/nex-go"
    )

func Friends_3DS_UpdateProfile(Client *NEX.Client, ProfileData NEX.MyProfile) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdateProfile")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateMii(Client *NEX.Client, Mii NEX.Mii) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdateMii")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateMiiList(Client *NEX.Client, MiiList NEX.MiiList) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdateMiiList")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdatePlayedGames(Client *NEX.Client, PlayedGames []NEX.PlayedGame) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdatePlayedGames")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdatePreference(Client *NEX.Client, Unknown bool, Unknown2 bool, Unknown3 bool) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdatePreference")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendMii(Client *NEX.Client, Friends []NEX.FriendMiiRequest) (rmcResult uint32, Miis []NEX.FriendMii) {
    println("Hit Friends_3DS_GetFriendMii")
    rmcResult = 0x80010002
    return
}
func Friends_3DS_GetFriendMiiList(Client *NEX.Client, Friends []NEX.FriendMiiRequest) (rmcResult uint32, MiiLists []NEX.FriendMiiList) {
    println("Hit Friends_3DS_GetFriendMiiList")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendRelationships(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, FriendRelationships []NEX.FriendRelationship) {
    println("Hit Friends_3DS_GetFriendRelationships")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_AddFriendByPrincipalID(Client *NEX.Client, Unknown uint64, PrincipalId uint32) (rmcResult uint32, FriendRelationship NEX.FriendRelationship) {
    println("Hit Friends_3DS_AddFriendByPrincipalID")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetAllFriends(Client *NEX.Client) (rmcResult uint32, FriendRelationships []NEX.FriendRelationship) {
    println("Hit Friends_3DS_GetAllFriends")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_SyncFriend(Client *NEX.Client, Unknown uint64, Unknown2 []uint32, Unknown3 []uint64) (rmcResult uint32, FriendList []NEX.FriendRelationship) {
    println("Hit Friends_3DS_SyncFriend")
    rmcResult = 0x00010001
    FriendList = make([]NEX.FriendRelationship, 0)
    return
}

func Friends_3DS_UpdatePresence(Client *NEX.Client, PresenceInfo NEX.NintendoPresence, Unknown bool) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdatePresence")
    rmcResult = 0x00010001
    return
}

func Friends_3DS_UpdateFavoriteGameKey(Client *NEX.Client, GameKey NEX.GameKey) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdateFavoriteGameKey")
    rmcResult = 0x00010001
    return
}

func Friends_3DS_UpdateComment(Client *NEX.Client, Comment string) (rmcResult uint32, ) {
    println("Hit Friends_3DS_UpdateComment")
    rmcResult = 0x00010001
    return
}

func Friends_3DS_GetFriendPresence(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, FriendPresenceList []NEX.FriendPresence) {
    println("Hit Friends_3DS_GetFriendPresence")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendPicture(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, FriendPictures []NEX.FriendPicture) {
    println("Hit Friends_3DS_GetFriendPicture")
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendPersistentInfo(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, PersistentInfo []NEX.FriendPersistentInfo) {
    println("Hit Friends_3DS_GetFriendPersistentInfo")
    rmcResult = 0x80010002
    return
}
func Friends_3DS_SendInvitation(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, ) {
    println("Hit Friends_3DS_SendInvitation")
    rmcResult = 0x80010002
    return
}