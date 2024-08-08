package main

import (
	"flag"
	"fmt"
	"mall-admin/cmd/initdb/logic"
	"mall-admin/cmd/initdb/system"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/init_db.yaml", "the config file")

func main() {
	flag.Parse()

	var c logic.InitDB
	conf.MustLoad(*configFile, &c)
	system.Register()

	var s logic.InitDBService
	err := s.InitDB(&c)
	if err != nil {
		fmt.Println("创建数据库失败", err.Error())
	} else {
		fmt.Println("创建数据库成功")
	}
}
