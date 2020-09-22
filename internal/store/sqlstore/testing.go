package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

//Возврашает тестовый сторе который сконфигурирован и также возвращает функцию которая очищает таблицы при
//последующих вызовах. Делаем фековую бд и передаем ее в качестве аргумента
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	//Метод говорит тестам что этот метод ненужно тестировать, нигде не учитывыать
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	//Функция которая перегружает таблицу, очищает после тестов
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}





