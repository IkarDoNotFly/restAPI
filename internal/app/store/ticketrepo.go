package store

type TicketRepo struct{
	store *Store
}
/*
//Create...
func(r *TicketRepo) CreateTicket(ticket *models.Ticket) (*models.Ticket,error){
	if err:=r.store.db.QueryRow(
		"INSERT INTO ticket(baggage,type_menu,date,cost) VALUES($1) RETURNING id",
		ticket.name,).Scan(&ticket.id);err!=nil{
		return nil,err
	}
	return ticket,nil
}
//FindById...
func(r *TicketRepo) FindTicketById(id int) (*models.Ticket,error){
	c:=&models.Ticket{}
	if err:=r.store.db.QueryRow(
		"SELECT id,name FROM ticket WHERE id=$1;",id).Scan(&c.id);err!=nil{
		return nil,err
	}

	return c, nil
}

//DeleteById...
func(r *TicketRepo) DeleteTicketById(id int) (error){
	query:=`DELETE FROM ticket WHERE id=$1`
	_,err:=r.store.db.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}

//UpdateById

func(r *TicketRepo) UpdateTicketById(id int,baggage string,type_menu string,date string,cost float64) (error){
	query:=`UPDATE ticket SET baggage=$1,type_menu=$2,date=$3,cost=$4 WHERE id=$5;`
	_,err:=r.store.db.Exec(query,baggage,type_menu,date,cost,id)
	if err!=nil{
		return err
	}
	return nil
}*/