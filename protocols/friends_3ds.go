package protocols

import (
    NEX "github.com/Stary2001/nex-go"
    )

func Friends_3DS_UpdateProfile(ProfileData NEX.MyProfile) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateMii(Mii NEX.Mii) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateMiiList(MiiList NEX.MiiList) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdatePlayedGames(PlayedGames []NEX.PlayedGame) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdatePreference(Unknown bool, Unknown2 bool, Unknown3 bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendMii(Friends []NEX.FriendMiiRequest) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_3DS_GetFriendMiiList(Friends []NEX.FriendMiiRequest) (rmcResult uint32, MiiLists []NEX.FriendMiiList) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendRelationships(Unknown []uint32) (rmcResult uint32, FriendRelationships []NEX.FriendRelationship) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_AddFriendByPrincipalID(Unknown uint64, PrincipalId uint32) (rmcResult uint32, FriendRelationship NEX.FriendRelationship) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetAllFriends() (rmcResult uint32, FriendRelationships []NEX.FriendRelationship) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_SyncFriend(Unknown uint64, Unknown2 []uint32, Unknown3 []uint64) (rmcResult uint32, FriendList []NEX.FriendRelationship) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdatePresence(PresenceInfo NEX.NintendoPresence, Unknown bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateFavoriteGameKey(GameKey NEX.GameKey) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_UpdateComment(Comment string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendPresence(Unknown []uint32) (rmcResult uint32, FriendPresenceList []NEX.FriendPresence) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendPicture(Unknown []uint32) (rmcResult uint32, FriendPictures []NEX.FriendPicture) {
    rmcResult = 0x80010002
    return
}

func Friends_3DS_GetFriendPersistentInfo(Unknown []uint32) (rmcResult uint32, PersistentInfo []NEX.FriendPersistentInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_3DS_SendInvitation(Unknown []uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}