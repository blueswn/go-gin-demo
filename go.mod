module go-gin-demo

go 1.12

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190609082536-301114b31cce
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190313210603-aa82965741a9
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.42.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	gopkg.in/ini.v1 v1.42.0 // indirect
)
