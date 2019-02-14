module hatgo

require (
	github.com/astaxie/beego v1.10.1
	github.com/davecgh/go-spew v0.0.0-20180830191138-d8f796af33cc
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7
	github.com/gin-gonic/gin v0.0.0-20180703091708-85221af84cf6
	github.com/go-ini/ini v0.0.0-20180615003539-cec2bdc49009
	github.com/go-redis/redis v6.14.0+incompatible
	github.com/go-sql-driver/mysql v0.0.0-20180618115901-749ddf1598b4
	github.com/go-xorm/builder v0.2.1
	github.com/go-xorm/core v0.6.0
	github.com/go-xorm/xorm v0.0.0-20180623065901-f16ce722ec15
	github.com/golang/protobuf v0.0.0-20180622174009-9eb2c01ac278
	github.com/json-iterator/go v1.1.5
	github.com/mattn/go-isatty v0.0.4
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/qiniu/api.v7 v0.0.0-20180813085939-10e8409dc200
	github.com/qiniu/x v0.0.0-20150721034113-f512abcf45ab
	github.com/ugorji/go v0.0.0-20180628102755-7d51bbe6161d
	golang.org/x/net v0.0.0
	golang.org/x/sync v0.0.0
	golang.org/x/sys v0.0.0
	golang.org/x/text v0.0.0
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.1
	qiniupkg.com/x v0.0.0-20150721034113-f512abcf45ab
)

replace golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd

replace golang.org/x/text v0.0.0 => github.com/golang/text v0.3.0

replace golang.org/x/sys v0.0.0 => github.com/golang/sys v0.0.0-20190213121743-983097b1a8a3

replace golang.org/x/sync v0.0.0 => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
