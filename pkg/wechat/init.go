package wechat

import (
	wechat "github.com/go-pay/gopay/wechat"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	"time"
)

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type Callback struct {
	Id           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     Resource `json:"resource"`
}
type Response struct {
	Mchid          string    `json:"mchid"`
	Appid          string    `json:"appid"`
	OutTradeNo     string    `json:"out_trade_no"`
	TransactionId  string    `json:"transaction_id"`
	TradeType      string    `json:"trade_type"`
	TradeState     string    `json:"trade_state"`
	TradeStateDesc string    `json:"trade_state_desc"`
	BankType       string    `json:"bank_type"`
	Attach         string    `json:"attach"`
	SuccessTime    time.Time `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int    `json:"total"`
		PayerTotal    int    `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}
var wepay *WehcatConfig
var wechatClient WechatClient
type WehcatConfig struct {
	AppId            string
	Appkey           string
	MchId            string
	SerialNo         string
	ApiKey3          string
	ApiclientKeyPath string
	ApiclientCerPath string
	NotifyUrl        string
	IsProd           bool
	ApiclientCer12Path string
}
var v2client *wechat.Client
var v3client *wechat3.ClientV3
