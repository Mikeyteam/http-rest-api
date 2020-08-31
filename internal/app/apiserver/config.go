package apiserver

type Config struct {
	BindAddr string `toml:"bin_addr"`
	Loglevel string `toml:"log_level"`
}

//Отдает указатель на инициализированный обьект - конфиг с дефолтными парамаметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		Loglevel: "debug",
	}
}
