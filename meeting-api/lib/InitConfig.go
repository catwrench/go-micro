package lib

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/spf13/viper"
	"os"
)

type ConfigStruct struct {
	Etcd     Etcd
	Jaeger   Jaeger
	GoMicro  GoMicro
	Db       Db
	Datetime Datetime
	Url      Url
}

//------GoMicro 基础组件-------
type Etcd struct {
	Port string
	Addr string
}
type Jaeger struct {
	Port string
}

//-------GoMicro 服务-------
type GoMicro struct {
	Gateway Gateway
	Api     Api
	Srv     Srv
}
type Gateway struct {
	Web  string
	Name string
	Port string
}
type Api struct {
	Meeting Meeting
}

type Meeting struct {
	Name string
	Port string
}

type Srv struct {
	User    User
	Meeting Meeting
	Notice  Notice
}
type User struct {
	Name string
	Port string
}
type Notice struct {
	Name string
	Port string
}

//--------其他配置----------
type Db struct {
	Name     string
	Addr     string
	Port     string
	Username string
	Password string
}
type Datetime struct {
	Datetime string
	Date     string
	Time     string
}
type Url struct {
	Uims Uims
}
type Uims struct {
	Scheme string
	Host   string
}

var Config *ConfigStruct

var ConfigViper *viper.Viper
var Err *viper.Viper

func init() {
	//加载配置
	ConfigViper = newViper(getConfigNameByEnv())
	//将config映射到结构体
	Config = &ConfigStruct{}
	MapToStruct(ConfigViper, &Config)
	//加载错误码
	Err = newViper("errors")
}

//创建viper实例
func newViper(configName string) *viper.Viper {
	v := viper.New()
	//直接指定文件
	//v.SetConfigFile("./submodules/common/config/errors.yml")
	// 设置文件名称（不包含后缀），多个设置后面会覆盖前面
	v.SetConfigName(configName)
	// 设置后缀名，支持多个设置，会按优先级读取文件类型
	v.SetConfigType("yml")
	// 设置文件所在路径,应该是相对的go.mod所在路径，支持多个设置
	v.AddConfigPath("./submodules/common/config")
	// 设置默认参数
	v.Set("verbose", true)

	//找到并读取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("viper：配置文件未找到")
		} else {
			log.Error("viper：配置错误：", err)
			panic("viper：配置错误")
		}
	}

	// 监控配置和重新获取配置
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变更:", e.Name)
	})
	return v
}

//根据环境变量读取配置文件
func getConfigNameByEnv() string {
	//读取环境变量
	env := os.Getenv("env")
	configName := "config"
	switch env {
	case "dev":
		configName = configName + ".dev"
	case "qa":
		configName = configName + ".qa"
	case "prod":
		configName = configName + ".prod"
	default:
	}
	return configName
}

//讲配置映射到结构体
func MapToStruct(v *viper.Viper, s interface{}) {
	err := v.Unmarshal(&s)
	if err != nil {
		log.Error("映射配置到结构体失败：", err)
		return
	}
	return
}
