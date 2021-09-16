package main

import (
	"restAPI/internal/app/apiserver"
	"log"
)

type Passenger struct{
	id int
	name_passenger string
	surname_passenger string

}
type Cities struct {
	id int
	name string
}

type Ticket struct{
	id int
	baggage int
	type_menu string
	date string
	cost float64
	//id_passenger int
	//id_point_a int
	//id_point_b int


}



func main(){
	config:=apiserver.NewConfig()

	s:=apiserver.New(config)
	err:=s.Start()
	if err!=nil{
		log.Fatal(err)
	}
}

