package model

import (
	"time"

	"github.com/bits-and-atoms/Go_REST_API/db"
)

type Event struct{
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}
func (e *Event) Save() error{
	query := `
	insert into events(name,description,location,dateTime,user_id)
	values (?,?,?,?,?)
	`
	// this ? way prevents sql injection 
	st,err := db.DB.Prepare(query)
	defer st.Close()
	if err != nil{
		return err
	}
	result, err := st.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	//exec used for INSERT, UPDATE, DELETE and create
	if err != nil{
		return err
	}
	id,err := result.LastInsertId()
	if err != nil{
		return err
	}
	e.ID = id
	return nil
}
func GetAllEvents() ([]Event,error){
	query := "select * from events"
	st,err := db.DB.Prepare(query)
	defer st.Close()
	if err != nil{
		return nil,err
	}
	result,err := st.Query()
	//query used for select
	if err != nil{
		return nil,err
	}
	var events []Event
	for result.Next(){
		var curr Event
		err := result.Scan(&curr.ID,&curr.Name,&curr.Description,&curr.Location,&curr.DateTime,&curr.UserID)
		if err != nil{
			return nil,err
		}
		events = append(events, curr)
	}
	return events,nil
}