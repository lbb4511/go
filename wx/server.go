package main

import (
	"log"
	"net/http"

	"github.com/sidbusy/weixinmp"
)

func main() {
	// 注册处理函数
	http.HandleFunc("/", receiver)
	log.Fatal(http.ListenAndServe(":8098", nil))
}

func receiver(w http.ResponseWriter, r *http.Request) {
	token := "Adhsu1H3SAU5hw3s3ad"               // 微信公众平台的Token
	appid := "wx5630e04b1c3f6f2a"                // 微信公众平台的AppID
	secret := "b9aebac4dd35d61e7991760462803c1b" // 微信公众平台的AppSecret
	// 仅被动响应消息时可不填写appid、secret
	// 仅主动发送消息时可不填写token
	mp := weixinmp.New(token, appid, secret)
	// 检查请求是否有效
	// 仅主动发送消息时不用检查
	if !mp.Request.IsValid(w, r) {
		return
	}
	// 判断消息类型
	if mp.Request.MsgType == weixinmp.MsgTypeText {
		// 回复消息
		mp.ReplyTextMsg(w, "Hello, 世界")
	}
}
