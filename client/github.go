package client

import (
	"fmt"
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/core"
	"github.com/ccqy66/relay/utils"
	"net/http"
)

type GithubOAuth struct {
	core.ConfigValue
}

func (github *GithubOAuth) GenerateAuthUrl() (string, error) {
	builder := utils.NewUrlBuilder(github.RedirectUri)
	builder.Add(consts.CLIENT_ID, github.ClientId)
	builder.Add(consts.SCOPE, github.Scope)
	builder.Add(consts.STATE, github.State)
	return builder.Build(), nil
}

func (github *GithubOAuth) GetAccessToken(request map[string]string) (*core.Token, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var t *core.Token
	callback.SendAccessToken(request, github.ConfigValue, func(err error, token *core.Token) {
		if err != nil {
			e = err
		}else {
			t = token
		}
	})
	return t, e
}

func (github *GithubOAuth) GetUserInfo(token *core.Token) (map[string]interface{}, error) {
	callback := core.DefaultOAuthCallback{}
	var e error
	var ret map[string]interface{}
	request, _ := http.NewRequest("GET", github.UserInfoUri, nil)
	request.Header.Add(consts.AUTHORIZATION_KEY, fmt.Sprintf("token %s", token.AccessToken))
	callback.SendUserInfo(request, func(err error, response map[string]interface{}) {
		if err != nil {
			e = err
		}else {
			ret = response
		}
	})
	return ret, e
}
