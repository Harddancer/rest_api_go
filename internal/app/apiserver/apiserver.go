package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Harddancer/rest_api_go/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Api server

type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New
func New(config *Config) *Apiserver {
	fmt.Println("Done")
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *Apiserver) Start() error {
	if err := s.cofigureLogger(); err != nil {
		fmt.Print("Ошибка на старте")
		return err
	}
	s.configureRouter()

	if err := s.configureStore();err != nil{
		return err
	}




	s.logger.Info("starting API  server")

	return http.ListenAndServe(s.config.BindAddr, s.router)

}

func (s *Apiserver) cofigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *Apiserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *Apiserver) configureStore() error{
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil{
		return err
	}


s.store = st
return nil
}


func (s *Apiserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}

}
