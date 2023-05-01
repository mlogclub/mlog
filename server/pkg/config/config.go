package config

import (
	"io/ioutil"

	"github.com/mlogclub/simple/sqls"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Instance *Config

type Config struct {
	Env        string `yaml:"Env"`        // 环境：prod、dev
	BaseUrl    string `yaml:"BaseUrl"`    // base url
	QQUrl      string `yaml:"QQUrl"`      // QQ url
	Port       string `yaml:"Port"`       // 端口
	LogFile    string `yaml:"LogFile"`    // 日志文件
	StaticPath string `yaml:"StaticPath"` // 静态文件目录
	IpDataPath string `yaml:"IpDataPath"` // IP数据文件

	// 数据库配置
	DB sqls.DbConfig `yaml:"DB"`

	Jwt struct {
		SignKey       string `yaml:"SignKey"`
		ExpireSeconds int    `yaml:"ExpireSeconds"`
		Issuer        string `yaml:"Issuer"`
	} `yaml:"Jwt"`

	// Github
	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	// OSChina
	OSChina struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"OSChina"`

	// QQ登录
	QQConnect struct {
		AppId  string `yaml:"AppId"`
		AppKey string `yaml:"AppKey"`
	} `yaml:"QQConnect"`

	// 阿里云oss配置
	Uploader struct {
		Enable    string `yaml:"Enable"`
		AliyunOss struct {
			Host          string `yaml:"Host"`
			Bucket        string `yaml:"Bucket"`
			Endpoint      string `yaml:"Endpoint"`
			AccessId      string `yaml:"AccessId"`
			AccessSecret  string `yaml:"AccessSecret"`
			StyleSplitter string `yaml:"StyleSplitter"`
			StyleAvatar   string `yaml:"StyleAvatar"`
			StylePreview  string `yaml:"StylePreview"`
			StyleSmall    string `yaml:"StyleSmall"`
			StyleDetail   string `yaml:"StyleDetail"`
		} `yaml:"AliyunOss"`
		TxyCos struct {
			Host      string `yaml:"Host"`
			Bucket    string `yaml:"Bucket"`
			Region    string `yaml:"Region"`
			SecretID  string `yaml:"SecretID"`
			SecretKey string `yaml:"SecretKey"`
		} `yaml:"TxyOss"`
		Server struct {
			URL  string `yaml:"URL"`
			Host string `yaml:"Host"`
		} `yaml:"Server"`
		Local struct {
			Host string `yaml:"Host"`
			Path string `yaml:"Path"`
		} `yaml:"Local"`
	} `yaml:"Uploader"`

	// 百度SEO相关配置
	// 文档：https://ziyuan.baidu.com/college/courseinfo?id=267&page=2#h2_article_title14
	BaiduSEO struct {
		Site  string `yaml:"Site"`
		Token string `yaml:"Token"`
	} `yaml:"BaiduSEO"`

	// 神马搜索SEO相关
	// 文档：https://zhanzhang.sm.cn/open/mip
	SmSEO struct {
		Site     string `yaml:"Site"`
		UserName string `yaml:"UserName"`
		Token    string `yaml:"Token"`
	} `yaml:"SmSEO"`

	// smtp
	Smtp struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SSL      bool   `yaml:"SSL"`
	} `yaml:"Smtp"`

	// es
	Es struct {
		Url           string `yaml:"Url"`
		IndexTopic    string `yaml:"IndexTopic"`
		IndexArticles string `yaml:"IndexArticles"`
	} `yaml:"Es"`

	// alipay
	Alipay struct {
		AppId         string `yaml:"AppId"`
		AppPrivateKey string `yaml:"AppPrivateKey"`
		AppPublicKey  string `yaml:"AppPublicKey"`
		BaseUrl       string `yaml:"BaseUrl"`
	} `yaml:"Alipay"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}
