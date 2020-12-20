#iFofa by:i11us0ry

#开发环境：
go version go1.15.3 windows/amd64
libvcl-2.0.8

#编译
#目前只做了win的，mac的没条件，linux觉得没必要
main.go目录下执行go build -buildmode=exe
#有编译好的文件，不想编译的可以直接使用

#菜单功能说明
1.用户设置
	检测到用户输入的email和key后会对email进行格式判断，若正确则向fofa请求验证用户信息，这一步是为了获取用户会员等级为后面其他功能做铺垫，若fofa返回正确信息，则将用户email、key、Vip_level保存到main同级目录下config.ini文件中，以后启动iFofa时程序会自动从config.ini读取用户信息

2.请求参数
	2.1 请求数量
		默认为100条，fofa官方推荐每次<=100条，理由是body字段包含内容较多，在实际操作中最好设置为可请求最大数量的整除数，如最多可请求100条的则推荐设置10、20、25、50等，不推荐11、21、26、51等
	2.2 可选参数
		官方介绍可选的列表有：host title ip domain port country province city country_name header server protocol banner cert isp as_number as_organization latitude longitude structinfo。
		iFofa只设置了常用的可选的列表：host,ip,title,domain,port,country,province,city,country_name,header,server,protocol,banner
		初次启动默认为host，ip, title用逗号分隔多个参数
		注意：country是国家代码，例如CN, country_name是国家名称；structinfo仅限企业会员调用

	请求数量和可选参数设置好后会被记录到config.ini文件中，以后启动iFofa时程序会自动从config.ini请求参数

3.语法参考
	将官方给的参考给拷贝了下来，方便随时查看


#右键功能说明
1.页数跳转
	#页数说明，fofa请求时除了请求数量、可选参数之外还可以设置请求页数，根据多次实验猜测可请求页数与资产数MaxSize、会员最大请求数Maxnums、及每次请求数量num有关,当然只是个人猜测

	#用逻辑语来讲就是可请求页数pages = MaxSize>Maxnums?math.floor(Maxnums/num):math.ceil(MaxSize/num)

	#以普通会员为例,搜索关键字title="fofa and goby nb!",
	#假设title="fofa and goby nb!"的资产MaxSize=120条，每次请求num=15条，可请求页pages=6
	#假设title="fofa and goby nb!"的资产MaxSize=80条，每次请求num=15条，可请求页pages=6

	假设当前
		处于关于title="fofa and goby nb!"的前一百条中的第三页的十五条数据（31-45）
	则以下功能为：
	1.1.首页
		请求关于title="fofa and goby nb!"的前一百条中的第一页的十五条数据（1-15）
	1.2.上一页
		请求关于title="fofa and goby nb!"的前一百条中的第二页的十五条数据（16-30）
	1.3.下一页
		请求关于title="fofa and goby nb!"的前一百条中的第四页的十五条数据（46-60）
	1.4.尾页
		如果MaxSize=120，num=15
			请求关于title="fofa and goby nb!"的前一百条中的第六页的十五条数据（76-90）
		如果MaxSize=80，num=15
			请求关于title="fofa and goby nb!"的前一百条中的第六页的十五条数据（76-80）
	因为15不是100的整除数，所以最终获取到的数据会相对于少一点，所以说了这么多，还是希望在请求的时候将请求数num设置为最大请求数Maxnums的整除数

2.保存功能
	目前只做了csv

3.清除功能
	清除面板内容，不改变其他参数

4.退出功能
	退出

