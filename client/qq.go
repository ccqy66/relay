package client

import (
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/core"
	"github.com/ccqy66/relay/utils"
	"net/http"
)

type QQOAuth struct {
	core.ConfigValue
}

func (qq *QQOAuth) GenerateAuthUrl() (string, error) {
	builder := utils.NewUrlBuilder(qq.RedirectUri)
	builder.Add(consts.CLIENT_ID, qq.ClientId)
	builder.Add(consts.SCOPE, qq.Scope)
	builder.Add(consts.STATE, qq.State)
	builder.Add(consts.RESPONSE_TYPE, consts.CODE_TYPE)
	return builder.Build(), nil
}

func (qq *QQOAuth) GetAccessToken(request map[string]string) (*core.Token, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var t *core.Token
	request[consts.GRANT_TYPE] = consts.AUTHORIZATION_KEY
	callback.SendAccessToken(request, qq.ConfigValue, func(err error, token *core.Token) {
		if err != nil {
			e = err
		}else {
			t = token
		}
	})
	return t, e
}

func (qq *QQOAuth) GetUserInfo(token *core.Token) (map[string]interface{}, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var ret map[string]interface{}
	builder := utils.NewUrlBuilder(qq.UserInfoUri)
	reqPath := builder.Add(consts.ACCESS_TOKEN, token.AccessToken).
		Build()
	request, _ := http.NewRequest("GET", reqPath, nil)
	callback.SendUserInfo(request, func(err error, response map[string]interface{}) {
		if err != nil {
			e = err
		}else {
			ret = response
		}
	})
	return ret, e
}




