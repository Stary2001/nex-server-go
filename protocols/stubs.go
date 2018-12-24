package protocols
import (
	NEX "github.com/Stary2001/nex-go"
	)
func Account_Management_CreateAccount(StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string) (rmcResult uint32, returnValue NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Account_Management_DeleteAccount(IdPrincipal NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_DisableAccount(IdPrincipal NEX.PID, DtUntil NEX.DateTime, StrMessage string) (rmcResult uint32, returnValue NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Account_Management_ChangePassword(StrNewKey string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Account_Management_TestCapability(UiCapability uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetName(IdPrincipal NEX.PID) (rmcResult uint32, StrName string) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetAccountData() (rmcResult uint32, returnValue uint32, OAccountData NEX.AccountData) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetPrivateData() (rmcResult uint32, returnValue bool, OData NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetPublicData(IdPrincipal NEX.PID) (rmcResult uint32, returnValue bool, OData NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetMultiplePublicData(LstPrincipals []NEX.PID) (rmcResult uint32, returnValue bool, OData []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateAccountName(StrName string) (rmcResult uint32, returnValue NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateAccountEmail(StrName string) (rmcResult uint32, returnValue NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateCustomData(OPublicData NEX.Data, OPrivateData NEX.Data) (rmcResult uint32, returnValue NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Account_Management_FindByNameRegex(UiGroups uint32, StrRegex string, ResultRange NEX.ResultRange) (rmcResult uint32, PlstAccounts []NEX.BasicAccountInfo) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateAccountExpiryDate(IdPrincipal NEX.PID, DtExpiry NEX.DateTime, StrExpiredMessage string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateAccountEffectiveDate(IdPrincipal NEX.PID, DtEffectiveFrom NEX.DateTime, StrNotEffectiveMessage string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateStatus(StrStatus string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetStatus(IdPrincipal NEX.PID) (rmcResult uint32, StrStatus string) {
    rmcResult = 0x80010002
    return
}
func Account_Management_GetLastConnectionStats(IdPrincipal NEX.PID) (rmcResult uint32, DtLastSessionLogin NEX.DateTime, DtLastSessionLogout NEX.DateTime, DtCurrentSessionLogin NEX.DateTime) {
    rmcResult = 0x80010002
    return
}
func Account_Management_ResetPassword() (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Account_Management_CreateAccountWithCustomData(StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string, OPublicData NEX.Data, OPrivateData NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_RetrieveAccount() (rmcResult uint32, OAccountData NEX.AccountData, OPublicData NEX.Data, OPrivateData NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Account_Management_UpdateAccount(StrKey string, StrEmail string, OPublicData NEX.Data, OPrivateData NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_ChangePasswordByGuest(StrPrincipalName string, StrEmail string, StrKey string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Account_Management_FindByNameLike(UiGroups uint32, StrLike string, ResultRange NEX.ResultRange) (rmcResult uint32, PlstAccounts []NEX.BasicAccountInfo) {
    rmcResult = 0x80010002
    return
}
func Account_Management_CustomCreateAccount(StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string, OAuthData NEX.Data) (rmcResult uint32, Pid NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Account_Management_NintendoCreateAccount(StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string, OAuthData NEX.Data) (rmcResult uint32, Pid NEX.PID, PidHMAC string) {
    rmcResult = 0x80010002
    return
}
func Account_Management_LookupOrCreateAccount(StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string, OAuthData NEX.Data) (rmcResult uint32, Pid NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Account_Management_DisconnectPrincipal(IdPrincipal NEX.PID) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Account_Management_DisconnectAllPrincipals() (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
/*func Authentication_Login(StrUserName string) (rmcResult uint32, returnValue NEX.Result, PidPrincipal NEX.PID, PbufResponse NEX.Buffer, PConnectionData NEX.RVConnectionData, StrReturnMsg string) {
    rmcResult = 0x80010002
    return
}*/
func Authentication_LoginEx(StrUserName string, OExtraData NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
/*func Authentication_RequestTicket(IdSource NEX.PID, IdTarget NEX.PID) (rmcResult uint32, returnValue NEX.Result, BufResponse NEX.Buffer) {
    rmcResult = 0x80010002
    return
}*/
func Authentication_GetPID(StrUserName string) (rmcResult uint32, returnValue NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Authentication_GetName(Id NEX.PID) (rmcResult uint32, returnValue string) {
    rmcResult = 0x80010002
    return
}
func Authentication_LoginWithContext(LoginData NEX.Data) (rmcResult uint32, returnValue NEX.Result, PidPrincipal NEX.PID, PbufResponse NEX.Buffer, PConnectionData NEX.RVConnectionData) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PrepareGetObjectV1(Param NEX.DataStorePrepareGetParamV1) (rmcResult uint32, PReqGetInfo NEX.DataStoreReqGetInfoV1) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PreparePostObjectV1(Param NEX.DataStorePreparePostParamV1) (rmcResult uint32, PReqPostInfo NEX.DataStoreReqPostInfoV1) {
    rmcResult = 0x80010002
    return
}
func Data_Store_CompletePostObjectV1(Param NEX.DataStoreCompletePostParamV1) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_DeleteObject(Param NEX.DataStoreDeleteParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_DeleteObjects(Params []NEX.DataStoreDeleteParam, Transactional bool) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ChangeMetaV1(Param NEX.DataStoreChangeMetaParamV1) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ChangeMetasV1(DataIds []uint64, Params []NEX.DataStoreChangeMetaParamV1, Transactional bool) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetMeta(Param NEX.DataStoreGetMetaParam) (rmcResult uint32, PMetaInfo NEX.DataStoreMetaInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetMetas(DataIds []uint64, Param NEX.DataStoreGetMetaParam) (rmcResult uint32, PMetaInfo []NEX.DataStoreMetaInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PrepareUpdateObject(Param NEX.DataStorePrepareUpdateParam) (rmcResult uint32, PReqUpdateInfo NEX.DataStoreReqUpdateInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_CompleteUpdateObject(Param NEX.DataStoreCompleteUpdateParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_SearchObject(Param NEX.DataStoreSearchParam) (rmcResult uint32, PSearchResult NEX.DataStoreSearchResult) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetNotificationUrl(Param NEX.DataStoreGetNotificationUrlParam) (rmcResult uint32, Info NEX.DataStoreReqGetNotificationUrlInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetNewArrivedNotificationsV1(Param NEX.DataStoreGetNewArrivedNotificationsParam) (rmcResult uint32, PResult []NEX.DataStoreNotificationV1, PHasNext bool) {
    rmcResult = 0x80010002
    return
}
func Data_Store_RateObject(Target NEX.DataStoreRatingTarget, Param NEX.DataStoreRateObjectParam, FetchRatings bool) (rmcResult uint32, PRating NEX.DataStoreRatingInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetRating(Target NEX.DataStoreRatingTarget, AccessPassword uint64) (rmcResult uint32, PRating NEX.DataStoreRatingInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetRatings(DataIds []uint64, AccessPassword uint64) (rmcResult uint32, PRatings [][]NEX.DataStoreRatingInfoWithSlot, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ResetRating(Target NEX.DataStoreRatingTarget, UpdatePassword uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ResetRatings(DataIds []uint64, Transactional bool) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetSpecificMetaV1(Param NEX.DataStoreGetSpecificMetaParamV1) (rmcResult uint32, PMetaInfos []NEX.DataStoreSpecificMetaInfoV1) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PostMetaBinary(Param NEX.DataStorePreparePostParam) (rmcResult uint32, DataId uint64) {
    rmcResult = 0x80010002
    return
}
func Data_Store_TouchObject(Param NEX.DataStoreTouchObjectParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetRatingWithLog(Target NEX.DataStoreRatingTarget, AccessPassword uint64) (rmcResult uint32, PRating NEX.DataStoreRatingInfo, PRatingLog NEX.DataStoreRatingLog) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PreparePostObject(Param NEX.DataStorePreparePostParam) (rmcResult uint32, PReqPostInfo NEX.DataStoreReqPostInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PrepareGetObject(Param NEX.DataStorePrepareGetParam) (rmcResult uint32, PReqGetInfo NEX.DataStoreReqGetInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_CompletePostObject(Param NEX.DataStoreCompletePostParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetNewArrivedNotifications(Param NEX.DataStoreGetNewArrivedNotificationsParam) (rmcResult uint32, PResult []NEX.DataStoreNotification, PHasNext bool) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetSpecificMeta(Param NEX.DataStoreGetSpecificMetaParam) (rmcResult uint32, PMetaInfos []NEX.DataStoreSpecificMetaInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetPersistenceInfo(OwnerId uint64, PersistenceSlotId uint16) (rmcResult uint32, PPersistenceInfo NEX.DataStorePersistenceInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetPersistenceInfos(OwnerId uint64, PersistenceSlotIds []uint16) (rmcResult uint32, PPersistenceInfo []NEX.DataStorePersistenceInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PerpetuateObject(PersistenceSlotId uint16, DataId uint64, DeleteLastObject bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_UnperpetuateObject(PersistenceSlotId uint16, DeleteLastObject bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PrepareGetObjectOrMetaBinary(Param NEX.DataStorePrepareGetParam) (rmcResult uint32, PReqGetInfo NEX.DataStoreReqGetInfo, PReqGetAdditionalMeta NEX.DataStoreReqGetAdditionalMeta) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetPasswordInfo(DataId uint64) (rmcResult uint32, PPasswordInfo NEX.DataStorePasswordInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetPasswordInfos(DataIds []uint64) (rmcResult uint32, PPasswordInfos []NEX.DataStorePasswordInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetMetasMultipleParam(Params []NEX.DataStoreGetMetaParam) (rmcResult uint32, PMetaInfo []NEX.DataStoreMetaInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_CompletePostObjects(DataIds []uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ChangeMeta(Param NEX.DataStoreChangeMetaParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ChangeMetas(DataIds []uint64, Params []NEX.DataStoreChangeMetaParam, Transactional bool) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_RateObjects(Targets []NEX.DataStoreRatingTarget, Params []NEX.DataStoreRateObjectParam, Transactional bool, FetchRatings bool) (rmcResult uint32, PRatings []NEX.DataStoreRatingInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PostMetaBinaryWithDataId(DataId uint64, Param NEX.DataStorePreparePostParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_PostMetaBinariesWithDataId(DataIds []uint64, Params []NEX.DataStorePreparePostParam, Transactional bool) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_RateObjectWithPosting(Target NEX.DataStoreRatingTarget, RateParam NEX.DataStoreRateObjectParam, PostParam NEX.DataStorePreparePostParam, FetchRatings bool) (rmcResult uint32, PRating NEX.DataStoreRatingInfo) {
    rmcResult = 0x80010002
    return
}
func Data_Store_RateObjectsWithPosting(Targets []NEX.DataStoreRatingTarget, RateParams []NEX.DataStoreRateObjectParam, PostParams []NEX.DataStorePreparePostParam, Transactional bool, FetchRatings bool) (rmcResult uint32, PRatings []NEX.DataStoreRatingInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetObjectInfos(DataIds []uint64) (rmcResult uint32, PInfos []NEX.DataStoreReqGetInfo, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_SearchObjectLight(Param NEX.DataStoreSearchParam) (rmcResult uint32, PSearchResult NEX.DataStoreSearchResult) {
    rmcResult = 0x80010002
    return
}
func Data_Store_AddToBufferQueue(Param NEX.BufferQueueParam, Buffer NEX.QBuffer) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Data_Store_AddToBufferQueues(Params []NEX.BufferQueueParam, Buffers []NEX.QBuffer) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetBufferQueue(Param NEX.BufferQueueParam) (rmcResult uint32, PBufferQueue []NEX.QBuffer) {
    rmcResult = 0x80010002
    return
}
func Data_Store_GetBufferQueues(Params []NEX.BufferQueueParam) (rmcResult uint32, PBufferQueueLst [][]NEX.QBuffer, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_ClearBufferQueues(Params []NEX.BufferQueueParam) (rmcResult uint32, PResults []NEX.Result) {
    rmcResult = 0x80010002
    return
}
func Data_Store_SearchBalloon(Param NEX.DataStoreSearchBalloonParam) (rmcResult uint32, PResults []NEX.DataStoreSearchBalloonResultSet) {
    rmcResult = 0x80010002
    return
}
func Data_Store_FetchMyInfos(Patam NEX.DataStoreFetchMyInfosParam) (rmcResult uint32, PResult NEX.DataStoreFetchMyInfosResult) {
    rmcResult = 0x80010002
    return
}
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
func Friends_Wii_U_GetAllInformation(NNAInfo NEX.NNAInfo, NintendoPresence NEX.NintendoPresenceV2, Birthday NEX.DateTime) (rmcResult uint32, PrincipalPreference NEX.PrincipalPreference, StatusMessage NEX.Comment, FriendList []NEX.FriendInfo, SentFriendRequests []NEX.FriendRequest, ReceivedFriendRequests []NEX.FriendRequest, Blacklist []NEX.BlacklistedPrincipal, Unknown bool, Notifications []NEX.PersistentNotification, Unknown2 bool) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddFriend(Pid NEX.PID) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddFriendByName(Name string) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_RemoveFriend(Pid NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddFriendRequest(Unknown uint32, Unknown2 uint8, Unknown3 string, Unknown4 uint8, Unknown5 string, GameKey NEX.GameKey, Unknown6 NEX.DateTime) (rmcResult uint32, FriendRequest NEX.FriendRequest, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_CancelFriendRequest(Id uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AcceptFriendRequest(Id uint64) (rmcResult uint32, FriendInfo NEX.FriendInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DeleteFriendRequest(Id uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DenyFriendRequest(Id uint64) (rmcResult uint32, BlacklistedPrincipal NEX.BlacklistedPrincipal) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_MarkFriendRequestsAsReceived(FriendRequests []uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_AddBlackList(BlacklistedPrincipal NEX.BlacklistedPrincipal) (rmcResult uint32, BlacklistedPrincipalOut NEX.BlacklistedPrincipal) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_RemoveBlackList(Pid NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdatePresence(NintendoPresence NEX.NintendoPresenceV2) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdateMii(Mii NEX.MiiV2) (rmcResult uint32, Unknown NEX.DateTime) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdateComment(StatusMessage NEX.Comment) (rmcResult uint32, Unknown NEX.DateTime) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_UpdatePreference(PrincipalPreferenc NEX.PrincipalPreference) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_GetBasicInfo(Pids []NEX.PID) (rmcResult uint32, Infos []NEX.PrincipalBasicInfo) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_DeleteFriendFlags(Unknown []NEX.PersistentNotification) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_CheckSettingStatus() (rmcResult uint32, Unknown uint8) {
    rmcResult = 0x80010002
    return
}
func Friends_Wii_U_GetRequestBlockSettings(Unknown []uint32) (rmcResult uint32, Settings []NEX.PrincipalRequestBlockSetting) {
    rmcResult = 0x80010002
    return
}
func Friends_AddFriend(UiPlayer uint32, UiDetails uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_AddFriendByName(StrPlayerName string, UiDetails uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_AddFriendWithDetails(UiPlayer uint32, UiDetails uint32, StrMessage string) (rmcResult uint32, RelationshipData NEX.RelationshipData) {
    rmcResult = 0x80010002
    return
}
func Friends_AddFriendByNameWithDetails(UiPlayer uint32, UiDetails uint32, StrMessage string) (rmcResult uint32, RelationshipData NEX.RelationshipData) {
    rmcResult = 0x80010002
    return
}
func Friends_AcceptFriendship(UiPlayer uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_DeclineFriendship(UiPlayer uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_BlackList(UiPlayer uint32, UiDetails uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_BlackListByName(StrPlayerName string, UiDetails uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_ClearRelationship(UiPlayer uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_UpdateDetails(UiPlayer uint32, UiDetails uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Friends_GetList(ByRelationship uint8, BReversed bool) (rmcResult uint32, LstFriendsList []uint32) {
    rmcResult = 0x80010002
    return
}
func Friends_GetDetailedList(ByRelationship uint8, BReversed bool) (rmcResult uint32, LstFriendsList []NEX.FriendData) {
    rmcResult = 0x80010002
    return
}
func Friends_GetRelationships(ResultRange NEX.ResultRange) (rmcResult uint32, UiTotalCount uint32, LstRelationshipsList []NEX.RelationshipData) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_EndParticipation(IdGathering uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_GetParticipants(IdGathering uint32, BOnlyActive bool) (rmcResult uint32, LstParticipants []NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_GetDetailedParticipants(IdGathering uint32, BOnlyActiv bool) (rmcResult uint32, LstParticipants []NEX.ParticipantDetails) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_GetParticipantsURLs(LstGatherings []uint32) (rmcResult uint32, LstGatheringURLs []NEX.GatheringURLs) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_GetGatheringRelations(Id uint32, Descr string) (rmcResult uint32, returnValue string) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Ext_DeleteFromDeletions(LstDeletions []uint32, Pid NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Match_Making_RegisterGathering(AnyGathering NEX.Data) (rmcResult uint32, returnValue uint32) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UnregisterGathering(IdGathering uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UnregisterGatherings(LstGatherings []uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UpdateGathering(AnyGathering NEX.Data) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Invite(IdGathering uint32, LstPrincipals []NEX.PID, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_AcceptInvitation(IdGathering uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_DeclineInvitation(IdGathering uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_CancelInvitation(IdGathering uint32, LstPrincipals []NEX.PID, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetInvitationsSent(IdGathering uint32) (rmcResult uint32, LstInvitations []NEX.Invitation) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetInvitationsReceived() (rmcResult uint32, LstInvitations []NEX.Invitation) {
    rmcResult = 0x80010002
    return
}
func Match_Making_Participate(IdGathering uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_CancelParticipation(IdGathering uint32, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetParticipants(IdGathering uint32) (rmcResult uint32, LstParticipants []NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Match_Making_AddParticipants(IdGathering uint32, LstParticipants []NEX.PID, StrMessage string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetDetailedParticipants(IdGathering uint32) (rmcResult uint32, LstParticipants []NEX.ParticipantDetails) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetParticipantsURLs(IdGathering uint32) (rmcResult uint32, LstStationURL []NEX.StationURL) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByType(StrType string, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByDescription(StrDescription string, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByDescriptionRegex(StrDescriptionRegex string, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByID(LstID []uint32) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindBySingleID(Id uint32) (rmcResult uint32, BResult bool, PGathering NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByOwner(Id NEX.PID, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByParticipants(Pid []NEX.PID) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindInvitations(ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindBySQLQuery(StrQuery string, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_LaunchSession(IdGathering uint32, StrURL string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UpdateSessionURL(IdGathering uint32, StrURL string) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetSessionURL(IdGathering uint32) (rmcResult uint32, returnValue bool, StrURL string) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetState(IdGathering uint32) (rmcResult uint32, returnValue bool, UiState uint32) {
    rmcResult = 0x80010002
    return
}
func Match_Making_SetState(IdGathering uint32, UiNewState uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_ReportStats(IdGathering uint32, LstStats []NEX.GatheringStats) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetStats(IdGathering uint32, LstParticipants []uint32, LstColumns []uint8) (rmcResult uint32, returnValue bool, PlstStats []NEX.GatheringStats) {
    rmcResult = 0x80010002
    return
}
func Match_Making_DeleteGathering(Gid uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetPendingDeletions(UiReason uint32, ResultRange NEX.ResultRange) (rmcResult uint32, returnValue bool, LstDeletions []NEX.DeletionEntry) {
    rmcResult = 0x80010002
    return
}
func Match_Making_DeleteFromDeletions(LstDeletions []uint32) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_MigrateGatheringOwnershipV1(Gid uint32, LstPotentialNewOwnersID []NEX.PID) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_FindByDescriptionLike(StrDescriptionLike string, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Match_Making_RegisterLocalURL(Gid uint32, Url NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Match_Making_RegisterLocalURLs(Gid uint32, LstUrls []NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UpdateSessionHostV1(Gid uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Match_Making_GetSessionURLs(Gid uint32) (rmcResult uint32, LstURLs []NEX.StationURL) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UpdateSessionHost(Gid uint32, IsMigrateOwner bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Match_Making_UpdateGatheringOwnership(Gid uint32, ParticipantsOnly bool) (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Match_Making_MigrateGatheringOwnership(Gid uint32, LstPotentialNewOwnersID []NEX.PID, ParticipantsOnly bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_CloseParticipation(Gid uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_OpenParticipation(Gid uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_AutoMatchmake_Postpone(AnyGathering NEX.Data, StrMessage string) (rmcResult uint32, JoinedGathering NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSession(SearchCriteria NEX.MatchmakeSessionSearchCriteria, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSessionWithHostUrls(SearchCriteria NEX.MatchmakeSessionSearchCriteria, ResultRange NEX.ResultRange) (rmcResult uint32, LstGathering []NEX.Data, LstGatheringUrls []NEX.GatheringURLs) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_CreateMatchmakeSession(AnyGathering NEX.Data, StrMessage string, ParticipationCount uint16) (rmcResult uint32, Gid uint32, SessionKey NEX.Buffer) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_JoinMatchmakeSession(Gid uint32, StrMessage string) (rmcResult uint32, SessionKey NEX.Buffer) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_ModifyCurrentGameAttribute(Gid uint32, AttribIndex uint32, NewValue uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateNotificationData(UiType uint32, UiParam1 uint64, UiParam2 uint64, StrParam string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetFriendNotificationData(UiType int32) (rmcResult uint32, DataList []NEX.NotificationEvent) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateApplicationBuffer(Gid uint32, ApplicationBuffer NEX.Buffer) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateMatchmakeSessionAttribute(Gid uint32, Attribs []uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetlstFriendNotificationData(LstTypes []uint32) (rmcResult uint32, DataList []NEX.NotificationEvent) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateMatchmakeSession(AnyGathering NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_AutoMatchmakeWithSearchCriteria_Postpone(LstSearchCriteria []NEX.MatchmakeSessionSearchCriteria, AnyGathering NEX.Data, StrMessage string) (rmcResult uint32, JoinedGathering NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetPlayingSession(LstPid []NEX.PID) (rmcResult uint32, LstPlayingSession []NEX.PlayingSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_CreateCommunity(Community NEX.PersistentGathering, StrMessage string) (rmcResult uint32, Gid uint32) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateCommunity(Community NEX.PersistentGathering) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_JoinCommunity(Gid uint32, StrMessage string, StrPassword string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindCommunityByGatheringId(LstGid []uint32) (rmcResult uint32, LstCommunity []NEX.PersistentGathering) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindOfficialCommunity(IsAvailableOnly bool, ResultRange NEX.ResultRange) (rmcResult uint32, LstCommunity []NEX.PersistentGathering) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindCommunityByParticipant(Pid NEX.PID, ResultRange NEX.ResultRange) (rmcResult uint32, LstCommunity []NEX.PersistentGathering) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdatePrivacySetting(OnlineStatus bool, ParticipationCommunity bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetMyBlockList() (rmcResult uint32, LstPrincipalId []NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_AddToBlockList(LstPrincipalId []NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_RemoveFromBlockList(LstPrincipalId []NEX.PID) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_ClearMyBlockList() (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_ReportViolation(Pid NEX.PID, UserName string, ViolationCode uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_IsViolationUser() (rmcResult uint32, Flag bool, Score uint32) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_JoinMatchmakeSessionEx(Gid uint32, StrMessage string, DontCareMyBlockList bool, ParticipationCount uint16) (rmcResult uint32, SessionKey NEX.Buffer) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetSimplePlayingSession(LstPrincipalId []NEX.PID, IncludeLoginUser bool) (rmcResult uint32, LstSimplePlayingSession []NEX.SimplePlayingSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GetSimpleCommunity(GatheringIdList []uint32) (rmcResult uint32, LstSimpleCommunityList []NEX.SimpleCommunity) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_AutoMatchmakeWithGatheringId_Postpone(LstGid []uint32, AnyGathering NEX.Data, StrMessage string) (rmcResult uint32, JoinedGathering NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateProgressScore(Gid uint32, ProgressScore uint8) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_DebugNotifyEvent(Pid NEX.PID, MainType uint32, SubType uint32, Param1 uint64, Param2 uint64, StringParam string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_GenerateMatchmakeSessionSystemPassword(Gid uint32) (rmcResult uint32, Password string) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_ClearMatchmakeSessionSystemPassword(Gid uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_CreateMatchmakeSessionWithParam(CreateMatchmakeSessionParam NEX.CreateMatchmakeSessionParam) (rmcResult uint32, JoinedMatchmakeSession NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_JoinMatchmakeSessionWithParam(JoinMatchmakeSessionParam NEX.JoinMatchmakeSessionParam) (rmcResult uint32, JoinedMatchmakeSession NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_AutoMatchmakeWithParam_Postpone(AutoMatchmakeParam NEX.AutoMatchmakeParam) (rmcResult uint32, JoinedMatchmakeSession NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindMatchmakeSessionByGatheringIdDetail(Gid uint32) (rmcResult uint32, MatchmakeSession NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSessionNoHolder(SearchCriteria NEX.MatchmakeSessionSearchCriteria, ResultRange NEX.ResultRange) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSessionWithHostUrlsNoHolder(SearchCriteria NEX.MatchmakeSessionSearchCriteria, ResultRange NEX.ResultRange) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession, LstGatheringUrls []NEX.GatheringURLs) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_UpdateMatchmakeSessionPart(UpdateMatchmakeSessionParam NEX.UpdateMatchmakeSessionParam) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_RequestMatchmaking(AutoMatchmakeParam NEX.AutoMatchmakeParam) (rmcResult uint32, RequestId uint64) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_WithdrawMatchmaking(RequestId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_WithdrawMatchmakingAll() (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindMatchmakeSessionByGatheringId(LstGid []uint32) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindMatchmakeSessionBySingleGatheringId(Gid uint32) (rmcResult uint32, MatchmakeSession NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindMatchmakeSessionByOwner(Id NEX.PID, ResultRange NEX.ResultRange) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_FindMatchmakeSessionByParticipant(Param NEX.FindMatchmakeSessionByParticipantParam) (rmcResult uint32, LstSession []NEX.FindMatchmakeSessionByParticipantResult) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSessionNoHolderNoResultRange(SearchCriteria NEX.MatchmakeSessionSearchCriteria) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession) {
    rmcResult = 0x80010002
    return
}
func Matchmake_Extension_BrowseMatchmakeSessionWithHostUrlsNoHolderNoResultRange(SearchCriteria NEX.MatchmakeSessionSearchCriteria) (rmcResult uint32, LstMatchmakeSession []NEX.MatchmakeSession, LstGatheringUrls []NEX.GatheringURLs) {
    rmcResult = 0x80010002
    return
}
func Message_Delivery_DeliverMessage(OUserMessage NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Messaging_DeliverMessage(OUserMessage NEX.Data) (rmcResult uint32, OModifiedMessage NEX.Data, LstSandboxNodeId []uint32, LstParticipants []NEX.PID) {
    rmcResult = 0x80010002
    return
}
func Messaging_GetNumberOfMessages(Recipient NEX.MessageRecipient) (rmcResult uint32, UiNbMessages uint32) {
    rmcResult = 0x80010002
    return
}
func Messaging_GetMessagesHeaders(Recipient NEX.MessageRecipient, Range NEX.ResultRange) (rmcResult uint32, LstMsgHeaders []NEX.UserMessage) {
    rmcResult = 0x80010002
    return
}
func Messaging_RetrieveAllMessagesWithinRange(Recipient NEX.MessageRecipient, Range NEX.ResultRange) (rmcResult uint32, LstMessages []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Messaging_RetrieveMessages(Recipient NEX.MessageRecipient, LstMsgIDs []uint32, BLeaveOnServer bool) (rmcResult uint32, LstMessages []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Messaging_DeleteMessages(Recipient NEX.MessageRecipient, LstMessagesToDelete []uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Messaging_DeleteAllMessages(Recipient NEX.MessageRecipient) (rmcResult uint32, UiNbDeletedMessages uint32) {
    rmcResult = 0x80010002
    return
}
func Monitoring_PingDaemon() (rmcResult uint32, returnValue bool) {
    rmcResult = 0x80010002
    return
}
func Monitoring_GetClusterMembers() (rmcResult uint32, StrValues []string) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_RequestProbeInitiation(UrlTargetList []NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_InitiateProbe(UrlStationToProbe NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_RequestProbeInitiationExt(UrlTargetList []NEX.StationURL, UrlStationToProbe NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_ReportNATTraversalResult(Cid uint32, Result bool, Rtt uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_ReportNATProperties(Natmapping uint32, Natfiltering uint32, Rtt uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_GetRelaySignatureKey() (rmcResult uint32, RelayMode int32, CurrentUTCTime NEX.DateTime, Address string, Port uint16, RelayAddressType int32, GameServerID uint32) {
    rmcResult = 0x80010002
    return
}
func NAT_Traversal_ReportNATTraversalResultDetail(Cid uint32, Result bool, Detail int32, Rtt uint32) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Nintendo_Notifications_ProcessNintendoNotificationEvent(EventObject NEX.NintendoNotificationEvent) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Notifications_ProcessNotificationEvent(OEvent NEX.NotificationEvent) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_FindByGroup(UiGroup uint32) (rmcResult uint32, LstTags []string) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_InsertItem(UiGroup uint32, StrTag string, BufData NEX.Buffer, BReplace bool) (rmcResult uint32, BResult bool) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_RemoveItem(UiGroup uint32, StrTag string) (rmcResult uint32, BResult bool) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_GetItem(UiGroup uint32, StrTag string) (rmcResult uint32, BufData NEX.Buffer, BResult bool) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_InsertCustomItem(UiGroup uint32, StrTag string, HData NEX.Data, BReplace bool) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_GetCustomItem(UiGroup uint32, StrTag string) (rmcResult uint32, HData NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Persistent_Store_FindItemsBySQLQuery(UiGroup uint32, StrTag string, StrQuery string) (rmcResult uint32, LstData []NEX.Data) {
    rmcResult = 0x80010002
    return
}
func Ranking_UploadScore(ScoreData NEX.RankingScoreData, UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_DeleteScore(Category uint32, UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_DeleteAllScores(UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_UploadCommonData(CommonData NEX.Buffer, UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_DeleteCommonData(UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetCommonData(UniqueId uint64) (rmcResult uint32, CommonData NEX.Buffer) {
    rmcResult = 0x80010002
    return
}
func Ranking_ChangeAttributes(Category uint32, ChangeParam NEX.RankingChangeAttributesParam, UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_ChangeAllAttributes(ChangeParam NEX.RankingChangeAttributesParam, UniqueId uint64) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetRanking(RankingMode uint8, Category uint32, OrderParam NEX.RankingOrderParam, UniqueId uint64, PrincipalId NEX.PID) (rmcResult uint32, PResult NEX.RankingResult) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetApproxOrder(Category uint32, OrderParam NEX.RankingOrderParam, Score uint32, UniqueId uint64, PrincipalId uint32) (rmcResult uint32, POrder uint32) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetStats(Category uint32, OrderParam NEX.RankingOrderParam, Flags uint32) (rmcResult uint32, PStats NEX.RankingStats) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetRankingByPIDList(PrincipalIdList []NEX.PID, RankingMode uint8, Category uint32, OrderParam NEX.RankingOrderParam, UniqueId uint64) (rmcResult uint32, PResult NEX.RankingResult) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetRankingByUniqueIdList(NexUniqueIdList []uint64, RankingMode uint8, Category uint32, OrderParam NEX.RankingOrderParam, UniqueId uint64) (rmcResult uint32, PResult NEX.RankingResult) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetCachedTopXRanking(Category uint32, OrderParam NEX.RankingOrderParam) (rmcResult uint32, PResult NEX.RankingCachedResult) {
    rmcResult = 0x80010002
    return
}
func Ranking_GetCachedTopXRankings(Categories []uint32, OrderParams []NEX.RankingOrderParam) (rmcResult uint32, PResults []NEX.RankingCachedResult) {
    rmcResult = 0x80010002
    return
}
func Remote_Log_Device_Log(Message string) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_Register(VecMyURLs []NEX.StationURL) (rmcResult uint32, returnValue NEX.Result, PidConnectionID uint32, UrlPublic NEX.StationURL) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_RequestConnectionData(CidTarget uint32, PidTarget uint32) (rmcResult uint32, returnValue bool, PvecConnectionsData []NEX.ConnectionData) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_RequestUrls(CidTarget uint32, PidTarget uint32) (rmcResult uint32, returnValue bool, PlstURLs []NEX.StationURL) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_RegisterEx(VecMyURLs []NEX.StationURL, HCustomData NEX.Data) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_TestConnectivity() (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_UpdateURLs(VecMyURLs []NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_ReplaceURL(Target NEX.StationURL, Url NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Secure_Connection_SendReport(ReportId uint32, ReportData NEX.QBuffer) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Simple_Authentication_LoginWithTokenEx(StrToken string, PConnectionData NEX.RVConnectionData, OAnyData NEX.Data) (rmcResult uint32, returnValue NEX.Result, PidPrincipal uint32, PConnectionDataOut NEX.RVConnectionData, StrReturnMsg string) {
    rmcResult = 0x80010002
    return
}
func Utility_AcquireNexUniqueId() (rmcResult uint32, PNexUniqueId uint64) {
    rmcResult = 0x80010002
    return
}
func Utility_AcquireNexUniqueIdWithPassword() (rmcResult uint32, PNexUniqueIdInfo NEX.UniqueIdInfo) {
    rmcResult = 0x80010002
    return
}
func Utility_AssociateNexUniqueIdWithMyPrincipalId(UniqueIdInfo NEX.UniqueIdInfo) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Utility_AssociateNexUniqueIdsWithMyPrincipalId(UniqueIdInfo []NEX.UniqueIdInfo) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
func Utility_GetAssociatedNexUniqueIdWithMyPrincipalId() (rmcResult uint32, PUniqueIdInfo NEX.UniqueIdInfo) {
    rmcResult = 0x80010002
    return
}
func Utility_GetAssociatedNexUniqueIdsWithMyPrincipalId() (rmcResult uint32, PUniqueIdInfo []NEX.UniqueIdInfo) {
    rmcResult = 0x80010002
    return
}
func Utility_GetIntegerSettings(IntegerSettingIndex uint32) (rmcResult uint32, IntegerSettings map[uint16]int32) {
    rmcResult = 0x80010002
    return
}
func Utility_GetStringSettings(StringSettingIndex uint32) (rmcResult uint32, StringSettings map[uint16]string) {
    rmcResult = 0x80010002
    return
}
