package client

import (
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/core"
	"github.com/ccqy66/relay/utils"
	"net/http"
	"strconv"
)

type WeiboOAuth struct {
	core.ConfigValue
}


func (wb *WeiboOAuth) GenerateAuthUrl() (string, error) {
	builder := utils.NewUrlBuilder(wb.RedirectUri)
	builder.Add(consts.CLIENT_ID, wb.ClientId)
	builder.Add(consts.SCOPE, wb.Scope)
	builder.Add(consts.STATE, wb.State)
	builder.Add(consts.RESPONSE_TYPE, consts.CODE_TYPE)
	return builder.Build(), nil
}

func (wb *WeiboOAuth) GetAccessToken(request map[string]string) (*core.Token, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var t *core.Token
	callback.SendAccessToken(request, wb.ConfigValue, func(err error, token *core.Token) {
		if err != nil {
			e = err
		}else {
			t = token
		}
	})
	return t, e
}

func (wb *WeiboOAuth) GetUserInfo(token *core.Token) (map[string]interface{}, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var ret map[string]interface{}
	builder := utils.NewUrlBuilder(wb.UserInfoUri)
	reqPath := builder.Add(consts.ACCESS_TOKEN, token.AccessToken).
		Add(consts.UID_KEY, strconv.FormatInt(token.Uid, 10)).
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

