package config

import (
	"log"
	"sync"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var (
	config Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(initconfig)
	return &config
}

type Config struct {
	Loglevel      string `env:"LogLevel,default=info"`
	LogFile       string `env:"LogFile,default=log/backup.log"`
	LogMaxSize    int    `env:"LogMaxSize,default=10"`
	LogMaxBackups int    `env:"LogMaxBackups,default=50"`
	LogMaxAge     int    `env:"LogMaxAge,default=30"`

	Location string `env:"Location,default=unknown"`

	Webdav            bool   `env:"Webdav,default=false"`
	WebdavRoot        string `env:"WebdavRoot"`
	WebdavUser        string `env:"WebdavUser"`
	WebdavPassword    string `env:"WebdavPassword"`
	WebdavStoragePath string `env:"WebdavStoragePath"`
}

func initconfig() {
	// load from .env
	if err := godotenv.Load("/etc/ftbbackup/.env"); err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Load .env Error: ", err)
		}
	}

	config = Config{}

	// load to config
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal("Load Environment Error: ", err)
	}

	log.Printf("Load Config Success: %+v\n", config)
}
