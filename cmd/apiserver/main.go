package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Mikeyteam/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

//Предварительная инициализация пути для конфига. Записываем в флаги
func init()  {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()  {
	//Парсим наши флаги которые ранее созданы при инициализации
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath,config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

