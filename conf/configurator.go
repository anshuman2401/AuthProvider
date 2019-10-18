package conf

import (
	"fmt"
	"os"
	"time"
	"github.com/jinzhu/configor"
)

type Cassandra struct {
	Cluster []string
	DC string
	Keyspace string
	Backoff                   time.Duration
	Timeout                   time.Duration
	KeepAlive                 time.Duration
	NumConns                  int
}

type Config struct {
	Cassandra Cassandra
}

func LoadConfig()  Config{
	var config Config
	if err:= configor.Load(&config,  os.Getenv("GOPATH")+"/src/AuthProvider/conf/config.json"); err != nil {
		fmt.Print("Error in loading confif")
	}

	fmt.Print("Config values are = ", config)
	return config
}

var GetConfig = LoadConfig()