package config

import (
    "log"    
	"os"
    "time"    
	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {   
	Env string `yaml:"env" env-default:"prod" env-required:"true"`
	HTTP        HTTP   `yaml:"http_server"`
	DB          DB     `yaml:"db"`
}

type HTTP struct {    
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080" env-required:"true"`
    Timeout     time.Duration `yaml:"timeout" env-default:"4s"`  
    IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type DB struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	DBName   string `yaml:"dbname" env-default:"postgres"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}


func MustLoad() *Config {    

	configPath := os.Getenv("CONFIG_PATH")

    if configPath == "" {        
		log.Fatal("config_path is not set")
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {        
		log.Fatalf("config file does not exist: %s", configPath)
    }

    var cfg Config

    if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {        log.Fatalf("cannot read confige file %v", err)
    }
    return &cfg
}