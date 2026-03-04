package model

import (
	"time"

	"github.com/bits-and-atoms/Go_REST_API/db"
)

// exec used for insert,update ,delete create
//query used for select
//prepare can be used with both exec and query, it is for performance , so if you dont close it then it saves
//the query in memory so increases speed but if you close it after that it behaves as non prepare exec or query only

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	insert into events(name,description,location,dateTime,user_id)
	values (?,?,?,?,?)
	`
	// this ? way prevents sql injection
	st, err := db.DB.Prepare(query)
	if err != nil {
		defer st.Close()
		return err
	}
	result, err := st.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	//exec used for INSERT, UPDATE, DELETE and create
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}
func (e *Event) Update() error {
	query := `
	update events
	set name =?,description =?,location=?,dateTime=?
	where id = ?
	`
	st, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer st.Close()
	_, err = st.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	return nil
}
func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()
	result, err := st.Query()
	//query used for select
	if err != nil {
		return nil, err
	}
	var events []Event
	for result.Next() {
		var curr Event
		err := result.Scan(&curr.ID, &curr.Name, &curr.Description, &curr.Location, &curr.DateTime, &curr.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, curr)
	}
	return events, nil
}
func GetEventById(id int64) (*Event, error) {
	// without prepare
	// query := "select * from events where id = ?"
	// result := db.DB.QueryRow(query, id)
	// var curr Event
	// err := result.Scan(&curr.ID, &curr.Name, &curr.Description, &curr.Location, &curr.DateTime, &curr.UserID)
	// if err != nil {
	// 	return nil, err
	// }
	// return &curr,nil

	//with prepare
	query := "select * from events where id = ?"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	result := st.QueryRow(id)
	var curr Event
	err = result.Scan(&curr.ID, &curr.Name, &curr.Description, &curr.Location, &curr.DateTime, &curr.UserID)
	if err != nil {
		return nil, err
	}
	return &curr, nil
}

func (e *Event) Delete() error {
	query := "delete from events where id = ?"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer st.Close()
	_, err = st.Exec(e.ID)
	return err
}
func (e *Event) Register(userId int64) error {
	query := "insert into registrations (event_id,user_id) values (?,?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer st.Close()
	_, err = st.Exec(e.ID,userId)
	return err
}

func (e *Event) CancelRegistration(userId int64)error{
	query := "delete from registrations where event_id = ? and user_id = ?"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer st.Close()
	_, err = st.Exec(e.ID,userId)
	return err
}