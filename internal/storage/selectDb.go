package storage

import (
	"database/sql"
	"fmt"
	"os"
)

func SelectDb(whichDB string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch whichDB {
	case "sqlite3":
		os.Remove("./CapTracker.db")
		db, err = sql.Open("sqlite3", "./CapTracker.db")
		if err != nil {
			return nil, fmt.Errorf("Error; %s", err)
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cap (
    	id INTEGER PRIMARY KEY, 
    	instrument TEXT,
		timestamp datetime,
		bid_price float,
		ask_price float,
        last_price float,
		volume float)`)
		if err != nil {
			return nil, fmt.Errorf("Error; %s", err)
		}

	case "postgres":
		db, err = sql.Open("postgres", "user=postgres password=postgres dbname=captracker sslmode=disable")
		if err != nil {
			return nil, fmt.Errorf("Error; %s", err)
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cap (
    	id SERIAL PRIMARY KEY,
    	instrument TEXT,
    	timestamp TIMESTAMP,
    	bid_price float,
    	ask_price float,
    	last_price float,
    	volume float)`)
		if err != nil {
			return nil, fmt.Errorf("Error; %s", err)
		}

	default:
		err = fmt.Errorf("unsupported DB type: %s", whichDB)
		return nil, err
	}
	return db, nil
}
