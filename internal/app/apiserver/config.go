package apiserver

type Config struct {
	BindAddr string `toml:"bin_addr"`
	Loglevel string `toml:"log_level"`
	Store    *store.Config
}

//Типо конструктора Config. Отдает указатель на инициализированный обьект - конфиг с дефолтными
//парамаметрами для Аписервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		Loglevel: "debug",
		Store:    store.NewConfig(),
	}
}
