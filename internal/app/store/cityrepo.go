package store

import "api_server/models"

type CityRepo struct {

	store *Store
}
//Create...
func(r *CityRepo) Create(city *models.City) (*models.City,error){
	if err:=r.store.db.QueryRow("INSERT INTO city(name) VALUES($1) RETURNING id",city.name,).Scan(&city.id);err!=nil{
	return nil,err
	}
	return city,nil
}
//FindById...
func(r *CityRepo) FindById(id int) (*models.City,error){
	c:=&models.City{}
	if err:=r.store.db.QueryRow(
		"SELECT id,name FROM city WEHERE city=$1",id).Scan(&c.id,&c.name);err!=nil{
		return nil,err
	}

	return c, nil
}