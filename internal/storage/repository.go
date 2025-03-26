package storage

import (
	"database/sql"

	"CapitalParser/internal/models"
)

func InsertCapData(db *sql.DB, data models.CapitalData, dbType string) error {
	if dbType == "sqlite3" {
		query := "INSERT INTO cap (uuid, instrument, timestamp, bid_price, ask_price, last_price, volume) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(query, data.UUID, data.InstrumentName, data.Timestamp, data.BidPrice, data.AskPrice, data.LastPrice, data.Volume)
		if err != nil {
			return err
		}
	}
	if dbType == "postgres" {
		query := "INSERT INTO cap (uuid, instrument, timestamp, bid_price, ask_price, last_price, volume) VALUES ($1, $2, $3, $4, $5, $6, $7)"
		_, err := db.Exec(query, data.UUID, data.InstrumentName, data.Timestamp, data.BidPrice, data.AskPrice, data.LastPrice, data.Volume)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCapData(db *sql.DB, uuid string, dbType string) (data models.CapitalData, err error) {
	if dbType == "sqlite3" {
		row := db.QueryRow("SELECT uuid, instrument, timestamp, bid_price, ask_price, last_price, volume FROM cap WHERE uuid = ?", uuid)
		err := row.Scan(&data.UUID, &data.InstrumentName, &data.Timestamp, &data.BidPrice, &data.AskPrice, &data.LastPrice, &data.Volume)
		if err != nil {
			return models.CapitalData{}, err
		}
		return data, nil
	}
	if dbType == "postgres" {
		row := db.QueryRow("SELECT uuid, instrument, timestamp, bid_price, ask_price, last_price, volume FROM cap WHERE uuid = $1", uuid)
		err := row.Scan(&data.UUID, &data.InstrumentName, &data.Timestamp, &data.BidPrice, &data.AskPrice, &data.LastPrice, &data.Volume)
		if err != nil {
			return models.CapitalData{}, err
		}
		return data, nil
	}
	return models.CapitalData{}, nil
}

func UpdateCapData(db *sql.DB, data models.CapitalData, dbType string) (ok bool, err error) {
	if dbType == "sqlite3" {
		_, err := db.Exec("UPDATE cap SET instrument = ?, timestamp = ?, bid_price = ?, ask_price = ?, last_price = ?, volume = ? WHERE uuid = ?", data.InstrumentName, data.Timestamp, data.BidPrice, data.AskPrice, data.LastPrice, data.Volume, data.UUID)
		if err != nil {
			return false, err
		}
	}
	if dbType == "postgres" {
		_, err := db.Exec("UPDATE cap SET instrument = $1, timestamp = $2, bid_price = $3, ask_price = $4, last_price = $5, volume = $6 WHERE uuid = $7", data.InstrumentName, data.Timestamp, data.BidPrice, data.AskPrice, data.LastPrice, data.Volume, data.UUID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func DeleteCapData(db *sql.DB, uuid string, dbType string) (ok bool, err error) {
	if dbType == "sqlite3" {
		_, err := db.Exec("DELETE FROM cap WHERE uuid = ?", uuid)
		if err != nil {
			return false, err
		}
	}
	if dbType == "postgres" {
		_, err := db.Exec("DELETE FROM cap WHERE uuid = $1", uuid)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
