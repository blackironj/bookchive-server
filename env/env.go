package env

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

const configPath = "/config"

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

type Database struct {
	Type           string
	User           string
	Password       string
	Host           string
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
	Auth     Auth
	Server   Server
	Database Database
	Oauth    Oauth
}

var Conf Env

func Setup() {
	configName := flag.String("config", "", "config file name")
	flag.Parse()

	if flag.NFlag() == 0 {
		log.Fatal("need a config file name")
	}

	rootDir, _ := filepath.Abs("./")

	viper.SetConfigName(*configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(rootDir + configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("cannot read a config file")
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatal("config has a wrong schema")
	}
}
