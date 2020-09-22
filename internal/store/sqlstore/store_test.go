package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

//Конфигурирует из env переменней DATABASE_URL. Эта функция вызывается1 раз перед всеми тестамами в данном пакете
//В ней хорошо создавать какието разовые вещи которые иницилизируются. В данном случе конфигурируем DATABASE_URL
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=indrik sslmode=disable"
	}

	os.Exit(m.Run())
}
