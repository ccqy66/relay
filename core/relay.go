package core

// mapping config
// if you add a new config properties you need to add new one

type ConfigValue struct {
	RedirectUri    string `json:"redirect_uri"`
	AccessTokenUri string `json:"access_token_uri"`
	UserInfoUri    string `json:"user_info_uri"`
	ClientId       string `json:"client_id"`
	AppId          string `json:"appid"`
	ClientSecret   string `json:"client_secret"`
	Scope          string `json:"scope"`
	State          string `json:"state"`
}

func (cv *ConfigValue) Check() bool {
	return true
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Uid         int64  `json:"uid"`
	ExpiresIn   int64  `json:"expires_in"`
	RemindIn    int64  `json:"remind_in"`
}

func (t *Token) Check() bool {
	return t.AccessToken != ""
}
func (t *Token) String() string {
	return "{" +
		"AccessToken=" + t.AccessToken + "," +
		"TokenType=" + t.TokenType + "," +
		"Scope=" + t.Scope + "," +
		"}"
}

// as a oauth interface
// if you implement a new oauth client, you must assign a local var for ConfigValue to receive config
// for example:
// type GithubOauth struct {
//		ConfigValue
// }
// ConfigValue will get github's config at config.json

type OAuth interface {
	// generate redirect url, Request a user's identity
	GenerateAuthUrl() (string, error)

	// get access_token by code and other extra params
	GetAccessToken(request map[string]string) (*Token, error)

	// get user info by access_token
	GetUserInfo(token *Token) (map[string]interface{}, error)
}
