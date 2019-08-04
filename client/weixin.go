package client

import (
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/core"
	"github.com/ccqy66/relay/utils"
)

type WeChatOAuth struct {
	core.ConfigValue
}

func (wco *WeChatOAuth) GenerateAuthUrl() (string, error) {
	builder := utils.NewUrlBuilder(wco.RedirectUri)
	builder.Add(consts.APPID, wco.AppId)
	builder.Add(consts.SCOPE, wco.Scope)
	builder.Add(consts.STATE, wco.State)
	builder.Add(consts.RESPONSE_TYPE, consts.CODE_TYPE)
	return builder.Build(), nil
}

func (wco *WeChatOAuth) GetUserInfo(token *core.Token) (map[string]interface{}, error) {
	return nil,nil
}

func (wco *WeChatOAuth) GetAccessToken(request map[string]string) (*core.Token, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var t *core.Token
	request[consts.APPID] = wco.AppId
	request[consts.SECRET] = wco.ClientSecret
	request[consts.GRANT_TYPE] = consts.AUTHORIZATION_CDOE
	callback.SendAccessToken(request, wco.ConfigValue, func(err error, token *core.Token) {
		if err != nil {
			e = err
		}else {
			t = token
		}
	})
	return t, e
}
