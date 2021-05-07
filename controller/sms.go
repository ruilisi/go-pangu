package controller

import (
	"fmt"
	"go-pangu/conf"
	"go-pangu/redis"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qinxin0720/QcloudSms-go/QcloudSms"
)

//随机数组，生成随机数需要选一个组，随机的字符在其中选择
var (
	LetterRunes    = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	SMSletterRunes = []rune("0123456789")
	smserror       string
)

func callback(err error, resp *http.Response, resData string) {
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		smserror = resData
		fmt.Println("response data: ", resData)
	}
}

func RandStringRunes(n int, letterRunes []rune) string {
	//生成6位随机数
	//需要一个int类型参数 n
	//生成一个n位随机数转换成string类型返回
	var (
		//	letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b = make([]rune, n)
	)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SMSFunction(mobile string, message []string, templID int) error {
	//发送短信  使用的是腾讯云短信服务  需要在根目录的配置文件自行配置你的参数
	var (
		appID, _ = strconv.Atoi(conf.GetEnv("SMS_APPID"))
		appKey   = conf.GetEnv("SMS_APPKEY")
	)

	//单发短信
	qcloudsms, err := QcloudSms.NewQcloudSms(appID, appKey)
	if err != nil {
		panic(err)
	}
	err = qcloudsms.SmsSingleSender.SendWithParam(86, mobile, templID, message, "瑞立思科技", "", "", callback)
	return err
}

func SMSHandler(c *gin.Context) {
	//获取手机号，发送随机生成的验证码
	var (
		mobile     map[string]interface{}
		SMSCode    string
		templID, _ = strconv.Atoi(conf.GetEnv("SMS_TEMPID"))
	)

	//绑定
	if err := c.ShouldBind(&mobile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//生成随机数 保存验证码在redis
	mobilenum := mobile["mobile"].(string)
	SMSCode = RandStringRunes(6, SMSletterRunes)
	redis.SetEx(mobilenum, SMSCode, time.Duration(5)*time.Minute)
	message := []string{SMSCode, "5"}
	err := SMSFunction(mobilenum, message, templID)

	//返回结果
	if err != nil {
		StatusError(c, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "send success",
	})
}

func SMSVerify(code string, mobile string) bool {
	//验证手机号 在需要验证用户提交的code是否与redis数据库一致时使用
	smsCode := redis.Get(mobile)
	if code == string(smsCode) {
		return true
	}
	return false
}
