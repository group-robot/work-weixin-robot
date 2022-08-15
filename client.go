package work_weixin_robot

import (
	"github.com/go-resty/resty/v2"
)

// WorkWeixinRobotClient 企业微信-机器人客户端
type WorkWeixinRobotClient struct {
	// Webhook webhook address
	Webhook string
	client  *resty.Client
}

// NewRobotClient create WorkWeixinRobotClient
func NewRobotClient() *WorkWeixinRobotClient {
	return NewRobotClientByWebHook("")
}

// NewRobotClientByWebHook create WorkWeixinRobotClient By Webhook
func NewRobotClientByWebHook(webhook string) *WorkWeixinRobotClient {
	return &WorkWeixinRobotClient{
		Webhook: webhook,
		client:  resty.New(),
	}
}

// RobotResponse robot response
type RobotResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsSuccess is success
func (rep *RobotResponse) IsSuccess() bool {
	return rep.ErrCode == 0
}

// SendMessage send message
func (client *WorkWeixinRobotClient) SendMessage(message Message) (*RobotResponse, error) {
	return client.SendMessageByUrl(client.Webhook, message)
}

// SendMessageStr send message json string
func (client *WorkWeixinRobotClient) SendMessageStr(message string) (*RobotResponse, error) {
	return client.SendMessageStrByUrl(client.Webhook, message)
}

// SendMessageByUrl send message custom url
func (client *WorkWeixinRobotClient) SendMessageByUrl(url string, message Message) (*RobotResponse, error) {
	return client.send(url, message.ToMessageMap())

}

// SendMessageStrByUrl send message custom url and json string message
func (client *WorkWeixinRobotClient) SendMessageStrByUrl(url, message string) (*RobotResponse, error) {
	return client.send(url, message)
}

func (client *WorkWeixinRobotClient) send(url string, body interface{}) (*RobotResponse, error) {
	resp, err := client.client.R().
		SetHeader("Content-Type", "application/json").
		ForceContentType("application/json").
		SetBody(body).
		SetResult(&RobotResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*RobotResponse)
	return result, nil
}
