package work_weixin_robot

import (
	"os"
	"testing"
)

func TestWorkWeixinRobotClient_SendMessageStr(t *testing.T) {
	json := `{
    "msgtype":"template_card",
    "template_card":{
        "card_type":"news_notice",
        "source":{
            "icon_url":"https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
            "desc":"企业微信",
            "desc_color":0
        },
        "main_title":{
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信"
        },
        "card_image":{
            "url":"https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0",
            "aspect_ratio":2.25
        },
        "image_text_area":{
            "type":1,
            "url":"https://work.weixin.qq.com",
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信",
            "image_url":"https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0"
        },
        "quote_area":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH",
            "title":"引用文本标题",
            "quote_text":"Jack：企业微信真的很好用~\nBalian：超级好的一款软件！"
        },
        "vertical_content_list":[
            {
                "title":"惊喜红包等你来拿",
                "desc":"下载企业微信还能抢红包！"
            }
        ],
        "horizontal_content_list":[
            {
                "keyname":"邀请人",
                "value":"张三"
            },
            {
                "keyname":"企微官网",
                "value":"点击访问",
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi"
            }
        ],
        "jump_list":[
            {
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi",
                "title":"企业微信官网"
            }
        ],
        "card_action":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH"
        }
    }
}`
	client := NewRobotClient()
	client.Webhook = os.Getenv("webhook")
	res, err := client.SendMessageStr(json)
	if err != nil {
		t.Error("send message error", err)
	}
	if res.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf(res.ErrMsg)
	}
}
