package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	MySQL       string            `yaml:"MySQL"`
	Redis       RedisConfig       `yaml:"Redis"`
	VideoServer VideoServerConfig `yaml:"VideoServer"`
}

type RedisConfig struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
	PoolSize int    `yaml:"PoolSize"`
}

type VideoServerConfig struct {
	Ip                   string `yaml:"Ip"`
	Host                 string `yaml:"Host"`
	User                 string `yaml:"User"`
	Password             string `yaml:"Password"`
	FTPServerIp          string `yaml:"FTPServerIp"`
	FTPServerGetFilePort string `yaml:"FTPServerGetFilePort"`
}

var Config Configuration

var (

	// redis
	TokenTTL float64 = 3600
	TokenKey string  = "token:"

	//dirPath
	CommonFilePath = "home/douyin/FTPServer"
	VideoDir       = "videos"
	PhotoDir       = "photos"
)

var MailPassword = os.Getenv("MailPassword")

type ProblemBasic struct {
	Identity          string      `json:"identity"`           // 问题表的唯一标识
	Title             string      `json:"title"`              // 问题标题
	Content           string      `json:"content"`            // 问题内容
	ProblemCategories []int       `json:"problem_categories"` // 关联问题分类表
	MaxRuntime        int         `json:"max_runtime"`        // 最大运行时长
	MaxMem            int         `json:"max_mem"`            // 最大运行内存
	TestCases         []*TestCase `json:"test_cases"`         // 关联测试用例表
}

type TestCase struct {
	Input  string `json:"input"`  // 输入
	Output string `json:"output"` // 输出
}

var (
	DateLayout            = "2006-01-02 15:04:05"
	ValidGolangPackageMap = map[string]struct{}{
		"bytes":   {},
		"fmt":     {},
		"math":    {},
		"sort":    {},
		"strings": {},
	}
)

// 首字母大写其他的包才能调用
func ReadConfig() {
	configFile, err := ioutil.ReadFile("config/configuration.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
