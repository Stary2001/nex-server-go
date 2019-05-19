package protocols

import (
    NEX "github.com/Stary2001/nex-go"
    )

func Friends_Wii_U_GetAllInformation(Client *NEX.Client, NNAInfo NEX.NNAInfo, NintendoPresence NEX.NintendoPresenceV2, Birthday NEX.DateTime) (rmcResult uint32, PrincipalPreference NEX.PrincipalPreference, StatusMessage NEX.Comment, FriendList []NEX.FriendInfo, SentFriendRequests []NEX.FriendRequest, ReceivedFriendRequests []NEX.FriendRequest, Blacklist []NEX.BlacklistedPrincipal, Unknown bool, Notifications []NEX.PersistentNotification, Unknown2 bool) {
    println("Hit Friends_Wii_U_GetAllInformation")
    rmcResult = 0x00010001
    return
}
func Friends_Wii_U_AddFriend(Client *NEX.Client, Pid NEX.PID) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddFriendByName(Client *NEX.Client, Name string) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_RemoveFriend(Client *NEX.Client, Pid NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddFriendRequest(Client *NEX.Client, Unknown uint32, Unknown2 uint8, Unknown3 string, Unknown4 uint8, Unknown5 string, GameKey NEX.GameKey, Unknown6 NEX.DateTime) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_CancelFriendRequest(Client *NEX.Client, Id uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AcceptFriendRequest(Client *NEX.Client, Id uint64) (rmcResult uint32, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DeleteFriendRequest(Client *NEX.Client, Id uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DenyFriendRequest(Client *NEX.Client, Id uint64) (rmcResult uint32, BlacklistedPrincipal NEX.BlacklistedPrincipal) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_MarkFriendRequestsAsReceived(Client *NEX.Client, FriendRequests []uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddBlackList(Client *NEX.Client, BlacklistedPrincipalIn NEX.BlacklistedPrincipal) (rmcResult uint32, BlacklistedPrincipal NEX.BlacklistedPrincipal) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_RemoveBlackList(Client *NEX.Client, Pid NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdatePresence(Client *NEX.Client, NintendoPresence NEX.NintendoPresenceV2) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdateMii(Client *NEX.Client, Mii NEX.MiiV2) (rmcResult uint32, Unknown NEX.DateTime) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdateComment(Client *NEX.Client, StatusMessage NEX.Comment) (rmcResult uint32, Unknown NEX.DateTime) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdatePreference(Client *NEX.Client, PrincipalPreference NEX.PrincipalPreference) (rmcResult uint32, ) {
    rmcResult = 0x00010001
    return
}
func Friends_Wii_U_GetBasicInfo(Client *NEX.Client, Pids []NEX.PID) (rmcResult uint32, Infos []NEX.PrincipalBasicInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DeleteFriendFlags(Client *NEX.Client, Unknown []NEX.PersistentNotification) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_CheckSettingStatus(Client *NEX.Client) (rmcResult uint32, Unknown uint8) {
    rmcResult = 0x00010001
    Unknown = 1 // it's ok
    return
}
func Friends_Wii_U_GetRequestBlockSettings(Client *NEX.Client, Unknown []uint32) (rmcResult uint32, Settings []NEX.PrincipalRequestBlockSetting) {
    rmcResult = 0x80010002
    return
}