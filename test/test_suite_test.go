package test_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-pangu/conf"
	"go-pangu/db"
	"go-pangu/jwt"
	"go-pangu/models"
	"go-pangu/routers"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
	"unsafe"

	"github.com/google/uuid"
	"github.com/imroc/req"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var Ts *httptest.Server

var token string
var accessToken string

var Alphanumeric = []rune("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

var ROCKETCHAT_PASS = conf.GetEnv("ROCKETCHAT_PASS")

var UNMISTAKABLE_CHARS = []rune("23456789ABCDEFGHJKLMNPQRSTWXYZabcdefghijkmnopqrstuvwxyz")

func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

func RandStringRunes(n int, runename []rune) string {
	var b = make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = runename[rand.Intn(len(runename))]
	}
	return string(b)
}

func Mergemap(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

func Params(method, api string, params url.Values) *http.Response {
	req, _ := http.NewRequest(method, Ts.URL+"/"+api, nil)
	if len(params) > 0 {
		req.URL.RawQuery = params.Encode()
	}
	resp, _ := http.DefaultClient.Do(req)
	return resp
}

func headers(token string) req.Header {
	return req.Header{
		"Accept":        "application/json",
		"Authorization": token,
	}
}

func Get(body interface{}, s, token string) *req.Resp {
	resp, _ := req.Get(Ts.URL+s, headers(token), req.BodyJSON(&body))
	return resp
}

func Post(body interface{}, s, token string) *req.Resp {
	resp, _ := req.Post(Ts.URL+s, headers(token), req.BodyJSON(&body))
	return resp
}

func Put(body interface{}, s, token string) *req.Resp {
	resp, _ := req.Put(Ts.URL+s, headers(token), req.BodyJSON(&body))
	return resp
}

func Delete(body interface{}, s, token string) *req.Resp {
	resp, _ := req.Delete(Ts.URL+s, headers(token), req.BodyJSON(&body))
	return resp
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BoxIn() *gorm.DB {
	tmp := db.DB
	if db.DB != nil {
		fmt.Println("formal", tmp)
		db.DB = db.DB.Begin()
	}
	return tmp
}

func BoxOut(dbCache *gorm.DB) {
	if db.DB != nil {
		db.DB.Rollback()
	}
	db.DB = dbCache
}

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

func genNewUserToken(deviceType string) (models.User, string) {
	var user models.User
	id := uuid.New()
	user.ID = id

	db.DB.Create(&user)
	payload := jwt.GenPayload("", "user", id.String())
	token := jwt.Encoder(payload)
	token = "Bearer " + token
	return user, token
}

var _ = BeforeSuite(func() {
	db.Open("test")
	fmt.Println("Server starting ...")
	db.Migrate("test", &models.User{})
	// models.Seed()
	Ts = httptest.NewServer(routers.SetupRouter())
})

var _ = AfterSuite(func() {
	db.Close()
	db.CleanTablesData()
})
