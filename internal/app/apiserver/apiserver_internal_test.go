package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandlerHello(t *testing.T) {
	//Создаем новый сервер с новым конфигом
	server := New(NewConfig())
	//Создадим ответ типо новый Response от сервера
	recorder := httptest.NewRecorder()
	//Создаем запрос request на необходимый url
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	//Вызываем наш обработчимк и передадим на вход ответ response и наш запрос request
	server.handlerHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Hello")
}
