package apiserver

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const(
	host="172.17.0.2"
	port="8080"
	user="postgres"
	password="password"
	dbname="Airport"
)
//APIServer...
type APIServer struct{
	psqlInfo string
	logger *logrus.Logger
}
//New...
func New() *APIServer{
	return &APIServer{
		psqlInfo:fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disable",host,port,user,password,dbname),
		logger:logrus.New(),
	}
}
//Start...
func(s *APIServer) Start() error{
	return nil
}
//Какой логер использовать?
//configureLoger..
func (s* APIServer) configureLoger() error{

}