package apiserver

import (
	"api_server/internal/app/store"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)


//APIServer...
type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}
//New...
func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) configureLogger() error{
	level,err:=logrus.ParseLevel(s.config.LogLevel)
	if err!=nil{
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer)configureRouter(){
	s.router.HandleFunc("/",s.handleHelp())
	s.router.HandleFunc("/getCity",s.handleGetCity())
	s.router.HandleFunc("/getPassenger",s.handleGetPassenger())
}

func (s *APIServer) handleHelp() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"I can find city,passenger,ticket by id. " +
			"Example: localhost:8080/getCity?id=1")
	}
}
func(s *APIServer) handleGetPassenger() http.HandlerFunc{
	store:=store.PassengerRepo{}
	return func(w http.ResponseWriter,r *http.Request){
		id,err:=strconv.Atoi(r.URL.Query().Get("id"))
		if err!=nil{
			io.WriteString(w,"Wrong id! Try again")
		}
		responce,err:=store.FindPaseengerById(id)

		if(err!=nil){
			io.WriteString(w,"Passenger not found!")
		}
		str,err:=json.Marshal(responce)
		fmt.Println(str)
		if err!=nil{
			io.WriteString(w,"Something wrong ")
		}


		io.WriteString(w,string(str))
	}

}
func(s *APIServer) handleGetCity() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"City")
	}
}
//handleGetTrades..
func(s *APIServer) handleGetTrades(){

}
//Start...
func(s *APIServer) Start() error{
	if err:=s.configureLogger();err!=nil{
		return err
	}
	s.configureRouter()

	if err:=s.configureStore();err!=nil{
		return err
	}
	s.logger.Info("API server started!")
	return http.ListenAndServe(s.config.BindAddr,s.router)
}

func(s *APIServer) configureStore() error{
	st:=store.New(s.config.Store)
	if err:=st.Open();err!=nil{
		return err
	}
	s.store=st
	return nil
}




