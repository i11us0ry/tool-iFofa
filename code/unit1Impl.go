package code

import (
    "encoding/base64"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "github.com/ying32/govcl/vcl/win"
    "io/ioutil"
    "math"
    "net/http"
    "os"
    "reflect"
    "strconv"
    "strings"
)
type ResultErr struct {
    Errmsg      string   `json:"errmsg"`
    Error       bool     `json:"error"`
}

type ResultJson1 struct {
   Mode        string      `json:"mode"`
   Error       bool        `json:"error"`
   Query       string      `json:"query"`
   Page        int         `json:"page"`
   Size        int64       `json:"size"`
   Results  [] string      `json:"results"`
}

type ResultJson2 struct {
    Mode        string      `json:"mode"`
    Error       bool        `json:"error"`
    Query       string      `json:"query"`
    Page        int         `json:"page"`
    Size        int64       `json:"size"`
    Results [][] string     `json:"results"`
}

type InKey struct {
    key   string
}

type ReqNum1 struct {
    Num     int32   //需要请求的数量
    NumMax  int32   //最大请求书数
}

type UserInfo1 struct {
    Email       string
    Key         string
    VipLevel    int32
}

type CheckState1 struct {
    host            bool
    ip              bool
    title           bool
    domain          bool
    port            bool
    country         bool
    province        bool
    city            bool
    country_name    bool
    header          bool
    server          bool
    protocol        bool
    banner          bool
}

//::private::
type TForm1Fields struct {
    inkey       InKey
    Num         int32
    ResultErr
    ResultJson1
    ResultJson2
    Page        int
    MaxPage     int64
    MaxSize     int64
    C           int
    Fields      string
    Datas  []   string
    subItemHit  win.TLVHitTestInfo
}

var resultErr   ResultErr
var resultJson1 ResultJson1
var resultJson2 ResultJson2

func (f *TForm1) OnFormCreate(sender vcl.IObject){
    f.SetCaption("iFofa                                                                                             by:  i11us0ry")
    //不允许最大化
    f.EnabledMaximize(false)
    //定位屏幕中央
    f.ScreenCenter()
    //主题颜色
    f.SetColor(0x241813)

    //设置用户信息
    f.MenuItem1.SetOnClick(f.OnsetUserInfo)
    //设置请求参数
    f.MenuItem2.SetOnClick(f.OnsetReqInfo)
    //查看参考语法
    f.MenuItem3.SetOnClick(f.OngetHelpInfo)

    //默认第一页
    f.Page = 1

    //listview初始化
    f.ListView1.SetColor(0xcfb593)
    f.ListView1.SetRowSelect(true)
    f.ListView1.SetOnClick(f.OnListView1Click)
    col := f.ListView1.Columns().Add()
    col.SetCaption("序号")
    col.SetWidth(50)
    //读取配置文件
    f.OnReadConfig()
    //填充listview columns
    f.OnChangeListView()

    //右键功能事件
    item := vcl.NewPopupMenu(f)
    //上一页功能，会先判断当前是否为第一页！
    subMenu := vcl.NewMenuItem(f)
    subMenu.SetCaption("上一页(&U)")
    subMenu.SetShortCutFromString("Ctrl+U")
    subMenu.SetOnClick(func(sender vcl.IObject) {
        f.OnGoToPage(1)
    })
    //下一页功能，根据更改Page来达到下一页的请求
    subMenu1 := vcl.NewMenuItem(f)
    subMenu1.SetCaption("下一页(&U)")
    subMenu1.SetShortCutFromString("Ctrl+D")
    subMenu1.SetOnClick(func(sender vcl.IObject) {
        f.OnGoToPage(2)
    })
    //首页，会先判断当前页是否为第一页，然后将Page设为1去请求数据
    subMenu2 := vcl.NewMenuItem(f)
    subMenu2.SetCaption("首页(&F)")
    subMenu2.SetShortCutFromString("Ctrl+F")
    subMenu2.SetOnClick(func(sender vcl.IObject) {
        f.OnGoToPage(0)
    })
    //首页，会先判断当前页是否为第一页，然后将Page设为1去请求数据
    subMenu6 := vcl.NewMenuItem(f)
    subMenu6.SetCaption("尾页(&L)")
    subMenu6.SetShortCutFromString("Ctrl+L")
    subMenu6.SetOnClick(func(sender vcl.IObject) {
        f.OnGoToPage(3)
    })
    //退出程序
    subMenu3 := vcl.NewMenuItem(f)
    subMenu3.SetCaption("退出(&Q)")
    subMenu3.SetShortCutFromString("Ctrl+Q")
    subMenu3.SetOnClick(func(vcl.IObject) {
        f.Close()
    })
    //保存数据
    subMenu4 := vcl.NewMenuItem(f)
    subMenu4.SetCaption("保存(&S)")
    subMenu4.SetShortCutFromString("Ctrl+S")
    subMenu4.SetOnClick(func(sender vcl.IObject) {
        f.OnSaveFile()
    })
    //情况当前数据
    subMenu5 := vcl.NewMenuItem(f)
    subMenu5.SetCaption("清空(&C)")
    subMenu5.SetShortCutFromString("Ctrl+C")
    subMenu5.SetOnClick(func(object vcl.IObject){
        if f.ListView1.Items().Count() == 0{
            vcl.ShowMessage("当前没有数据哦！")
        } else{
            f.ListView1.Clear()
        }
    })
    item.Items().Add(subMenu2)
    item.Items().Add(subMenu)
    item.Items().Add(subMenu1)
    item.Items().Add(subMenu6)
    item.Items().Add(subMenu4)
    item.Items().Add(subMenu5)
    item.Items().Add(subMenu3)
    f.SetPopupMenu(item)
}

//监听搜索语法输入框
func (f *TForm1) OnEdit1Change(sender vcl.IObject) {
    val := f.Edit1.Text()
    if val !="" {
        f.inkey.key = val
    }
}

//用户信息设置
func (f *TForm1) OnsetUserInfo(sender vcl.IObject) {
    Form2.Show()
}

//设置请求信息
func (f *TForm1) OnsetReqInfo(sender vcl.IObject) {
    Form5.Show()
}

//查看请求参数
func (f *TForm1) OngetHelpInfo(sender vcl.IObject) {
    Form3.Show()
}

//确认按钮
func (f *TForm1) OnButton3Click(sender vcl.IObject) {
    //判断是否输入搜索搜索语法
    if f.Edit1.Text() == ""{
        vcl.ShowMessage("请输入搜索语法！")
    } else{
        //重置listview数据
        f.ListView1.Clear()
        //重置访问页数！
        f.Page = 1
        //读取配置信息
        f.OnReadConfig()
        //根据请求参数改变listview
        f.OnChangeListView()
        //将请求参数写入Fields
        f.Fields = ""
        t := reflect.TypeOf(f.CheckState1)
        v := reflect.ValueOf(f.CheckState1)
        f.C = 0
        for k := 0;k<t.NumField();k++{
            //iniFile.WriteBool("ReqInfo", t.Field(k).Name, v.Field(k).Bool())
            if v.Field(k).Bool(){
                f.C =f.C + 1
                f.Fields = f.Fields + "," + t.Field(k).Name
            }
        }
        f.Fields = strings.Replace(f.Fields,",","",1)
        f.OnReqData(f)
    }
}

//根据请求参数改变listview
func (f *TForm1) OnChangeListView(){
    t := reflect.TypeOf(f.CheckState1)
    v := reflect.ValueOf(f.CheckState1)
    c := 0
    //v.Field(k).Bool() t.Field(k).Name
    for k := 0;k<t.NumField();k++{
        if v.Field(k).Bool(){
            c = c + 1
            if c <= int(f.ListView1.Columns().Count()-1){
                f.ListView1.Columns().Items(int32(c)).SetCaption(t.Field(k).Name)
                f.ListView1.Columns().Items(int32(c)).SetWidth(150)
            } else{
               f.ListView1.Columns().Add()
               f.ListView1.Columns().Items(int32(c)).SetCaption(t.Field(k).Name)
               f.ListView1.Columns().Items(int32(c)).SetWidth(150)
            }
            //fmt.Println("count:",f.ListView1.Columns().Count()-1," c:",c,"caption:",f.ListView1.Columns().Items(int32(c)).Caption())
        } else {
           for i := 1;i<= int(f.ListView1.Columns().Count()-1);i++{
               Cap := f.ListView1.Columns().Items(int32(i)).Caption()
               //fmt.Println("Cap:",Cap," t.Field(k).Name:",t.Field(k).Name)
               if Cap == t.Field(k).Name{
                   f.ListView1.Columns().Delete(int32(i))
               }
           }
        }
    }
}

//读取配置文件
func (f *TForm1) OnReadConfig(){
    iniFile := vcl.NewIniFile(".\\config.ini")
    defer iniFile.Free()

    //读取用户信息
    f.UserInfo1.Email       = iniFile.ReadString("UserInfo", "Email", "")
    f.UserInfo1.Key         = iniFile.ReadString("UserInfo", "Key", "")
    f.UserInfo1.VipLevel    = iniFile.ReadInteger("UserInfo","VipLevel",0)
    switch f.UserInfo1.VipLevel {
    //用户可查询最大数量
    case 0:
    case 1:
        f.ReqNum1.NumMax = 100
        break
    case 2:
        f.ReqNum1.NumMax = 10000
        break
    case 3:
        f.ReqNum1.NumMax = 100000
        break
    }

    //读取请求数量
    f.ReqNum1.Num = iniFile.ReadInteger("ReqNum", "Num", 100)

    //从config.ini读取请求参数
    f.CheckState1.host = iniFile.ReadBool("ReqInfo", "host", true)
    f.CheckState1.ip = iniFile.ReadBool("ReqInfo", "ip", true)
    f.CheckState1.title = iniFile.ReadBool("ReqInfo", "title", true)
    f.CheckState1.domain = iniFile.ReadBool("ReqInfo", "domain", false)
    f.CheckState1.port = iniFile.ReadBool("ReqInfo", "port", false)
    f.CheckState1.country = iniFile.ReadBool("ReqInfo", "country", false)
    f.CheckState1.province = iniFile.ReadBool("ReqInfo", "province", false)
    f.CheckState1.city = iniFile.ReadBool("ReqInfo", "city", false)
    f.CheckState1.country_name = iniFile.ReadBool("ReqInfo", "country_name", false)
    f.CheckState1.header = iniFile.ReadBool("ReqInfo", "header", false)
    f.CheckState1.server = iniFile.ReadBool("ReqInfo", "server", false)
    f.CheckState1.protocol = iniFile.ReadBool("ReqInfo", "protocol", false)
    f.CheckState1.banner = iniFile.ReadBool("ReqInfo", "banner", false)
}

//results为一维数组时
func (f *TForm1) str2json1(msg string,resultJson ResultJson1){
    if err := json.Unmarshal([]byte(msg), &resultJson); err == nil {
        if resultJson.Error{
            vcl.ShowMessage(f.ResultErr.Errmsg)
        }else{
            f.MaxSize = resultJson.Size
            f.WriteListview1(resultJson.Results)
        }
    } else {
        fmt.Println(err)
    }
}

//results为二维数组时
func (f *TForm1) str2json2(msg string,resultJson ResultJson2){
    if err := json.Unmarshal([]byte(msg), &resultJson); err == nil {
        if resultJson.Error{
            vcl.ShowMessage(f.ResultErr.Errmsg)
        }else{
            f.MaxSize = resultJson.Size
            f.WriteListview2(resultJson.Results)
        }
    } else {
        fmt.Println(err)
    }
}

//results错误
func (f *TForm1) str2json3(msg string,resultErr ResultErr){
    if err := json.Unmarshal([]byte(msg), &resultErr); err == nil {
        if resultErr.Error{
            vcl.ShowMessage(resultErr.Errmsg)
        }
    } else{
        fmt.Println(err)
    }
}

//results为一维数组时填充数据
func (f *TForm1) WriteListview1(results []string){
    f.ListView1.Clear()
    f.ListView1.Items().BeginUpdate()
    for i:=0;i< len(results);i++{
        item := f.ListView1.Items().Add()
        item.SetCaption(fmt.Sprintf(strconv.Itoa(i+1)))
        item.SubItems().Add(fmt.Sprintf(results[i]))
    }
    f.ListView1.Items().EndUpdate()
}

//results为多维数组时填充数据
func (f *TForm1) WriteListview2(results [][]string){
    //填充前先清除原有的数据
    f.ListView1.Clear()
    f.ListView1.Items().BeginUpdate()
    for i:=0;i<len(results);i++{
        item := f.ListView1.Items().Add()
        item.SetCaption(fmt.Sprintf(strconv.Itoa(i+1)))
        for j:=0;j<len(results[i]);j++{
            item.SubItems().Add(fmt.Sprintf(results[i][j]))
        }
    }
    f.ListView1.Items().EndUpdate()
}

//双击选项事件
func (f *TForm1) OnListView1DblClick(sender vcl.IObject) {
    p := f.ListView1.ScreenToClient(vcl.Mouse.CursorPos())
    f.subItemHit.Pt.X = p.X
    f.subItemHit.Pt.Y = p.Y
    win.ListView_SubItemHitTest(f.ListView1.Handle(), &f.subItemHit)
    if f.subItemHit.IItem != -1 {
       var r types.TRect
       if f.ListView1.RowSelect() {
           r = f.ListView1.Selected().DisplayRect(types.DrBounds)
       } else {
           win.ListView_GetItemRect(f.ListView1.Handle(), f.subItemHit.IItem, &r, 0)
       }
       colWidht := f.ListView1.Column(f.subItemHit.ISubItem).Width()

       var left, i int32
       // 差10像素
       left += 10
       for i = 0; i < f.subItemHit.ISubItem; i++ {
           left += f.ListView1.Column(i).Width() //ListView_GetColumnWidth(f.ListView1.Handle(), i)
       }

       if f.subItemHit.ISubItem!=0{
           f.Edit2.SetText(f.ListView1.Items().Item(f.subItemHit.IItem).SubItems().Strings(f.subItemHit.ISubItem - 1))
           f.Edit2.SetBounds(left, f.ListView1.Top()+r.Top, colWidht, r.Bottom-r.Top)
           f.Edit2.Show()
           f.Edit2.SetFocus()
       }
    }
}

//点击事件，取消双击长生的选择框
func (f *TForm1) OnListView1Click(sender vcl.IObject){
    f.Edit2.Hide()
}

//请求数据
func (f *TForm1) OnReqData(sender vcl.IObject){
    //对请求语法进行base64编码
    kbase64 := base64.StdEncoding.EncodeToString([]byte(f.inkey.key))
    //生成url
    Url := fmt.Sprintf("https://fofa.so/api/v1/search/all?email=%s&key=%s&qbase64=%s&page=%d&size=%d&fields=%s",
        f.UserInfo1.Email,f.UserInfo1.Key,kbase64,f.Page,int(f.ReqNum1.Num),f.Fields)
    //获取http请求状态
    response,err:= http.Get(Url)
    if err != nil{
        fmt.Println("http.Get err=",err)
        return
    }
    //获取请求内容
    bytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("ioutil.ReadAll err=",err)
        return
    }
    if strings.Contains(string(bytes), "errmsg"){
        f.str2json3(string(bytes),resultErr)
    }else{
        //数据请求成功时记录已经请求的数量
        if f.C==1 {
            //results为一维数组时
            f.str2json1(string(bytes),resultJson1)
        } else{
            //results为二维数组时
            f.str2json2(string(bytes),resultJson2)
        }
    }
}

//页数跳转
func (f *TForm1) OnGoToPage(IsNext int) {
    switch  IsNext {
        case 0://首页
            fmt.Println("点击了首页！")
            if f.Page == 1{
                vcl.ShowMessage("这就是首页啊！")
            } else{
                if f.ListView1.Items().Count() == 0{
                    vcl.ShowMessage("请先实现一次数据请求！")
                } else{
                    f.Page = 1
                    f.OnReqData(f)
                }
            }
            break
        case 1://上一页
            fmt.Println("点击了上一页！")
            if f.Page == 1{
                vcl.ShowMessage("这已经是第一页了呀！")
            } else{
                if f.ListView1.Items().Count() == 0{
                    vcl.ShowMessage("请先实现一次数据请求！")
                } else{
                    f.Page = f.Page - 1
                    f.OnReqData(f)
                }
            }
            break
        case 2://下一页
            fmt.Println("点击了下一页！")
            if f.ListView1.Items().Count() == 0{
                vcl.ShowMessage("请先实现一次数据请求！")
            } else{
                //如果资产数f.MaxSize大于可搜索最大数f.ReqNum1.NumMax，则最大页lastPage可搜索最大数f.ReqNum1.NumMax/单次搜索数f.ReqNum1.Num
                //反之若f.MaxSize小于f.ReqNum1.NumMax，则lastPage=f.MaxSize/f.ReqNum1.Num
                var page float64
                var lastPage int
                if f.MaxSize>int64(f.ReqNum1.NumMax){
                    page = float64(f.ReqNum1.NumMax)/float64(f.ReqNum1.Num)
                    lastPage = int(math.Floor(page))
                }else{
                    page = float64(f.MaxSize)/float64(f.ReqNum1.Num)
                    lastPage = int(math.Ceil(page))
                }
                if f.Page == lastPage{
                    vcl.ShowMessage(fmt.Sprintf("当前会员等级不支持再查看%d条数据",f.ReqNum1.Num))
                } else{
                    f.Page = f.Page + 1
                    f.OnReqData(f)
                }
            }
            break
        case 3://末页
            fmt.Println("点击了尾页！")
            if f.ListView1.Items().Count() == 0{
                vcl.ShowMessage("请先实现一次数据请求！")
            } else{
                var page float64
                var lastPage int
                if f.MaxSize>int64(f.ReqNum1.NumMax){
                    page = float64(f.ReqNum1.NumMax)/float64(f.ReqNum1.Num)
                    lastPage = int(math.Floor(page))
                }else{
                    page = float64(f.MaxSize)/float64(f.ReqNum1.Num)
                    lastPage = int(math.Ceil(page))
                }
                if f.Page == lastPage{
                    vcl.ShowMessage(fmt.Sprintf("当前会员等级不支持再查看%d条数据！",f.ReqNum1.Num))
                } else{
                    f.Page = lastPage
                    f.OnReqData(f)
                }
            }
    }
}

//文件保存
func (f *TForm1) OnSaveFile(){
    var h int32 //头序号
    var i int32 //行序号
    var j int32 //列序号
    //调用系统文件保存接口
    dlSave := vcl.NewSaveDialog(f)
    dlSave.SetFilter("文本文件(*.csv)|*.csv")
    dlSave.SetOptions(dlSave.Options().Include(types.OfShowHelp))
    dlSave.SetTitle("保存")

    if f.ListView1.Items().Count() == 0{
        vcl.ShowMessage("当前没有数据哦！")
    } else{
        if dlSave.Execute() {
            //创建用户选择的文件
            file, err := os.OpenFile(dlSave.FileName(), os.O_CREATE|os.O_RDWR, 0644)
            if err != nil {
                fmt.Println("open file is failed, err: ", err)
            }
            defer file.Close()
            // 写入UTF-8 BOM，防止中文乱码
            file.WriteString("\xEF\xBB\xBF")
            writer := csv.NewWriter(file)
            //写入标题
            Header := make([]string,0)
            for h=1;h<f.ListView1.Columns().Count();h++{
                Header = append(Header,f.ListView1.Column(h).Caption())
            }
            writer.Write(Header)
            //读取行数
            for i=0;i<f.ListView1.Items().Count();i++{
                ItemsData := make([]string,0)
                //读取列数
                for j=0;j<f.ListView1.Columns().Count()-1;j++{
                    //获取i行j列的值
                    s := f.ListView1.Items().Item(i).SubItems().S(j)
                    //UTF-8 to GBK
                    ItemsData = append(ItemsData,s)
                }
                writer.Write(ItemsData)
            }
            writer.Flush()
        }
    }
}



