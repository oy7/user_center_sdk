package test

import (
	"testing"

	"github.com/oy7/user_center_sdk/proto/user"
	"github.com/oy7/user_center_sdk/proxy"
)

const (
	url       = "192.168.254.93:9521"
	source    = "user_center_sdk"
	requestId = ""
)

// 测试获取用户信息
func TestGetBaseInfo(t *testing.T) {
	cases := []struct {
		Name string
		Uid  uint64
	}{
		{"十里", 150067},
		{"艾客", 150052},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			uc := proxy.Init(url, source, requestId)
			resp, err := uc.GetBaseInfo(c.Uid)
			if err != nil {
				t.Error(err)
			}
			if resp.Code != 0 {
				t.Error(resp.Msg)
			}
			t.Logf("resp: %+v", resp)
		})
	}

}

// 测试发送验证码
func TestSmsSendLogin(t *testing.T) {
	cases := []struct {
		Name string
		Req  *user.SMSSendLoginReq
	}{
		{"登录-十里", &user.SMSSendLoginReq{
			PhoneNumber: "18038024878",
			SmsCodeType: 2,
		}},
		{"登录-艾客", &user.SMSSendLoginReq{
			PhoneNumber: "15302738153",
			SmsCodeType: 2,
		}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			uc := proxy.Init(url, source, requestId)
			resp, err := uc.ServiceSmsSendLogin(c.Req.PhoneNumber, c.Req.SmsCodeType)
			if err != nil {
				t.Error(err)
			}
			if resp.Code != 0 {
				t.Error(resp.Msg)
			}
			t.Logf("resp: %+v", resp)
		})
	}

}

// 测试用户登录
func TestUserLogin(t *testing.T) {
	cases := []struct {
		Name string
		Req  *user.UserLoginReq
	}{
		{"短信登录-十里", &user.UserLoginReq{
			LoginType:   "sms",
			PhoneNumber: "18038024878",
			SmsCode:     "2719",
		}},
		{"短信登录-艾客", &user.UserLoginReq{
			LoginType:   "sms",
			PhoneNumber: "15302738153",
			SmsCode:     "9712",
		}},
		{"一键登录-十里", &user.UserLoginReq{
			LoginType:   "one_click",
			PhoneNumber: "18038024878",
		}},
		{"一键登录-艾客", &user.UserLoginReq{
			LoginType:   "one_click",
			PhoneNumber: "15302738153",
		}},
		{"密码登录-十里 ", &user.UserLoginReq{
			LoginType:   "password",
			PhoneNumber: "18038024878",
			Password:    "4878",
		}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			uc := proxy.Init(url, source, requestId)
			resp, err := uc.ApiUserLogin(c.Req)
			if err != nil {
				t.Error(err)
			}
			t.Logf("resp: %+v", resp)
		})
	}
}
