package relay

import (
	"fmt"
	"testing"
)

func TestGithub(t *testing.T) {
	client,_ := NewClient("config.json")
	github := client.Github()
	fmt.Println(github.GenerateAuthUrl())

	//fmt.Println(github.GetAccessToken(map[string]string{
	//	"code":"0af2bb7bda379f478c1a",
	//}))

	//fmt.Println(github.GetUserInfo()
}

func TestWeibo(t *testing.T) {
	client,_ := NewClient("config.json")
	weibo := client.Weibo()
	fmt.Println(weibo.GenerateAuthUrl())

	//fmt.Println(github.GetAccessToken(map[string]string{
	//	"code":"0af2bb7bda379f478c1a",
	//}))

	//fmt.Println(github.GetUserInfo("1724ab4926e19985e27437948325a4880dbd9ef4"))
}


func TestQQ(t *testing.T) {
	client,_ := NewClient("config.json")
	qq := client.QQ()
	fmt.Println(qq.GenerateAuthUrl())
}