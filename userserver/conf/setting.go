package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	App      *App      `yaml:"app"`
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
	Redis    *Redis    `yaml:"redis"`
}

type App struct {
	PageSize     int    `json:"pageSize,omitempty" yaml:"pageSize"`
	JwtSecret    string `json:"jwtSecret,omitempty" yaml:"jwtSecret"`
	PassWordCost int    `json:"passWordCost,omitempty" yaml:"passWordCost"`

	ImagePrefixUrl string   `json:"imagePrefixUrl,omitempty" yaml:"imagePrefixUrl"`
	ImageSavePath  string   `json:"imageSavePath,omitempty" yaml:"imageSavePath"`
	ImageMaxSize   int      `json:"imageMaxSize,omitempty" yaml:"imageMaxSize"`
	ImageAllowExts []string `json:"imageAllowExts,omitempty" yaml:"imageAllowExts"`

	LogSavePath string `json:"logSavePath,omitempty" yaml:"logSavePath"`
	LogSaveName string `json:"logSaveName,omitempty" yaml:"logSaveName"`
	LogFileExt  string `json:"logFileExt,omitempty" yaml:"logFileExt"`
	TimeFormat  string `json:"timeFormat,omitempty" yaml:"timeFormat"`
}

type Server struct {
	RunMode      string        `json:"runMode,omitempty" yaml:"runMode"`
	HttpPort     int           `json:"httpPort,omitempty" yaml:"httpPort"`
	ReadTimeout  time.Duration `json:"readTimeout,omitempty" yaml:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout,omitempty" yaml:"writeTimeout"`
}

type Database struct {
	Type               string        `json:"type,omitempty" yaml:"type"`
	User               string        `json:"user,omitempty" yaml:"user"`
	Password           string        `json:"password,omitempty" yaml:"password"`
	Host               string        `json:"host,omitempty" yaml:"host"`
	Name               string        `json:"name,omitempty" yaml:"name"`
	Port               string        `json:"port,omitempty" yaml:"port"`
	SSLMode            string        `json:"SSLMode,omitempty" yaml:"SSLMode"`
	TimeZone           string        `json:"timeZone,omitempty" yaml:"timeZone"`
	TablePrefix        string        `json:"tablePrefix,omitempty" yaml:"tablePrefix"`
	SetMaxIdleConns    int           `json:"setMaxIdleConns,omitempty" yaml:"setMaxIdleConns"`
	SetMaxOpenConns    int           `json:"setMaxOpenConns,omitempty" yaml:"setMaxOpenConns"`
	SetConnMaxLifetime time.Duration `json:"setConnMaxLifetime,omitempty" yaml:"setConnMaxLifetime"`
}

type Redis struct {
	Host        string        `json:"host,omitempty" yaml:"host"`
	Password    string        `json:"password,omitempty" yaml:"password"`
	MaxIdle     int           `json:"maxIdle,omitempty" yaml:"maxIdle"`
	MaxActive   int           `json:"maxActive,omitempty" yaml:"maxActive"`
	IdleTimeout time.Duration `json:"idleTimeout,omitempty" yaml:"idleTimeout"`
}

var Cfg *Config

// Setup initialize the configuration instance
func Setup() {
	var err error

	cfg, err := ioutil.ReadFile("./conf/server.yml")
	if err != nil {
		log.Fatalf("setting.Setup, Fail to parse './conf/server.yml': %v", err)
	}

	var _config *Config
	if err = yaml.Unmarshal(cfg, &_config); err != nil {
		log.Fatalf("setting.Setup, Fail to Load config: %v", err)
	}

	_config.App.ImageMaxSize = _config.App.ImageMaxSize * 1024 * 1024
	_config.Server.ReadTimeout = _config.Server.ReadTimeout * time.Second
	_config.Server.WriteTimeout = _config.Server.WriteTimeout * time.Second
	_config.Database.SetConnMaxLifetime = _config.Database.SetConnMaxLifetime * time.Hour
	_config.Redis.IdleTimeout = _config.Redis.IdleTimeout * time.Second

	Cfg = _config
	println(_config.App)
}
