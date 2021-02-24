package alipay

import (
	"fmt"
	"github.com/ooncn/common/util"
	"github.com/vmihailenco/msgpack"
)

const (
	GATEWAY = "https://openapi.alipay.com/gateway.do" //网关
	//会员API
	UserInfoShare = "alipay.user.info.share" //支付宝会员授权信息查询接口
	// 工具类
	UserInfoAuth = "alipay.user.info.auth" //(用户登陆授权)
	RSA2         = "RSA2"
)

type Config struct {
	AppId              string //AppId
	MerchantPrivateKey string //私钥
	PublicKey          string //请填写您的支付宝公钥
	///注：如果采用非证书模式，则无需赋值上面的三个证书路径，改为赋值如下的支付宝公钥字符串即可
	MerchantCertPath string //应用公钥证书文件路径
	CertPath         string //请填写您的支付宝公钥证书文件路径
	RootCertPath     string //请填写您的支付宝根证书文件路径
	NotifyUrl        string //可设置异步通知接收服务地址（可选）
	SignType         string //加密方式
}

func GetConfigOptions() *Config {
	return nil
}

var (
	Client               *AopClient
	AlipayConfigFileName = "alipayConfig.o"
)

func NewAopClient() *AopClient {
	str, err := util.ReaderFile(AlipayConfigFileName)
	if err != nil {
		fmt.Println("数据库配置文件获取失败" + err.Error())
		return nil
	}
	var set AopClient
	//err = msgpack.Unmarshal([]byte(util.PuDecrypt(str)), &set)
	err = msgpack.Unmarshal([]byte(str), &set) /*
		timeNow := time.Now().Unix()
		if timeNow >= set.AuthTime {
			fmt.Println("system setting error")
			return nil
		}*/
	return &set
}
func NewAopClientFile(file string) *AopClient {
	str, err := util.ReaderFile(file + "_" + AlipayConfigFileName)
	if err != nil {
		fmt.Println("数据库配置文件获取失败" + err.Error())
		return nil
	}
	var set AopClient
	//err = msgpack.Unmarshal([]byte(util.PuDecrypt(str)), &set)
	err = msgpack.Unmarshal([]byte(str), &set) /*
		timeNow := time.Now().Unix()
		if timeNow >= set.AuthTime {
			fmt.Println("system setting error")
			return nil
		}*/
	return &set
}

/*func init() {
	Client = NewAopClient()
	if Client==nil {
		return
	}
	r, err := util.NewRsa2PKCS1(Client.RsaPublicKey, Client.RsaPrivateKey)
	if err != nil {
		fmt.Println("支付宝加密组件失败" + err.Error())
	}
	Ras2Util = r
}*/
