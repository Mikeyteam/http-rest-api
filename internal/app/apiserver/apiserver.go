package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//Возвращает указатель на наш APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(), //Создали новый логер
		router: mux.NewRouter(), //Создали новый роутер
	}
}

func (s *APIServer) Start() error {
	//Вызываем при старте сервера в начале логгер
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	//Парсим уровень логирования у нас в конфиге - debug
	level, err := logrus.ParseLevel(s.config.Loglevel)
	if err != nil {
		return err
	}
	//Устанавливем данный уровень (debug) в наш логер
	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter () {
	s.router.HandleFunc("/hello", s.handlerHello())
}

func (s *APIServer) handlerHello() http.HandlerFunc {
    //Здесь можно описать какие то локальные перменные, только используются в данном handler

	return func(w http.ResponseWriter, r *http.Request){
       //Здесь можно описать бизнес логику запроса
		io.WriteString(w, "Hello")
	}
}
