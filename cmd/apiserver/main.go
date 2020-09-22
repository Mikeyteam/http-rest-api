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
	//Создаем флаг - переменную configPath, значение по умолчанию configs/apiserver.toml
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()  {
	//Парсим наши флаги которые ранее созданы при инициализации и записываем знчения configs/apiserver.toml в соответсвующую переменную
	//configPath. Это делается через библиотеку toml которая прочитает файл и распарсит apiserver.toml
	flag.Parse()

	config := apiserver.NewConfig()
	//Т.е мы из toml файла записываем в наш config
	_, err := toml.DecodeFile(configPath,config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

