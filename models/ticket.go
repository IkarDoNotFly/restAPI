package models

type Ticket struct {
	id int
	baggage string
	menu_type string
	date string
	cost float64
	id_passenger int
	point_a int //id of city
	point_b int //id of city
}
