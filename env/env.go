package env

import (
	"flag"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configPath    = "/config"
	componentName = "[ENV] "
)

type Auth struct {
	JWTKey              string
	Issuer              string
	TokenExpireTimeHour int
}

type Server struct {
	RunMode         string
	Port            string
	ReadTimeoutSec  int
	WriteTimeoutSec int
	NumCPU          string
}

type DB struct {
	Type           string
	User           string
	Password       string
	Host           string
	Port           int
	Name           string
	MaxIdleConn    int
	MaxOpenConn    int
	MaxLifeTimeSec int
}

type Oauth struct {
	Google Google
}

type Google struct {
	Scopes       []string
	CredFilePath string
	CallbackURL  string
}

type Env struct {
	Auth   Auth
	Server Server
	DB     DB
	Oauth  Oauth
}

var Conf Env

func Setup() {
	setupLog()

	configName := flag.String("config", "", "config file name")
	flag.Parse()

	if flag.NFlag() == 0 {
		log.Fatal(componentName, "need a config file name")
	}

	rootDir, _ := filepath.Abs("./")

	viper.SetConfigName(*configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(rootDir + configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(componentName, "cannot read a config file")
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatal(componentName, "config has a wrong schema")
	}
}

func setupLog() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}
