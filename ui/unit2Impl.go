package ui

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/govcl/vcl/types/keys"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type UserInfo struct {
	UserEmail string
	UserKey   string
	Vip       string
}

//用户转换检测用户信息失败json
type CheckUserErr struct {
	Error  bool   `json:"error"`
	Errmsg string `json:"errmsg"`
}

//用户转换检测用户信息成功json
type CheckUser struct {
	Email       string `json:"error"`
	Username    string `json:"username"`
	Fcoin       int    `json:"fcoin"`
	Isvip       bool   `json:"isvip"`
	Vip_level   int32  `json:"vip_level"`
	Is_verified bool   `json:"is_verified"`
	Avatar      string `json:"avatar"`
	Message     int    `json:"message"`
	Fofacli_ver string `json:"fofacli_ver"`
	Fofa_server bool   `json:"fofa_server"`
}

//::private::
type TForm2Fields struct {
	CheckUserErr
	CheckUser
	UserInfo
}

func (f *TForm2) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.EnabledMaximize(false)
	f.SetCaption("设置用户信息")
	f.Button1.SetOnClick(f.OnConfirmUserInfo)
	f.Button2.SetOnClick(f.OnClearUserInfo)
	f.StaticText1.SetColor(0x241813)
	f.StaticText2.SetColor(0x241813)
	f.StaticText1.Font().SetColor(colors.ClWhite)
	f.StaticText2.Font().SetColor(colors.ClWhite)
	f.SetColor(0x241813)
	iniFile := vcl.NewIniFile(".\\config.ini")
	defer iniFile.Free()

	e := iniFile.ReadString("UserInfo", "Email", "")
	p := iniFile.ReadString("UserInfo", "Key", "")

	f.Edit1.SetText(e)
	f.Edit2.SetText(p)
	//可以接收键盘操作
	f.SetKeyPreview(true)
	// 键盘弹起事件
	f.SetOnKeyUp(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == keys.VkEscape {
			f.Close()
		}
	})
}

func (f *TForm2) OnFormShow(sender vcl.IObject) {
	iniFile := vcl.NewIniFile(".\\config.ini")
	defer iniFile.Free()

	e := iniFile.ReadString("UserInfo", "Email", "")
	p := iniFile.ReadString("UserInfo", "Key", "")

	f.Edit1.SetText(e)
	f.Edit2.SetText(p)
}

//func (f *TForm2) OnFormShortCut(msg *types.TWMKey, handled *bool) {
//    fmt.Println("123")
//}
//监听email输入框
func (f *TForm2) OnEdit1Change(sender vcl.IObject) {
	val := f.Edit1.Text()
	if val != "" {
		f.UserInfo.UserEmail = val
	}
}

//监听key输入框
func (f *TForm2) OnEdit2Change(sender vcl.IObject) {
	val := f.Edit2.Text()
	if val != "" {
		f.UserInfo.UserKey = val
	}
}

//转换用户信息错误的json
func (f *TForm2) Str2jsonErr(msg string, checkUserErr CheckUserErr) {
	if err := json.Unmarshal([]byte(msg), &checkUserErr); err == nil {
		if checkUserErr.Error {
			vcl.ShowMessage(checkUserErr.Errmsg)
		}
	} else {
		fmt.Println(err)
	}
}

//转换用户信息的json
func (f *TForm2) Str2json(msg string, checkUser CheckUser) {
	if err := json.Unmarshal([]byte(msg), &checkUser); err == nil {
		switch checkUser.Vip_level {
		case 1:
			f.UserInfo.Vip = "普通会员"
			break
		case 2:
			f.UserInfo.Vip = "高级会员"
			break
		case 3:
			f.UserInfo.Vip = "企业会员"
			break
		}
		vcl.ShowMessage(fmt.Sprintf("验证成功,欢迎您,尊贵的%s:%s", f.UserInfo.Vip, checkUser.Username))
		Form2.Close()
		//Form1.Show()
		WriteUserInfo(f.UserInfo.UserEmail, f.UserInfo.UserKey, checkUser.Vip_level)
	} else {
		fmt.Println(err)
	}
}

//确认用户信息
//vip等级0 1 2 3
func (f *TForm2) OnConfirmUserInfo(sender vcl.IObject) {
	res := VerifyEmailFormat(f.UserInfo.UserEmail)
	if res {
		f.Button1.SetEnabled(false)
		apiCh := make(chan string)
		go f.ConfirmUserInfoReq(apiCh)
		go f.ConfirmUserInfoResp(apiCh)
		//fmt.Println(string(bytes))

	} else {
		vcl.ShowMessage("邮箱格式错误！")
	}
}
func (f *TForm2) ConfirmUserInfoReq(apiCh chan string) {
	Url := fmt.Sprintf("https://fofa.so/api/v1/info/my?email=%s&key=%s", f.UserInfo.UserEmail, f.UserInfo.UserKey)
	//WriteUserInfo(f.UserInfo.UserEmail,f.UserInfo.UserKey)
	// 构造 http client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", Url, nil)

	response, err := client.Do(req)
	if err != nil {
		errmsg := "http.Get err=" + err.Error()
		apiCh <- errmsg
		return
	}
	//获取请求内容
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errmsg := "ioutil.ReadAll err=" + err.Error()
		apiCh <- errmsg
		return
	}
	apiCh <- string(bytes)

}
func (f *TForm2) ConfirmUserInfoResp(apiCh chan string) {
	respMsg := <-apiCh
	vcl.ThreadSync(func() {
		if strings.Contains(respMsg, "err=") {
			vcl.ShowMessage(respMsg)
		} else if strings.Contains(respMsg, "errmsg") {
			f.Str2jsonErr(respMsg, f.CheckUserErr)
		} else {
			f.Str2json(respMsg, f.CheckUser)
		}
		f.Button1.SetEnabled(true)
	})
	close(apiCh)
}

//重置用户信息
func (f *TForm2) OnClearUserInfo(sender vcl.IObject) {
	f.Edit1.SetText("")
	f.Edit2.SetText("")
}

//验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//写入email和key
func WriteUserInfo(e string, k string, v int32) {
	//创建文件
	iniFile := vcl.NewIniFile(".\\config.ini")
	defer iniFile.Free()

	iniFile.WriteString("UserInfo", "Email", e)
	iniFile.WriteString("UserInfo", "Key", k)
	iniFile.WriteInteger("UserInfo", "VipLevel", v)

	Form2.Close()
	//Form1.Show()
}
