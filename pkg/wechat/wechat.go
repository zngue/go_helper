package wechat

import (
	"fmt"
	wechat "github.com/go-pay/gopay/wechat"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	commonWechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/spf13/viper"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
)

type WechatClient interface {

	V3Client() (*wechat3.ClientV3, error)
	CommonWechat() *commonWechat.Wechat
	V2Client() (*wechat.Client,error)
	OfficialAccount() *officialaccount.OfficialAccount
	DecryptToString(resource *Resource) (certificate string, err error)

}
type WechatService struct {
}

func (s *WechatService) DecryptToString(resource *Resource) (certificate string, err error) {
	return utils.DecryptAES256GCM(wepay.ApiKey3, resource.AssociatedData, resource.Nonce, resource.Ciphertext)
}
func (s *WechatService) ApiclientKey() (string, error) {
	paths := wepay.ApiclientKeyPath
	fmt.Println(wepay, paths)
	privateKeyBytes, keypathErr := ioutil.ReadFile(wepay.ApiclientKeyPath)
	if keypathErr != nil {
		return "", keypathErr
	}
	return string(privateKeyBytes), nil
}
func (s *WechatService) V3Client() (*wechat3.ClientV3, error) {
	var (
		err error
		privateKeyBytes string
	)
	if v3client!=nil {
		return v3client, nil
	}
	privateKeyBytes, err = s.ApiclientKey()
	if err != nil {
		return nil, err
	}
	v3client,err=wechat3.NewClientV3(wepay.MchId, wepay.SerialNo, wepay.ApiKey3, privateKeyBytes)
	return v3client,err

}
func (s *WechatService) V2Client() (*wechat.Client,error) {
	var err error
	if v2client!=nil {
		return v2client,err
	}
	v2client = wechat.NewClient(wepay.AppId, wepay.MchId, wepay.ApiKey3, wepay.IsProd)
	if err = v2client.AddCertPemFilePath(wepay.ApiclientCerPath, wepay.ApiclientKeyPath); err != nil {
		return nil, err
	}
	if err = v2client.AddCertPkcs12FilePath(wepay.ApiclientCer12Path); err != nil {
		return nil, err
	}
	return v2client,err

}
func (s *WechatService) CommonWechat() *commonWechat.Wechat {
	newWechat := commonWechat.NewWechat()
	return newWechat
}
func (s *WechatService) OfficialAccount() *officialaccount.OfficialAccount {
	newWechat := commonWechat.NewWechat()
	memory := cache.NewMemory()
	c := config.Config{
		AppID:     wepay.AppId,
		AppSecret: wepay.Appkey,
		Token:     "weixin",
		Cache:     memory,
	}
	return newWechat.GetOfficialAccount(&c)
}

func WechatConfig() *WehcatConfig {
	wen := &WehcatConfig{
		AppId:            viper.GetString("wechat.base.appid"),
		Appkey:           viper.GetString("wechat.base.appsecret"),
		MchId:            viper.GetString("wechat.pay.mchid"),
		SerialNo:         viper.GetString("wechat.pay.serialNo"),
		ApiKey3:          viper.GetString("wechat.pay.apiKey3"),
		ApiclientKeyPath: viper.GetString("wechat.pay.apiclientKeyPath"),
		ApiclientCerPath: viper.GetString("wechat.pay.apiclientCerPath"),
		ApiclientCer12Path: viper.GetString("wechat.pay.apiclientCer12Path"),
		IsProd:           viper.GetBool("wechat.pay.isProd"),
		NotifyUrl:        viper.GetString("wechat.pay.notifyUrl"),
	}
	return wen
}
// NewWechatClient /*
func NewWechatClient() WechatClient {
	if wepay==nil {
		wepay = WechatConfig()
	}
	if wechatClient==nil {
		 wechatClient=new(WechatService)
	}
	return wechatClient
}
