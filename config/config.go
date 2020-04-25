package config

import (
	"fmt"
	dlog "log"
	"github.com/lexkong/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// RunmodeDebug -
	RunmodeDebug = "debug"
	// RunmodeRelease -
	RunmodeRelease = "release"
	// RunmodeTest -
	RunmodeTest = "test"

	// 配置文件路径
	configFilePath = "./config.yaml"
	// 日志文件路径
	logFilePath = "storage/logs/dofun.log"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
	// MailConfig 邮件配置
	MailConfig *mailConfig
)

// InitConfig 初始化配置
func InitConfig(c string, hasLog bool) {
	if c == "" {
		c = configFilePath
	}
	// 初始化 viper 配置
	viper.SetConfigFile(c)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v", err))
	}

	// 初始化日志
	if hasLog {
		initLog()
		// 热更新配置文件
		watchConfig()
	}
	// 初始化 app 配置
	AppConfig = newAppConfig()
	// 初始化数据库配置
	DBConfig = newDBConfig()
	// 初始化邮件配置
	MailConfig = newMailConfig()
	//热加载
	//hotReload()
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		// 配置文件更新了
		log.Infof("Config file changed: %s", ev.Name)
	})
}

func hotReload() {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher();
	if err != nil {
		dlog.Fatal(err);
	}
	defer watch.Close();
	//添加要监控的对象，文件或文件夹
	err = watch.Add("app/");
	if err != nil {
		dlog.Fatal(err);
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Create == fsnotify.Create {
						dlog.Println("创建文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						dlog.Println("写入文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						dlog.Println("删除文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						dlog.Println("重命名文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						dlog.Println("修改权限 : ", ev.Name);
					}
				}
			case err := <-watch.Errors:
				{
					dlog.Println("error : ", err);
					return;
				}
			}
		}
	}();

	//循环
	//select {};
}