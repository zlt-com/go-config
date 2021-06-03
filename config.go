package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration 系统配置
type Configuration struct {
	//系统配置
	Port        string
	UseBindata  bool
	CookieStore string
	SessionName string
	MPVerify    string
	Uploads     string

	//数据库配置
	DBSource string
	DBType   string
	LogMode  bool
	LogFile  string

	// Elastic配置
	ElasticEnable       bool
	ElasticHostURL      string
	ElasticIndexName    string
	ElasticSQLLogName   string
	ElasticErrorLogName string
	ElasticInfoLogName  string
	ElasticLoginLogName string
	ElasticWxMsgLogName string
	ElasticUser         string
	ElasticPassword     string
	ElasticSourceHost   string

	//nsq消息队列配置
	NSQServerIP string
	NSQLookupd  string
	NSQTopic    string

	//Redis
	RedisHost string
	RedisPass string
	RedisType string

	//微信配置
	WechatInfo *Weixin

	//Oauth认证
	OauthClientID     string
	OauthClientSecret string
	OauthRedirectURI  string
	OauthServer       string
}

// Weixin 微信公众号基本信息
type Weixin struct {
	APPID          string
	SECRET         string
	TOKEN          string
	EncodingAESKey string
	Domain         string
	RedisHost      string
	RedisPass      string
	RedisType      string
	Validate       bool
}

// Config 全局配置
var (
	Config *Configuration
)

func init() {
	//读取配置文件
	fmt.Println("...读取配置文件...")
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("...读取配置文件错误：" + err.Error())
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
}
