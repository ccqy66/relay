package relay

import (
	"encoding/json"
	"errors"
	client2 "github.com/ccqy66/relay/client"
	"github.com/ccqy66/relay/consts"
	"github.com/ccqy66/relay/core"
	"io/ioutil"
)

type Client struct {
	ConfigTable map[string]core.ConfigValue
}

// get a relay client by config file path
// config file format is json:
// for example:
//	{
//  "weixin": {
//    "appid": "",
//    "redirect_uri": "",
//    "reponse_type": "",
//    "scope": "",
//    "state": ""
//  },
//  "github": {
//
//  }
//}
// you can define multi config for using json

func NewClient(configPath string) (*Client, error) {
	if configPath == "" {
		return nil, errors.New("config path is empty")
	}
	var config map[string]core.ConfigValue
	if c, err := ioutil.ReadFile(configPath); err != nil {
		return nil, err
	} else {
		if json.Unmarshal(c, &config) != nil {
			return nil, err
		}
	}
	if config == nil || len(config) <= 0 {
		return nil, errors.New("new client error")
	}
	return &Client{ConfigTable: config}, nil
}

func (client *Client) New(t string) core.OAuth {
	if client == nil {
		return nil
	}
	if client.ConfigTable == nil || len(client.ConfigTable) <= 0 {
		return nil
	}
	config := client.ConfigTable[t]
	if !config.Check() {
		return nil
	}
	switch t {
	case consts.WECHAT:
		return &client2.WeChatOAuth{
			ConfigValue: config,
		}
	case consts.GITHUB:
		return &client2.GithubOAuth{
			ConfigValue: config,
		}
	case consts.QQ:
		return &client2.QQOAuth{
			ConfigValue: config,
		}
	case consts.WEIBO:
		return &client2.WeiboOAuth{
			ConfigValue: config,
		}
	default:
		return nil
	}
}

func (client *Client) WeChat() core.OAuth {
	return client.New(consts.WECHAT)
}

func (client *Client) Github() core.OAuth {
	return client.New(consts.GITHUB)
}

func (client *Client) QQ() core.OAuth {
	return client.New(consts.QQ)
}

func (client *Client) Weibo() core.OAuth {
	return client.New(consts.WEIBO)
}

