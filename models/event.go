// storing logic of event data
package models

import (
	"time"

	"example.com/project/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"` //required field, it will error if it is missing
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time 
	UserID int64
}

//save data to database
//we should pass receiver as a pointer to get/manage the data from the original
func (e *Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,date_time,user_id)
	VALUES (?,?,?,?,?)`

	//prepare query statement. alternatively, we can directly execute Exec() method but we use prepare method for better performance
	//prepare method will store syntax in memory and easily reuse it in efficient way
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	e.DateTime = time.Now()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	//get the id of lastest inset
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	//sending query 
	//use query method instead exec beacause query use when we want get back a multiple row
	//use query for fetching data
	//query method return pointer of the row and error
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	//Next method return a boolean which is true as long as there are row left
	for rows.Next() {
		var event Event
		//populate content of row
		err := rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
		if err != nil {
			return nil,err
		}
		//store data from each row inside slice
		events = append(events, event)
	}
	return events,nil
}

func GetEventById(id int64) (*Event, error) {
	//to protect sequel injection by defining placeholder (?)
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil,err
	}
	return &event, nil
}

//we should pass receiver as a pointer to get/manage the data from the original
func (e Event) Updated() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.ID)
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID,userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE From registrations WHERE event_id = ? user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID,userId)
	return err
}