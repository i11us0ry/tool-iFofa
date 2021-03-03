package ui

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
)

//::private::
type TForm3Fields struct {
	Grammer1 []Grammer
}

type Grammer struct {
	Example     string
	Description string
	Notes       string
}

func (f *TForm3) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("查询语法参考")
	f.SetColor(0x241813)
	f.ListView1.SetColor(0xcfb593)

	f.Grammer1 = make([]Grammer, 35)
	f.Grammer1[0].Example = "title=\"beijing\""
	f.Grammer1[1].Example = "header=\"jboss\""
	f.Grammer1[2].Example = "body=\"Hacked by\""
	f.Grammer1[3].Example = "domain=\"qq.com\""
	f.Grammer1[4].Example = "icon_hash=\"-247388890\""
	f.Grammer1[5].Example = "host=\".gov.cn\""
	f.Grammer1[6].Example = "port=\"443\""
	f.Grammer1[7].Example = "ip=\"1.1.1.1\""
	f.Grammer1[8].Example = "ip=\"220.181.111.1/24\""
	f.Grammer1[9].Example = "status_code=\"402\""
	f.Grammer1[10].Example = "protocol=\"https\""
	f.Grammer1[11].Example = "city=\"Hangzhou\""
	f.Grammer1[12].Example = "region=\"Zhejiang\""
	f.Grammer1[13].Example = "country=\"CN\""
	f.Grammer1[14].Example = "cert=\"google\""
	f.Grammer1[15].Example = "banner=users && protocol=ftp"
	f.Grammer1[16].Example = "type=service"
	f.Grammer1[17].Example = "os=windows"
	f.Grammer1[18].Example = "server==\"Microsoft-IIS/7.5\""
	f.Grammer1[19].Example = "app=\"HIKVISION-视频监控\""
	f.Grammer1[20].Example = "after=\"2017\" && before=\"2017-10-01\""
	f.Grammer1[21].Example = "asn=\"19551\""
	f.Grammer1[22].Example = "org=\"Amazon.com, Inc.\""
	f.Grammer1[23].Example = "base_protocol=\"udp\""
	f.Grammer1[24].Example = "is_ipv6=true"
	f.Grammer1[25].Example = "is_domain=true"
	f.Grammer1[26].Example = "ip_ports=\"80,161\""
	f.Grammer1[27].Example = "port_size=\"6\""
	f.Grammer1[28].Example = "port_size_gt=\"3\""
	f.Grammer1[29].Example = "port_size_lt=\"12\""
	f.Grammer1[30].Example = "p_country=\"CN\""
	f.Grammer1[31].Example = "ip_region=\"Zhejiang\""
	f.Grammer1[32].Example = "ip_city=\"Hangzhou\""
	f.Grammer1[33].Example = "ip_after=\"2019-01-01\""
	f.Grammer1[34].Example = "ip_before=\"2019-07-01\""

	f.Grammer1[0].Description = "从标题中搜索“北京”"
	f.Grammer1[1].Description = "从http头中搜索“jboss”"
	f.Grammer1[2].Description = "从html正文中搜索abc"
	f.Grammer1[3].Description = "搜索根域名带有qq.com的网站。"
	f.Grammer1[4].Description = "搜索使用此icon的资产。"
	f.Grammer1[5].Description = "从url中搜索”.gov.cn”"
	f.Grammer1[6].Description = "查找对应“443”端口的资产"
	f.Grammer1[7].Description = "从ip中搜索包含“1.1.1.1”的网站"
	f.Grammer1[8].Description = "查询IP为“220.181.111.1”的C网段资产"
	f.Grammer1[9].Description = "查询服务器状态为“402”的资产"
	f.Grammer1[10].Description = "查询https协议资产"
	f.Grammer1[11].Description = "搜索指定城市的资产。"
	f.Grammer1[12].Description = "搜索指定行政区的资产。"
	f.Grammer1[13].Description = "搜索指定国家(编码)的资产。"
	f.Grammer1[14].Description = "搜索证书(https或者imaps等)中带有google的资产。"
	f.Grammer1[15].Description = "搜索FTP协议中带有users文本的资产。"
	f.Grammer1[16].Description = "搜索所有协议资产，支持subdomain和service两种。"
	f.Grammer1[17].Description = "搜索Windows资产。"
	f.Grammer1[18].Description = "搜索IIS 7.5服务器。"
	f.Grammer1[19].Description = "搜索海康威视设备"
	f.Grammer1[20].Description = "时间范围段搜索"
	f.Grammer1[21].Description = "搜索指定asn的资产。"
	f.Grammer1[22].Description = "搜索指定org(组织)的资产。"
	f.Grammer1[23].Description = "搜索指定udp协议的资产。"
	f.Grammer1[24].Description = "搜索ipv6的资产"
	f.Grammer1[25].Description = "搜索域名的资产"
	f.Grammer1[26].Description = "搜索同时开放80和161端口的ip"
	f.Grammer1[27].Description = "查询开放端口数量等于\"6\"的资产"
	f.Grammer1[28].Description = "查询开放端口数量大于\"3\"的资产"
	f.Grammer1[29].Description = "查询开放端口数量小于\"12\"的资产"
	f.Grammer1[30].Description = "搜索中国的ip资产(以ip为单位的资产数据)。"
	f.Grammer1[31].Description = "搜索指定行政区的ip资产(以ip为单位的资产数据)。"
	f.Grammer1[32].Description = "搜索指定城市的ip资产(以ip为单位的资产数据)。"
	f.Grammer1[33].Description = "搜索2019-01-01以后的ip资产(以ip为单位的资产数据)。"
	f.Grammer1[34].Description = "搜索2019-07-01以前的ip资产(以ip为单位的资产数据)。"

	f.Grammer1[0].Notes = "-"
	f.Grammer1[1].Notes = "-"
	f.Grammer1[2].Notes = "-"
	f.Grammer1[3].Notes = "-"
	f.Grammer1[4].Notes = "仅限高级会员使用"
	f.Grammer1[5].Notes = "搜索要用host作为名称"
	f.Grammer1[6].Notes = "-"
	f.Grammer1[7].Notes = "搜索要用ip作为名称"
	f.Grammer1[8].Notes = "-"
	f.Grammer1[9].Notes = "-"
	f.Grammer1[10].Notes = "搜索指定协议类型(在开启端口扫描的情况下有效)"
	f.Grammer1[11].Notes = "-"
	f.Grammer1[12].Notes = "-"
	f.Grammer1[13].Notes = "-"
	f.Grammer1[14].Notes = "-"
	f.Grammer1[15].Notes = "-"
	f.Grammer1[16].Notes = "搜索所有协议资产"
	f.Grammer1[17].Notes = "-"
	f.Grammer1[18].Notes = "-"
	f.Grammer1[19].Notes = "-"
	f.Grammer1[20].Notes = "-"
	f.Grammer1[21].Notes = "-"
	f.Grammer1[22].Notes = "-"
	f.Grammer1[23].Notes = "-"
	f.Grammer1[24].Notes = "搜索ipv6的资产,只接受true和false。"
	f.Grammer1[25].Notes = "搜索域名的资产,只接受true和false。"
	f.Grammer1[26].Notes = "搜索同时开放80和161端口的ip资产(以ip为单位的资产数据)"
	f.Grammer1[27].Notes = "仅限FOFA会员使用"
	f.Grammer1[28].Notes = "仅限FOFA会员使用"
	f.Grammer1[29].Notes = "仅限FOFA会员使用"
	f.Grammer1[30].Notes = "搜索中国的ip资产"
	f.Grammer1[31].Notes = "搜索指定行政区的资产"
	f.Grammer1[32].Notes = "搜索指定城市的资产"
	f.Grammer1[33].Notes = "搜索2019-01-01以后的ip资产"
	f.Grammer1[34].Notes = "搜索2019-07-01以前的ip资产"

	f.ListView1.Items().BeginUpdate()
	for i := 0; i <= 34; i++ {
		//行
		item := f.ListView1.Items().Add()
		//行标题
		item.SetCaption(fmt.Sprintf(f.Grammer1[i].Example))
		//行的第一列
		item.SubItems().Add(fmt.Sprintf(f.Grammer1[i].Description))
		//行的第二列
		item.SubItems().Add(fmt.Sprintf(f.Grammer1[i].Notes))
	}
	f.ListView1.Items().EndUpdate()

	//可以接收键盘操作
	f.SetKeyPreview(true)
	// 键盘弹起事件
	f.SetOnKeyUp(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == keys.VkEscape {
			f.Close()
		}
	})
}
