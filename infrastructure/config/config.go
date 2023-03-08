package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// ServerConf 服务端口，名称，mode
type ServerConf struct {
	Address    string `yaml:"address"`
	ServerName string `yaml:"serverName"`
	Mode       string `yaml:"mode"`
}

// MysqlConf mysql
type MysqlConf struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbHost   string `yaml:"dbHost"`
	DbPort   int    `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
}

// LogConf 日志
type LogConf struct {
	LogFileDir string `yaml:"logFileDir"`
	AppName    string `yaml:"appName"`
	MaxSize    int    `yaml:"maxSize"`    //文件多大开始切分
	MaxBackups int    `yaml:"maxBackups"` //保留文件个数
	MaxAge     int    `yaml:"maxAge"`     //文件保留最大实际
}

// NsqConf nsq
type NsqConf struct {
	NsqProducerHost  string `yaml:"nsq_producer_host"`
	NsqProducerPort  string `yaml:"nsq_producer_port"`
	NsqSubscribeHost string `yaml:"nsq_subscribe_host"`
	NsqSubscribePort string `yaml:"nsq_subscribe_port"`
}

// jwt
type JwtConf struct {
	Secret string `yaml:"secret"`
}

// image
type ImageConf struct {
	RuntimeRootPath string   `yaml:"runtime_root_path"`
	ImagePrefixUrl  string   `yaml:"image_prefix_url"`
	ImageSavePath   string   `yaml:"image_save_path"`
	ImageMaxSize    int      `yaml:"image_max_size"`
	ImageAllowExts  []string `yaml:"image_allow_exts"`
}

// app
type AppConf struct {
	RunMode     string `yaml:"run_mode"`
	UserPwdSalt string `yaml:"user_pwd_salt"`
}

type Config struct {
	Server ServerConf `yaml:"server"`
	Mysql  MysqlConf  `yaml:"mysql"`
	Log    LogConf    `yaml:"log"`
	Nsq    NsqConf    `yaml:"nsq"`
	Jwt    JwtConf    `yaml:"jwt"`
	Image  ImageConf  `yaml:"imgae"`
	App    AppConf    `yaml:"app"`
}

func NewConfig() (conf *Config) {
	conf = initConfig()
	return
}

var (
	configOnce sync.Once
	config     *Config
)

// InitConfig 读取配置
func initConfig() *Config {
	configOnce.Do(func() {
		wd, _ := os.Getwd()
		for !strings.HasSuffix(wd, "garden-be") {
			wd = filepath.Dir(wd)
		}
		fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/manifest/config/config_pro.yaml", wd))
		if err != nil {
			panic(fmt.Sprintf("load config.yaml failed: %v", err))
		}

		err = yaml.Unmarshal(fileBytes, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
	})

	return config
}
