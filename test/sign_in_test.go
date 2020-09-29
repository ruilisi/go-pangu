package test_test

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SignIn", func() {
	result := map[string]interface{}{}
	rand := RandStringRunes(8, Alphanumeric)

	BeforeEach(func() {
		result = map[string]interface{}{}
	})

	Context("signup ", func() {
		It("should show signup success", func() {
			signupParams := map[string]string{
				"email":            rand + "@te.com",
				"password":         "123456",
				"password_confirm": "123456",
				"signup_type":      "email",
			}
			resp := Post(signupParams, "/users/sign_up", "")
			resp.ToJSON(&result)
			fmt.Println(resp)
			Expect(resp.Response().StatusCode).To(Equal(http.StatusOK))
			Ω(fmt.Sprint(result["status"])).To(Equal("register success"))
		})
	})

	Context("signin ", func() {
		It("should show signin success", func() {
			signinParams := map[string]string{
				"email":       rand + "@te.com",
				"password":    "123456",
				"DEVICE_TYPE": "MAC",
				"login_type":  "email",
			}
			resp := Post(signinParams, "/users/sign_in", "")
			head := resp.Response().Header

			bear := head.Get("Authorization")
			token = strings.Replace(bear, "Bearer ", "", 1)

			resp.ToJSON(&result)
			Expect(resp.Response().StatusCode).To(Equal(http.StatusOK))
			Ω(fmt.Sprint(result["status"])).To(Equal("login success"))
		})
	})

	Context("change password", func() {
		It("should change password success", func() {
			changePassParams := map[string]string{
				"password":         "112233",
				"password_confirm": "112233",
				"origin_password":  "123456",
			}
			resp := Post(changePassParams, "/users/change_password", token)

			resp.ToJSON(&result)
			Expect(resp.Response().StatusCode).To(Equal(http.StatusOK))
			Ω(fmt.Sprint(result["status"])).To(Equal("update password success"))
		})
	})

	Context("after change password signin ", func() {
		It("should show signin success", func() {
			signinParams := map[string]string{
				"email":       rand + "@te.com",
				"password":    "112233",
				"DEVICE_TYPE": "MAC",
				"login_type":  "email",
			}
			resp := Post(signinParams, "/users/sign_in", "")
			head := resp.Response().Header

			bear := head.Get("Authorization")
			token = strings.Replace(bear, "Bearer ", "", 1)

			resp.ToJSON(&result)
			Expect(resp.Response().StatusCode).To(Equal(http.StatusOK))
			Ω(fmt.Sprint(result["status"])).To(Equal("login success"))
		})
	})

})
