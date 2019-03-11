package thirdparty

import (
	"fmt"

	"github.com/go-ini/ini"
)

// ConfigListをstructとして形成
type ConfigList struct {
	Port      int
	DBName    string
	SQLDriver string
}

// structの中身を持つ変数Config。
var Config ConfigList

func init() {
	ctg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      ctg.Section("web").Key("port").MustInt(),
		DBName:    ctg.Section("db").Key("name").MustString("example.com"),
		SQLDriver: ctg.Section("db").Key("driver").String(),
	}
}

// 全体の結果を呼び出すための関数
func Ini() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
}
