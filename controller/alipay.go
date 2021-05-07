package controller

import (
	"go-pangu/conf"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
)

func AliPayHandler(c *gin.Context) {
	privateKey := conf.GetEnv("ALIPAY_PRIVATE_KEY") //私钥
	appId := conf.GetEnv("ALIPAY_APP_ID")           //appid

	//启动client
	client := alipay.NewClient(appId, privateKey, true)

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	client.
		SetPrivateKeyType(alipay.PKCS8).                         // 设置 支付宝 私钥类型，alipay.PKCS1 或 alipay.PKCS8，默认 PKCS1
		SetCharset("utf-8").                                     // 设置字符编码，不设置默认 utf-8
		SetSignType(alipay.RSA2).                                // 设置签名类型，不设置默认 RSA2
		SetReturnUrl("https://www.baidu.com").                   // 设置返回URL，支付成功后跳转
		SetNotifyUrl("http://xxxxxxxxxx.ngrok.io/alipay_notify") //设置回调地址，应该为一个公网ip地址

	//证书设置，不需要修改
	client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	bm := make(gopay.BodyMap)
	bm.Set("subject", "支付测试")                                //商品名字会显示在客户支付界面
	bm.Set("out_trade_no", RandStringRunes(16, LetterRunes)) //系统自设订单号，并非支付宝的账单号，但一样会返回
	bm.Set("quit_url", "https://www.baidu.com")              //中途退出返回地址
	bm.Set("total_amount", "0.01")                           //金额，以元为单位
	bm.Set("product_code", "QUICK_WAP_WAY")                  //销售产品码，商家和支付宝签约的产品码

	//发送请求，获取返回支付的url，打开就可以支付
	payURL, err := client.TradePagePay(bm)
	if err != nil {
		StatusError(c, http.StatusBadRequest, "fail", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "url": payURL})
}

func AliPayNotifyHandler(c *gin.Context) {
	//支付宝支付通知接收接口
	aliPayPublicKey := conf.GetEnv("ALIPAY_PUBLIC_KEY") //公钥

	//获取返回内容
	req, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		StatusError(c, http.StatusBadRequest, "fail", err.Error())
		return
	}

	//验证发来签名是否合法
	ok, err := alipay.VerifySign(aliPayPublicKey, req)
	if err != nil || !ok {
		StatusError(c, http.StatusBadRequest, "fail", "sign error")
		return
	}

	//获取参数
	tradeStatus := req.Get("trade_status")
	if tradeStatus != "TRADE_SUCCESS" {
		StatusError(c, http.StatusBadRequest, "fail", "pay fail")
		return
	}

	//固定设置返回的值
	c.String(http.StatusOK, "%s", "success")
}
