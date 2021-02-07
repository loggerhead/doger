package api

const (
	AllGender = 0
	Male      = 1
	Female    = 2
)

type UserInfo struct {
	ID           int    `bson:"_id" json:"id,omitempty"`
	UserID       int    `json:"userId,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Country      string `json:"country,omitempty"`
	Status       int    `json:"status,omitempty"`
	Remark       string `json:"remark,omitempty"`
	CertStatus   int    `json:"certStatus,omitempty"`
	DataLevel    int    `json:"dataLevel,omitempty"`
	VipLevel     int    `json:"vipLevel,omitempty"`
	HiddenFields string `json:"hiddenFields,omitempty"`
	Basicinfo    struct {
		ID               int    `json:"id,omitempty"`
		LoverUserID      int    `json:"loverUserId,omitempty"`
		UID              string `json:"uid,omitempty"`
		Phone            string `json:"phone,omitempty"`
		Nickname         string `json:"nickname,omitempty"`
		Wechatid         string `json:"wechatid,omitempty"`
		Avatar           string `json:"avatar,omitempty"`
		Photos           string `json:"photos,omitempty"`
		Gender           int    `json:"gender,omitempty"`
		Birthday         string `json:"birthday,omitempty"`
		BirthYear        string `json:"birthYear,omitempty"`
		Constellation    string `json:"constellation,omitempty"`
		Height           int    `json:"height,omitempty"`
		Weight           int    `json:"weight,omitempty"`
		Company          string `json:"company,omitempty"`
		Position         string `json:"position,omitempty"`
		Isabroad         int    `json:"isabroad,omitempty"`
		Abroadcountry    string `json:"abroadcountry,omitempty"`
		HomelandProvince string `json:"homelandProvince,omitempty"`
		HomelandCity     string `json:"homelandCity,omitempty"`
		LivingProvince   string `json:"livingProvince,omitempty"`
		LivingCity       string `json:"livingCity,omitempty"`
		Highestschool    string `json:"highestschool,omitempty"`
		Highestdegree    int    `json:"highestdegree,omitempty"`
		Annualincome     int    `json:"annualincome,omitempty"`
		Marital          int    `json:"marital,omitempty"`
		DataDeleted      int    `json:"dataDeleted,omitempty"`
	} `json:"basicinfo,omitempty"`
	Card []struct {
		ID            int         `json:"id,omitempty"`
		LoverUserID   int         `json:"loverUserId,omitempty"`
		Type          string      `json:"type,omitempty"`
		Placeholder   string      `json:"placeholder,omitempty"`
		Topic         string      `json:"topic,omitempty"`
		Title         string      `json:"title,omitempty"`
		Contenttext   string      `json:"contenttext,omitempty"`
		Contentimages string      `json:"contentimages,omitempty"`
		Isshow        int         `json:"isshow,omitempty"`
		Isdefault     int         `json:"isdefault,omitempty"`
		Weigh         int         `json:"weigh,omitempty"`
		Createtime    int         `json:"createtime,omitempty"`
		Updatetime    interface{} `json:"updatetime,omitempty"`
		Record        interface{} `json:"record,omitempty"`
		DataDeleted   int         `json:"dataDeleted,omitempty"`
	} `json:"card,omitempty"`
	Certification []struct {
		Type        string `json:"type,omitempty"`
		Status      int    `json:"status,omitempty"`
		IdentifyWay string `json:"identifyWay,omitempty"`
		Credibility int    `json:"credibility,omitempty"`
	} `json:"certification,omitempty"`
	CertificationMap struct {
		Work   int `json:"work,omitempty"`
		Idcard int `json:"idcard,omitempty"`
		Degree int `json:"degree,omitempty"`
	} `json:"certificationMap,omitempty"`
	CertificationDetail struct {
		Credibility   int `json:"credibility,omitempty"`
		Certification []struct {
			Type        string `json:"type,omitempty"`
			Status      int    `json:"status,omitempty"`
			IdentifyWay string `json:"identifyWay,omitempty"`
			Credibility int    `json:"credibility,omitempty"`
		} `json:"certification,omitempty"`
	} `json:"certificationDetail,omitempty"`
	UserUUID       string `json:"userUuid,omitempty"`
	BlockSchool    bool   `json:"blockSchool,omitempty"`
	BlockCompany   bool   `json:"blockCompany,omitempty"`
	Collected      bool   `json:"collected,omitempty"`
	LoverBlocked   bool   `json:"loverBlocked,omitempty"`
	TargetBlocked  bool   `json:"targetBlocked,omitempty"`
	ShareTitle     string `json:"shareTitle,omitempty"`
	FriendStatus   int    `json:"friendStatus,omitempty"`
	FriendStatusV2 int    `json:"friendStatusV2,omitempty"`
	DynamicNum     int    `json:"dynamicNum,omitempty"`
	NeedHideRemark bool   `json:"needHideRemark,omitempty"`
	Prerogative    string `json:"prerogative,omitempty"`
	Role           int    `json:"role,omitempty"`
	Identity       string `json:"identity,omitempty"`
	Lbs            struct {
		District string `json:"district,omitempty"`
		Distance string `json:"distance,omitempty"`
	} `json:"lbs,omitempty"`
	NewTag []struct {
		SortID  int    `json:"sortId,omitempty"`
		Sort    string `json:"sort,omitempty"`
		Order   int    `json:"order,omitempty"`
		TagList []struct {
			Idx  int    `json:"idx,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"tagList,omitempty"`
	} `json:"newTag,omitempty"`
}

type ReplyInfo struct {
	FromUserAvatar   string `json:"fromUserAvatar,omitempty"`
	FromUserUid      string `json:"fromUserUid,omitempty"`
	FromUserNickname string `json:"fromUserNickname,omitempty"`
	Content          string `json:"content,omitempty"`
	LikeNum          int    `json:"likeNum,omitempty"`
	FromIsMask       bool   `json:"fromIsMask,omitempty"`
	IsOwner          bool   `json:"isOwner,omitempty"`
	IsOfficial       bool   `json:"isOfficial,omitempty"`
}

type DynamicDetail struct {
	DynamicID       int          `bson:"_id" json:"dynamicId,omitempty"`
	UID             string       `json:"uid,omitempty"`
	Nickname        string       `json:"nickname,omitempty"`
	City            string       `json:"city,omitempty"`
	HomelandCity    string       `json:"homelandCity,omitempty"`
	Birthday        string       `json:"birthday,omitempty"`
	BirthYear       string       `json:"birthYear,omitempty"`
	Avatar          string       `json:"avatar,omitempty"`
	Gender          int          `json:"gender,omitempty"`
	TopicID         int          `json:"topicId,omitempty"`
	Position        string       `json:"position,omitempty"`
	TopicContent    string       `json:"topicContent,omitempty"`
	DynamicImages   string       `json:"dynamicImages,omitempty"`
	DynamicContent  string       `json:"dynamicContent,omitempty"`
	DynamicType     int          `json:"dynamicType,omitempty"`
	CreateTime      string       `json:"createTime,omitempty"`
	DynamicAreaDesc string       `json:"dynamicAreaDesc,omitempty"`
	Longitude       string       `json:"longitude,omitempty"`
	Latitude        string       `json:"latitude,omitempty"`
	WatchNum        int          `json:"watchNum,omitempty"`
	LikeNum         int          `json:"likeNum,omitempty"`
	LikeUserAvatar  string       `json:"likeUserAvatar,omitempty"`
	CommentNum      int          `json:"commentNum,omitempty"`
	Tag             string       `json:"tag,omitempty"`
	IsOfficial      bool         `json:"isOfficial,omitempty"`
	IsStick         bool         `json:"isStick,omitempty"`
	IsOpen          bool         `json:"isOpen,omitempty"`
	IsMask          bool         `json:"isMask,omitempty"`
	OwnerLike       bool         `json:"ownerLike,omitempty"`
	UserDeleted     bool         `json:"userDeleted,omitempty"`
	ShareNum        int          `json:"shareNum,omitempty"`
	ExposeNum       int          `json:"exposeNum,omitempty"`
	ReplyInfos      []*ReplyInfo `json:"replyInfos,omitempty"`
}

type DynamicComment struct {
	CommentID            int               `bson:"_id" json:"commentId,omitempty"`
	FromUserAvatar       string            `json:"fromUserAvatar,omitempty"`
	FromUserNickname     string            `json:"fromUserNickname,omitempty"`
	FromUserCity         string            `json:"fromUserCity,omitempty"`
	FromUserHomelandCity string            `json:"fromUserHomelandCity,omitempty"`
	FromUserBirthday     string            `json:"fromUserBirthday,omitempty"`
	FromUserBirthYear    string            `json:"fromUserBirthYear,omitempty"`
	FromUserUID          string            `json:"fromUserUid,omitempty"`
	ToUserAvatar         string            `json:"toUserAvatar,omitempty"`
	ToUserNickname       string            `json:"toUserNickname,omitempty"`
	ToUserCity           string            `json:"toUserCity,omitempty"`
	ToUserBirthday       string            `json:"toUserBirthday,omitempty"`
	ToUserBirthYear      string            `json:"toUserBirthYear,omitempty"`
	ToUserUID            string            `json:"toUserUid,omitempty"`
	ToIsMask             bool              `json:"toIsMask,omitempty"`
	ToCost               int               `json:"toCost,omitempty"`
	Level                int               `json:"level,omitempty"`
	LikeNum              int               `json:"likeNum,omitempty"`
	CommentNum           int               `json:"commentNum,omitempty"`
	Position             string            `json:"position,omitempty"`
	Content              string            `json:"content,omitempty"`
	CreateTime           string            `json:"createTime,omitempty"`
	CreateTimeText       string            `json:"createTimeText,omitempty"`
	Gender               int               `json:"gender,omitempty"`
	ShareNum             int               `json:"shareNum,omitempty"`
	OwnerLike            bool              `json:"ownerLike,omitempty"`
	UserDeleted          bool              `json:"userDeleted,omitempty"`
	IsOfficial           bool              `json:"isOfficial,omitempty"`
	Comments             []*DynamicComment `json:"comments,omitempty"`
}

type DynamicListData struct {
	Type int            `json:"type,omitempty"`
	Data *DynamicDetail `json:"data,omitempty"`
}

type BaseResp struct {
	Success bool   `json:"success,omitempty"`
	Code    int    `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

type RespUserDetail struct {
	BaseResp
	Data struct {
		Type     int       `json:"type,omitempty"`
		UserInfo *UserInfo `json:"userInfo,omitempty"`
	} `json:"data,omitempty"`
}

type RespDynamicDetail struct {
	BaseResp
	Data *DynamicDetail `json:"data,omitempty"`
}

type RespDynamicList struct {
	BaseResp
	Data []DynamicListData `json:"data,omitempty"`
}

type RespDynamicCommentList struct {
	BaseResp
	Data []*DynamicComment `json:"data,omitempty"`
}

type RespDynamicLikeList struct {
	BaseResp
	Data []struct {
		UID        string `json:"uid,omitempty"`
		Avatar     string `json:"avatar,omitempty"`
		Nickname   string `json:"nickname,omitempty"`
		Birthday   string `json:"birthday,omitempty"`
		BirthYear  string `json:"birthYear,omitempty"`
		Position   string `json:"position,omitempty"`
		Gender     int    `json:"gender,omitempty"`
		IsOfficial bool   `json:"isOfficial,omitempty"`
	} `json:"data,omitempty"`
}
