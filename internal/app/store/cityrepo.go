package store

import "api_server/models"

type CityRepo struct {

	store *Store
}
//Create...
func(r *CityRepo) Create(city *models.City) (*models.City,error){
	query:="INSERT INTO city(name) VALUES($1) RETURNING id"
	if err:=r.store.db.QueryRow(query,city.Name).Scan(&city.Id);err!=nil{
	return nil,err
	}
	return city,nil
}
//FindById...
func(r *CityRepo) FindById(id int) (*models.City,error){
	c:=&models.City{
	}
	if err:=r.store.db.QueryRow(
		"SELECT id,name FROM city WHERE id=$1;",id).Scan(&c.Id,&c.Name);err!=nil{
		return nil,err
	}

	return c, nil
}

//DeleteById...
func(r *CityRepo) DeleteById(id int) (error){
	query:=`DELETE FROM city WHERE id=$1`
	_,err:=r.store.db.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}

//UpdateById

func(r *CityRepo) UpdateById(id int,name string) (error){
	query:=`UPDATE city SET name=$1 WHERE id=$2;`
	_,err:=r.store.db.Exec(query,name,id)
	if err!=nil{
		return err
	}
	return nil
}