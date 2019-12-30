package application

// CreateApplicatonRequest 创建应用请求
type CreateApplicatonRequest struct {
	Name            string `bson:"name" json:"name"`                           // 应用名称
	Website         string `bson:"website" json:"website,omitempty"`           // 应用的网站地址
	LogoImage       string `bson:"logo_image" json:"logo_image,omitempty"`     // 应用的LOGO
	Description     string `bson:"description" json:"description"`             // 应用简单的描述
	RedirectURI     string `bson:"redirect_uri" json:"redirect_uri"`           // 应用重定向URI, Oauht2时需要改参数
	ClientID        string `bson:"client_id" json:"client_id"`                 // 应用客户端ID
	ClientSecret    string `bson:"client_secret" json:"client_secret"`         // 应用客户端秘钥
	TokenExpireTime int64  `bson:"token_expire_time" json:"token_expire_time"` // 应用申请的token的过期时间
}

// Application is oauth2's client: https://tools.ietf.org/html/rfc6749#section-2
type Application struct {
	ID                       string `json:"id"`        // 唯一ID
	UserID                   string `json:"user_id"`   // 应用属于那个用户
	CreateAt                 int64  `json:"create_at"` // 应用创建的时间
	UpdateAt                 int64  `json:"update_at"` // 应用更新的时间
	Locked                   bool   `json:"locked"`    // 是否冻结应用, 冻结应用后, 该应用无法通过凭证获取访问凭证(token)
	*CreateApplicatonRequest `bson:",inline"`
}
