package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf 注册全局变量
var Conf = new(AppConfig)

// AppConfig 全局配置
type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Version       string `mapstructure:"version"`
	VideoSize     int64  `mapstructure:"video_size"`
	PicSize       int64  `mapstructure:"pic_size"`
	StartTime     string `mapstructure:"start_time"`
	MachineId     int64  `mapstructure:"machine_id"`
	RateLimitTime int64  `mapstructure:"rate_limit_time"`
	RateLimitNum  int64  `mapstructure:"rate_limit_num"`
	Port          int    `mapstructure:"port"`

	*AuthConfig  `mapstructure:"auth"`
	*LogConfig   `mapstructure:"log"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*OssConfig   `mapstructure:"oss"`
	*EmailConfig `mapstructure:"email"`
}

// AuthConfig jwt配置
type AuthConfig struct {
	JwtExpire int `mapstructure:"jwt_expire"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level             string `mapstructure:"level"`
	Filename          string `mapstructure:"filename"`
	UserClickFilename string `mapstructure:"userClickFileName"`
	MaxSize           int    `mapstructure:"max_size"`
	MaxAge            int    `mapstructure:"max_age"`
	MaxBackups        int    `mapstructure:"max_backups"`
}

// MySqlConfig mysql配置
type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// OssConfig oss配置
type OssConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
	BucketPoint     string `mapstructure:"bucket_point"`
}

// EmailConfig oss配置
type EmailConfig struct {
	User       string `mapstructure:"user"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Host       string `mapstructure:"host"`
	ReplyEmail string `mapstructure:"reply_email"`
}

func Init(filename string) (err error) {
	// 方法一
	viper.SetConfigFile(filename) //从指定文件中读取配置文件

	// 方法二
	//viper.SetConfigName("config") //指定配置文件名称
	//viper.SetConfigType("yaml")   //指定配置文件类型(远程获取配置文件时指定类型)
	//viper.AddConfigPath(".")   //指定查找配置文件的路径
	err = viper.ReadInConfig() //读取配置文件
	if err != nil {
		fmt.Printf("viper.ReadInConfig() err: %v \n", err)
		return
	}
	// 把读取到到配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v \n", err)
	}
	// 监控配置文件
	viper.WatchConfig()
	// 当配置文件修改了
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了！")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v \n", err)
		}
	})
	return
}
