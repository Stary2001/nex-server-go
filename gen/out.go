type BasicAccountInfo struct {
    Base Structure
    PidOwner uint32
    StrName String
}

type AccountData struct {
    Base Structure
    Pid uint32
    StrName String
    UiGroups uint32
    StrEmail String
    DtCreationDate uint64
    DtEffectiveDate uint64
    StrNotEffectiveMsg String
    DtExpiryDate uint64
    StrExpiredMsg String
}

type DataStorePrepareGetParamV1 struct {
    Base Structure
    DataId uint32
    LockId uint32
}

type DataStorePrepareGetParam struct {
    Base Structure
    DataId uint64
    LockId uint32
    PersistenceTarget DataStorePersistenceTarget
    AccessPassword uint64
    ExtraData []String
}

type DataStoreReqGetInfoV1 struct {
    Base Structure
    Url String
    RequestHeaders []DataStoreKeyValue
    Size uint32
    RootCaCert bytes
}

type DataStoreReqGetInfo struct {
    Base Structure
    Url String
    RequestHeaders []DataStoreKeyValue
    Size uint32
    RootCaCert bytes
    DataId uint64
}

type DataStoreReqGetAdditionalMeta struct {
    Base Structure
    OwnerId uint64
    DataType uint16
    Version uint16
    MetaBinary bytes
}

type DataStorePreparePostParamV1 struct {
    Base Structure
    Size uint32
    Name String
    DataType uint16
    MetaBinary bytes
    Permission DataStorePermission
    DelPermission DataStorePermission
    Flag uint32
    Period uint16
    ReferDataId uint32
    Tags []String
    RatingInitParams []DataStoreRatingInitParamWithSlot
}

type DataStorePreparePostParam struct {
    Base Structure
    Size uint32
    Name String
    DataType uint16
    MetaBinary bytes
    Permission DataStorePermission
    DelPermission DataStorePermission
    Flag uint32
    Period uint16
    ReferDataId uint32
    Tags []String
    RatingInitParams []DataStoreRatingInitParamWithSlot
    PersistenceInitParam DataStorePersistenceInitParam
    ExtraData []String
}

type DataStoreReqPostInfoV1 struct {
    Base Structure
    DataId uint32
    Url String
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert bytes
}

type DataStoreReqPostInfo struct {
    Base Structure
    DataId uint64
    Url String
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert bytes
}

type DataStoreCompletePostParamV1 struct {
    Base Structure
    DataId uint32
    IsSuccess bool
}

type DataStoreCompletePostParam struct {
    Base Structure
    DataId uint64
    IsSuccess bool
}

type DataStoreDeleteParam struct {
    Base Structure
    DataId uint64
    UpdatePassword uint64
}

type DataStoreChangeMetaParamV1 struct {
    Base Structure
    DataId uint64
    ModifiesFlag uint32
    Name String
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary bytes
    Tags []String
    UpdatePassword uint64
}

type DataStoreChangeMetaParam struct {
    Base Structure
    DataId uint64
    ModifiesFlag uint32
    Name String
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary bytes
    Tags []String
    UpdatePassword uint64
    ReferredCnt uint32
    DataType uint16
    Status uint8
    CompareParam DataStoreChangeMetaCompareParam
    PersistenceTarget DataStorePersistenceTarget
}

type DataStoreGetMetaParam struct {
    Base Structure
    DataId uint64
    PersistenceTarget DataStorePersistenceTarget
    ResultOption uint8
    AccessPassword uint64
}

type DataStoreRatingInfo struct {
    Base Structure
    TotalValue uint64
    Count uint32
    InitialValue uint64
}

type DataStoreRatingInfoWithSlot struct {
    Base Structure
    Slot uint8
    Rating DataStoreRatingInfo
}

type DataStoreMetaInfo struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    Name String
    DataType uint16
    MetaBinary bytes
    Permission DataStorePermission
    DelPermission DataStorePermission
    CreatedTime uint64
    UpdatedTime uint64
    Period uint16
    Status uint8
    ReferredCnt uint32
    ReferDataId uint32
    Flag uint32
    ReferredTime uint64
    ExpireTime uint64
    Tags []String
    Ratings []DataStoreRatingInfoWithSlot
}

type DataStorePrepareUpdateParam struct {
    Base Structure
    DataId uint64
    Size uint32
    UpdatePassword uint64
    ExtraData []String
}

type DataStoreReqUpdateInfo struct {
    Base Structure
    Version uint32
    Url String
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert bytes
}

type DataStoreCompleteUpdateParam struct {
    Base Structure
    DataId uint64
    Version uint32
    IsSuccess bool
}

type DataStoreSearchParam struct {
    Base Structure
    SearchTarget uint8
    OwnerIds []uint64
    OwnerType uint8
    DestinationIds []uint64
    DataType uint16
    CreatedAfter uint64
    CreatedBefore uint64
    UpdatedAfter uint64
    UpdatedBefore uint64
    ReferDataId uint32
    Tags []String
    ResultOrderColumn uint8
    ResultOrder uint8
    ResultRange ResultRange
    ResultOption uint8
    MinimalRatingFrequency uint32
    UseCache bool
    TotalCountEnabled bool
    DataTypes []uint16
}

type DataStoreSearchResult struct {
    Base Structure
    TotalCount uint32
    Result []DataStoreMetaInfo
    TotalCountType uint8
}

type DataStoreGetNotificationUrlParam struct {
    Base Structure
    PreviousUrl String
}

type DataStoreReqGetNotificationUrlInfo struct {
    Base Structure
    Url String
    Key String
    Query String
    RootCaCert bytes
}

type DataStoreGetNewArrivedNotificationsParam struct {
    Base Structure
    LastNotificationId uint64
    Limit uint16
}

type DataStoreNotificationV1 struct {
    Base Structure
    NotificationId uint64
    DataId uint32
}

type DataStoreNotification struct {
    Base Structure
    NotificationId uint64
    DataId uint64
}

type DataStoreRateObjectParam struct {
    Base Structure
    RatingValue uint32
    AccessPassword uint64
}

type DataStoreRatingTarget struct {
    Base Structure
    DataId uint64
    Slot uint8
}

type DataStoreGetSpecificMetaParamV1 struct {
    Base Structure
    DataIds []uint32
}

type DataStoreGetSpecificMetaParam struct {
    Base Structure
    DataIds []uint64
}

type DataStoreSpecificMetaInfoV1 struct {
    Base Structure
    DataId uint32
    OwnerId uint64
    Size uint32
    DataType uint16
    Version uint16
}

type DataStoreSpecificMetaInfo struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    DataType uint16
    Version uint32
}

type DataStoreTouchObjectParam struct {
    Base Structure
    DataId uint64
    LockId uint32
    AccessPassword uint64
}

type DataStoreRatingLog struct {
    Base Structure
    IsRated bool
    Pid uint64
    RatingValue uint32
    LockExpirationTime uint64
}

type DataStorePersistenceInfo struct {
    Base Structure
    OwnerId uint64
    PersistenceSlotId uint16
    DataId uint64
}

type DataStorePasswordInfo struct {
    Base Structure
    DataId uint64
    AccessPassword uint64
    UpdatePassword uint64
}

type BufferQueueParam struct {
    Base Structure
    DataId uint64
    Slot uint32
}

type DataStoreSearchBalloonParam struct {
    Base Structure
    DataType uint16
    UserRank uint8
    ResultSetCount uint8
}

type DataStoreSearchBalloonResultSet struct {
    Base Structure
    Balloons []DataStoreSearchBalloonResult
}

type DataStoreFetchMyInfosParam struct {
    Base Structure
    BalloonDataTypes []uint16
    AdditionalOperation uint16
}

type DataStoreFetchMyInfosResult struct {
    Base Structure
    Balloons []DataStoreFetchMyInfosBalloonResult
    Achievement DataStoreFetchMyInfosAchievementResult
}

type ResultRange struct {
    Base Structure
    UiOffset uint32
    UiSize uint32
}

type DataStorePersistenceTarget struct {
    Base Structure
    OwnerId uint64
    PersistenceSlotId uint16
}

type DataStoreKeyValue struct {
    Base Structure
    Key String
    Value String
}

type DataStorePermission struct {
    Base Structure
    Permission uint8
    RecipientIds []uint64
}

type DataStoreRatingInitParamWithSlot struct {
    Base Structure
    Slot uint8
    Param DataStoreRatingInitParam
}

type DataStorePersistenceInitParam struct {
    Base Structure
    PersistenceSlotId uint16
    DeleteLastObject bool
}

type DataStoreChangeMetaCompareParam struct {
    Base Structure
    ComparisonFlag uint32
    Name String
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary bytes
    Tags []String
    ReferredCnt uint32
    DataType uint16
    Status uint8
}

type DataStoreSearchBalloonResult struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    Name String
    DataType uint16
    MetaBinary bytes
    CreatedTime uint64
    UpdatedTime uint64
    OwnerDataId uint64
    OwnerName String
    IsFriendBalloon bool
    Ratings [uint8]DataStoreRatingInfo
    OwnerRatings [uint8]DataStoreRatingInfo
}

type DataStoreFetchMyInfosBalloonResult struct {
    Base Structure
    DataId uint64
    DataType uint16
    MetaBinary bytes
    CreatedTime uint64
    UpdatedTime uint64
    IsCleared bool
    Ratings [uint8]DataStoreRatingInfo
    Buffers [uint8][]bytes
}

type DataStoreFetchMyInfosAchievementResult struct {
    Base Structure
    DataId uint64
    DataType uint16
    MetaBinary bytes
    CreatedTime uint64
    Ratings [uint8]DataStoreRatingInfo
    Buffers [uint8][]bytes
}

type DataStoreRatingInitParam struct {
    Base Structure
    Flag uint8
    InternalFlag uint8
    LockType uint8
    InitialValue uint64
    RangeMin uint32
    RangeMax uint32
    PeriodHour uint8
    PeriodDuration uint16
}

type FriendPersistentInfo struct {
    Unknown uint32
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
    GameKey GameKey
    Unknown2 String
    Unknown3 uint64
    Unknown4 uint64
    Unknown5 uint64
}

type FriendPicture struct {
    Unknown uint32
    Data bytes
    DateTime uint64
}

type FriendPresence struct {
    Unknown uint32
    NintendoPresence NintendoPresence
}

type FriendRelationship struct {
    Unknown uint32
    Unknown2 uint64
    Unknown3 uint8
}

type GameKey struct {
    TitleId uint64
    TitleVersion uint16
}

type Mii struct {
    Unknown String
    Unknown2 bool
    Unknown3 uint8
    MiiData bytes
}

type MiiList struct {
    Unknown String
    Unknown2 bool
    Unknown3 uint8
    MiiDataList []bytes
}

type MyProfile struct {
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
    Unknown uint64
    Unknown2 String
    Unknown3 String
}

type NintendoPresence struct {
    ChangedBitFlag uint32
    GameKey GameKey
    GameModeDescription String
    JoinAvailabilityFlag uint32
    MatchmakeSystemType uint8
    JoinGameID uint32
    JoinGameMode uint32
    OwnerPrincipalID uint32
    JoinGroupID uint32
    ApplicationArg bytes
}

type PlayedGame struct {
    GameKey GameKey
    DateTime uint64
}

type BlacklistedPrincipal struct {
    PrincipalBasicInfo PrincipalBasicInfo
    GameKey GameKey
    BlacklistedSince uint64
}

type Comment struct {
    Unknown uint8
    StatusMessage String
    LastChangedOn uint64
}

type FriendInfo struct {
    NNAInfo NNAInfo
    NintendoPresence NintendoPresenceV2
    StatusMessage Comment
    BecameFriend uint64
    LastOnline uint64
    Unknown uint64
}

type FriendRequest struct {
    PrincipalBasicInfo PrincipalBasicInfo
    Message FriendRequestMessage
    SentOn uint64
}

type FriendRequestMessage struct {
    Unknown uint64
    Unknown2 uint8
    Unknown3 uint8
    Message String
    Unknown4 uint8
    Unknown5 String
    GameKey GameKey
    Unknown6 uint64
    ExpiresOn uint64
}

type GameKey struct {
    TitleId uint64
    TitleVersion uint16
}

type MiiV2 struct {
    Name String
    Unknown uint8
    Unknown2 uint8
    MiiDataFFLStoreData bytes
    Unknown3 uint64
}

type NintendoPresenceV2 struct {
    ChangedFlags uint32
    IsOnline bool
    GameKey GameKey
    Unknown1 uint8
    Message String
    Unknown2 uint32
    Unknown3 uint8
    GameServerId uint32
    Unknown4 uint32
    Pid uint32
    GatheringId uint32
    ApplicationData bytes
    Unknown5 uint8
    Unknown6 uint8
    Unknown7 uint8
}

type NNAInfo struct {
    PrincipalBasicInfo PrincipalBasicInfo
    Unknown uint8
    Unknown2 uint8
}

type PersistentNotification struct {
    Unknown uint64
    Unknown2 uint32
    Unknown3 uint32
    Unknown4 uint32
    Unknown5 String
}

type PrincipalBasicInfo struct {
    Pid uint32
    NNID String
    Mii MiiV2
    Unknown uint8
}

type PrincipalPreference struct {
    Unknown bool
    Unknown2 bool
    Unknown3 bool
}

type PrincipalRequestBlockSetting struct {
    Unknown uint32
    Unknown2 bool
}

type FriendData struct {
    Pid uint32
    StrName String
    ByRelationship uint8
    UiDetails uint32
    StrStatus String
}

type RelationshipData struct {
    Pid uint32
    StrName String
    ByRelationship uint8
    UiDetails uint32
    ByStatus uint8
}

type Gathering struct {
    Base Structure
    IdMyself uint32
    PidOwner uint32
    PidHost uint32
    UiMinParticipants uint16
    UiMaxParticipants uint16
    UiParticipationPolicy uint32
    UiPolicyArgument uint32
    UiFlags uint32
    UiState uint32
    StrDescription String
}

type PersistentGathering struct {
    Base Structure
    Base2 Gathering
    CommunityType uint32
    Password String
    Attribs []uint32
    ApplicationBuffer bytes
    ParticipationStartDate uint64
    ParticipationEndDate uint64
    MatchmakeSessionCount uint32
    ParticipationCount uint32
}

type MatchmakeSession struct {
    Base Structure
    Base2 Gathering
    GameMode uint32
    Attribs []uint32
    OpenParticipation bool
    MatchmakeSystemType uint32
    ApplicationBuffer bytes
    ParticipationCount uint32
    ProgressScore uint8
    SessionKey bytes
    Option0 uint32
    MatchmakeParam MatchmakeParam
    StartedTime uint64
    UserPassword String
    ReferGid uint32
    UserPasswordEnabled bool
    SystemPasswordEnabled bool
    Codeword String
}

type MatchmakeSessionSearchCriteria struct {
    Base Structure
    Attribs []String
    GameMode String
    MinParticipants String
    MaxParticipants String
    MatchmakeSystemType String
    VacantOnly bool
    ExcludeLocked bool
    ExcludeNonHostPid bool
    SelectionMethod uint32
    VacantParticipants uint16
    MatchmakeParam MatchmakeParam
    ExcludeUserPasswordSet bool
    ExcludeSystemPasswordSet bool
    ReferGid uint32
    Codeword String
    ResultRange ResultRange
}

type CreateMatchmakeSessionParam struct {
    Base Structure
    SourceMatchmakeSession MatchmakeSession
    AdditionalParticipants []uint32
    GidForParticipationCheck uint32
    CreateMatchmakeSessionOption uint32
    JoinMessage String
    ParticipationCount uint16
}

type JoinMatchmakeSessionParam struct {
    Base Structure
    Gid uint32
    AdditionalParticipants []uint32
    GidForParticipationCheck uint32
    JoinMatchmakeSessionOption uint32
    JoinMatchmakeSessionBehavior uint8
    StrUserPassword String
    StrSystemPassword String
    JoinMessage String
    ParticipationCount uint16
    ExtraParticipants uint16
    BlockListParam MatchmakeBlockListParam
}

type UpdateMatchmakeSessionParam struct {
    Base Structure
    Gid uint32
    ModificationFlag uint32
    Attributes []uint32
    OpenParticipation bool
    ApplicationBuffer bytes
    ProgressScore uint8
    MatchmakeParam MatchmakeParam
    StartedTime uint64
    UserPassword String
    GameMode uint32
    Description String
    MinParticipants uint16
    MaxParticipants uint16
    MatchmakeSystemType uint32
    ParticipationPolicy uint32
    PolicyArgument uint32
    Codeword String
}

type MatchmakeBlockListParam struct {
    Base Structure
    OptionFlag uint32
}

type MatchmakeParam struct {
    Base Structure
    Params [String]Variant
}

type AutoMatchmakeParam struct {
    Base Structure
    SourceMatchmakeSession MatchmakeSession
    AdditionalParticipants []uint32
    GidForParticipationCheck uint32
    AutoMatchmakeOption uint32
    JoinMessage String
    ParticipationCount uint16
    LstSearchCriteria []MatchmakeSessionSearchCriteria
    TargetGids []uint32
    BlockListParam MatchmakeBlockListParam
}

type FindMatchmakeSessionByParticipantParam struct {
    PrincipalIdList []uint32
    ResultOptions uint32
    BlockListParam MatchmakeBlockListParam
}

type FindMatchmakeSessionByParticipantResult struct {
    PrincipalId uint32
    Session MatchmakeSession
}

type GatheringURLs struct {
    Base Structure
    Gid uint32
    LstStationURLs []String
}

type Gathering Stats struct {
    Base Structure
    PidParticipant uint32
    UiFlags uint32
    LstValues []float
}

type Invitation struct {
    Base Structure
    IdGathering uint32
    IdGuest uint32
    StrMessage String
}

type Participant Details struct {
    Base Structure
    IdParticipant uint32
    StrName String
    StrMessage String
    UiParticipants uint16
}

type Deletion Entry struct {
    Base Structure
    IdGathering uint32
    Pid uint32
    UiReason uint32
}

type PlayingSession struct {
    Base Structure
    PrincipalId uint32
    Gathering Data
}

type SimplePlayingSession struct {
    Base Structure
    PrincipalID uint32
    GatheringID uint32
    GameMode uint32
    Attribute_0 uint32
}

type SimpleCommunity struct {
    Base Structure
    GatheringID uint32
    MatchmakeSessionCount uint32
}

type MessageRecipient struct {
    UiRecipientType uint32
    PrincipalId uint32
    GatheringId uint32
}

type UserMessage struct {
    UiID uint32
    UiParentID uint32
    PidSender uint32
    Receptiontime uint64
    UiLifeTime uint32
    UiFlags uint32
    StrSubject String
    StrSender String
    MessageRecipient MessageRecipient
}

type NintendoNotificationEvent struct {
    Base Structure
    UiType uint32
    PidSender uint32
    Dataholder Data
}

type NintendoNotificationEventGeneral struct {
    Base Structure
    U32Param uint32
    U64Param1 uint64
    U64Param2 uint64
    StrParam String
}

type NintendoNotificationEventProfile struct {
    Base Structure
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
}

type RankingOrderParam struct {
    Base Structure
    OrderCalculation uint8
    GroupIndex uint8
    GroupNum uint8
    TimeScope uint8
    Offset uint32
    Length uint8
}

type RankingRankData struct {
    Base Structure
    PrincipalId uint32
    UniqueId uint64
    Order uint32
    Category uint32
    Score uint32
    Groups []uint8
    Param uint64
    CommonData bytes
}

type RankingResult struct {
    Base Structure
    RankDataList []RankingRankData
    TotalCount uint32
    SinceTime uint64
}

type RankingCachedResult struct {
    CreatedTime uint64
    ExpiredTime uint64
    MaxLength uint8
}

type RankingStats struct {
    Base Structure
    StatsList []double
}

type RankingScoreData struct {
    Base Structure
    PrincipalId uint32
    UniqueId uint64
    Order uint32
    Category uint32
    Score uint32
    Groups []uint8
    Param uint64
    CommonData bytes
}

type RankingChangeAttributesParam struct {
    Base Structure
    ModificationFlag uint8
    Groups []uint8
    Param uint64
}

type UniqueIdInfo struct {
    Base Structure
    NexUniqueId uint64
    NexUniqueIdPassword uint64
}

