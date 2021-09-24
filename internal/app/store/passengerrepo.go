package store

import "api_server/models"

type PassengerRepo struct {

	store *Store
}

//Create...
func(r *PassengerRepo) CreatePassenger(passenger *models.Passenger) (*models.Passenger,error){
	if err:=r.store.db.QueryRow(
		"INSERT INTO passenger(name,surname) VALUES($1,$2) RETURNING id",
		passenger.Name,passenger.Surname).Scan(&passenger.Id);err!=nil{
		return nil,err
	}
	return passenger,nil
}
//FindById...
func(r *PassengerRepo) FindPaseengerById(id int) (*models.Passenger,error){
	p:=&models.Passenger{}
	if err:=r.store.db.QueryRow(
		"SELECT id,name,surname FROM passenger WHERE id=$1;",id).Scan(&p.Id,&p.Name,&p.Surname);err!=nil{
		return nil,err
	}

	return p, nil
}

//DeleteById...
func(r *PassengerRepo) DeletePassengerById(id int) (error){
	query:=`DELETE FROM passenger WHERE id=$1`
	_,err:=r.store.db.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}

//UpdateById

func(r *PassengerRepo) UpdatePassengerById(id int,name string,surname string) (error){
	query:=`UPDATE passenger SET name=$1,surname=$2 WHERE id=$3;`
	_,err:=r.store.db.Exec(query,name,surname,id)
	if err!=nil{
		return err
	}
	return nil
}