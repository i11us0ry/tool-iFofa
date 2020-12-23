#### 0x01 iFofa 
#### by:i11us0ry

#### 0x02 开发环境：
	go version go1.15.3 windows/amd64
	libvcl-2.0.8

#### 0x03 编译

	main.go目录下执行go build -buildmode=exe

	目前只做了win的，mac的没条件做适配，linux觉得没必要、有编译好的文件，不想编译的可以直接使用

#### 0x04菜单功能说明

### 1.用户设置

![](https://s3.ax1x.com/2020/12/23/r6PQRP.md.png)

	检测到用户输入的email和key后会对email进行格式判断，若正确则向fofa请求验证用户信息，这一步是为了获取用户会员等级为后面其他功能做铺垫，若fofa返回正确信息，则将用户email、key、Vip_level保存到main同级目录下config.ini文件中，以后启动iFofa时程序会自动从config.ini读取用户信息

### 2.请求参数

	### 2.1 请求数量

		默认为100条，fofa官方推荐每次<=100条，理由是body字段包含内容较多，在实际操作中最好设置为可请求最大数量的整除数，如最多可请求100条的则推荐设置10、20、25、50等，不推荐11、21、26、51等
		
	### 2.2 可选参数

		官方介绍可选的列表有：host title ip domain port country province city country_name header server protocol banner cert isp as_number as_organization latitude longitude structinfo。
		
		iFofa只设置了常用的可选的列表：host,ip,title,domain,port,country,province,city,country_name,header,server,protocol,banner

		初次启动默认为host，ip, title用逗号分隔多个参数

		注意：country是国家代码，例如CN, country_name是国家名称；structinfo仅限企业会员调用

	请求数量和可选参数设置好后会被记录到config.ini文件中，以后启动iFofa时程序会自动从config.ini请求参数

![](https://s3.ax1x.com/2020/12/23/r6nvNT.png)

### 3.语法参考

![](https://s3.ax1x.com/2020/12/23/r6up34.png)

	将官方给的参考给拷贝了下来，方便随时查看

#### 0x05 右键功能说明

![](https://s3.ax1x.com/2020/12/23/r6nx4U.png)

### 1.页数跳转

1.1首页

1.2上一页

1.3下一页

1.4尾页

### 2.保存功能

2.1目前只做了csv

### 3.清除功能

3.1清除面板内容，不改变其他参数

### 4.退出功能

4.1退出

