package ui

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/govcl/vcl/types/keys"
	"reflect"
	"strconv"
)

type ReqNum struct {
	Num int
}

type CheckState struct {
	host         bool
	ip           bool
	title        bool
	domain       bool
	port         bool
	country      bool
	province     bool
	city         bool
	country_name bool
	header       bool
	server       bool
	protocol     bool
	banner       bool
}

//::private::
type TForm4Fields struct {
	CheckState
	ReqNum
}

func (f *TForm4) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.EnabledMaximize(false)
	f.SetColor(0x241813)
	f.StaticText1.Font().SetColor(colors.ClWhite)
	f.CheckListBox1.SetColor(0xcfb593)

	//添加选项
	f.CheckListBox1.Items().Add(fmt.Sprintf("host"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("ip"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("title"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("domain"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("port"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("country"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("province"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("city"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("country_name"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("header"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("server"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("protocol"))
	f.CheckListBox1.Items().Add(fmt.Sprintf("banner"))

	//根据配置文件设置选项
	iniFile := vcl.NewIniFile(".\\config.ini")
	defer iniFile.Free()

	f.CheckState.host = iniFile.ReadBool("ReqInfo", "host", true)
	f.CheckState.ip = iniFile.ReadBool("ReqInfo", "ip", true)
	f.CheckState.title = iniFile.ReadBool("ReqInfo", "title", true)
	f.CheckState.domain = iniFile.ReadBool("ReqInfo", "domain", false)
	f.CheckState.port = iniFile.ReadBool("ReqInfo", "port", false)
	f.CheckState.country = iniFile.ReadBool("ReqInfo", "country", false)
	f.CheckState.province = iniFile.ReadBool("ReqInfo", "province", false)
	f.CheckState.city = iniFile.ReadBool("ReqInfo", "city", false)
	f.CheckState.country_name = iniFile.ReadBool("ReqInfo", "country_name", false)
	f.CheckState.header = iniFile.ReadBool("ReqInfo", "header", false)
	f.CheckState.server = iniFile.ReadBool("ReqInfo", "server", false)
	f.CheckState.protocol = iniFile.ReadBool("ReqInfo", "protocol", false)
	f.CheckState.banner = iniFile.ReadBool("ReqInfo", "banner", false)

	//读取请求数量
	f.ReqNum.Num = int(iniFile.ReadInteger("ReqNum", "Num", 100))
	f.Edit1.SetText(strconv.Itoa(f.ReqNum.Num))

	t := reflect.TypeOf(f.CheckState)
	v := reflect.ValueOf(f.CheckState)
	c := 0
	//v.Field(k).Bool() t.Field(k).Name
	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Bool() {
			f.CheckListBox1.SetChecked(int32(c), true)
			c = c + 1
		}
	}

	f.CheckListBox1.SetOnClick(func(sender vcl.IObject) {
		//改变状态
		//fmt.Println(f.CheckListBox1.ItemIndex(),f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex()))
		//如果选项选中则添加到CheckState
		switch f.CheckListBox1.ItemIndex() {
		case 0:
			f.CheckState.host = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 1:
			f.CheckState.ip = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 2:
			f.CheckState.title = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 3:
			f.CheckState.domain = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 4:
			f.CheckState.port = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 5:
			f.CheckState.country = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 6:
			f.CheckState.province = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 7:
			f.CheckState.city = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 8:
			f.CheckState.country_name = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 9:
			f.CheckState.header = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 10:
			f.CheckState.server = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 11:
			f.CheckState.protocol = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		case 12:
			f.CheckState.banner = f.CheckListBox1.Checked(f.CheckListBox1.ItemIndex())
		}
	})
	//按钮事件
	f.Button1.SetOnClick(f.OnConReqInfo)
	f.Button2.SetOnClick(f.OnClearReqInfo)
	f.Edit1.SetOnChange(f.OnReqInfoChange)

	//可以接收键盘操作
	f.SetKeyPreview(true)
	// 键盘弹起事件
	f.SetOnKeyUp(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == keys.VkEscape {
			f.Close()
		}
	})
}

//数量改变
func (f *TForm4) OnReqInfoChange(sender vcl.IObject) {
	val, _ := strconv.Atoi(f.Edit1.Text())
	f.ReqNum.Num = val
}

//确认选项
func (f *TForm4) OnConReqInfo(sender vcl.IObject) {
	t := reflect.TypeOf(f.CheckState)
	v := reflect.ValueOf(f.CheckState)
	//打开文件
	iniFile := vcl.NewIniFile(".\\config.ini")
	defer iniFile.Free()
	//将checkState写入config.ini
	for k := 0; k < t.NumField(); k++ {
		//fmt.Println("t.Field(k).Name",t.Field(k).Name,"v.Field(k).Bool()",v.Field(k).Bool())
		iniFile.WriteBool("ReqInfo", t.Field(k).Name, v.Field(k).Bool())
	}
	iniFile.WriteInteger("ReqNum", "Num", int32(f.ReqNum.Num))
	vcl.ShowMessage("设置成功！")
	//读取配置信息
	Form1.OnReadConfig()
	//根据请求参数改变listview
	Form1.OnChangeListView()
	Form4.Close()
	//Form1.Show()
}

//重置选项
func (f *TForm4) OnClearReqInfo(sender vcl.IObject) {
	//取消选项
	f.CheckListBox1.CheckAll(types.CbUnchecked, true, true)
	//取消CheckState状态
	f.CheckState.host = false
	f.CheckState.ip = false
	f.CheckState.title = false
	f.CheckState.domain = false
	f.CheckState.port = false
	f.CheckState.country = false
	f.CheckState.province = false
	f.CheckState.city = false
	f.CheckState.country_name = false
	f.CheckState.header = false
	f.CheckState.server = false
	f.CheckState.protocol = false
	f.CheckState.banner = false
}
