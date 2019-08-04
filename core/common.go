package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/utils"
	"net/http"
)

type OAuthCallback interface {
	SendAccessToken(params map[string]string, config ConfigValue, callback func(err error, token *Token))
	SendUserInfo(request *http.Request, callback func(err error, response map[string]interface{}))
}

type DefaultOAuthCallback struct {

}

func (cb *DefaultOAuthCallback) SendAccessToken(params map[string]string,
	config ConfigValue,
	callback func(err error, token *Token)) {

	builder := utils.NewUrlBuilder(config.AccessTokenUri)
	builder.Add(consts.CLIENT_ID, config.ClientId)
	builder.Add(consts.CLIENT_SECRET, config.ClientSecret)
	builder.Add(consts.STATE, config.State)

	utils.MapForeach(params, func(key string, value string) {
		builder.Add(key, value)
	})
	postPath := builder.Build()

	req, _ := http.NewRequest("GET", postPath, nil)
	req.Header.Add(consts.ACCEPT_KEY, consts.CONTENT_TYPE_JSON)


	utils.NewHttpClient().Do(req, func(response []byte) {
		tmp := &Token{}
		if err := json.Unmarshal(response, tmp) ; err != nil{
			callback(err, nil)
		}else {
			if tmp.Check() {
				callback(nil, tmp)
				return
			}
			callback(errors.New(fmt.Sprintf("get access_token error, response content=%v", string(response))),
				nil)
		}

	}, func(err error, code int, response []byte) {
		callback(err, nil)
	})
}

func (cb *DefaultOAuthCallback) SendUserInfo(request *http.Request,
	callback func(err error, response map[string]interface{})) {

	request.Header.Add(consts.ACCEPT_KEY, consts.CONTENT_TYPE_JSON)
	ret := make(map[string]interface{})

	utils.NewHttpClient().Do(request, func(response []byte) {
		if err := json.Unmarshal(response, &ret); err != nil {
			callback(err, nil)
		}else {
			callback(nil, ret)
		}
	}, func(err error, code int, response []byte) {
		callback(err, ret)
	})

}
