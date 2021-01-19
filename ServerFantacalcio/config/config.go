package config

import "database/sql"

//Funzione che si occupa del collegamento al database
func GetDB()(db *sql.DB, err error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "a1b2c3d4e5"
	dbName := "fantacalcio"

	db,err = sql.Open(dbDriver,dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/"+dbName)

	return
}
