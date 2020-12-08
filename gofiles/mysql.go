package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

//Mysql ----------------------------------------------------------
type Mysql struct {
	DB        *sql.DB
	execute   mysqlExecute
	getStrAry mysqlGetStrAry
}

//Mysql constructor
func newMysql() (mysql Mysql, err error) {
	err = godotenv.Load("../mysql.env")
	DB, err := sql.Open("mysql", os.Getenv("ID")+":"+os.Getenv("PS")+"@tcp("+os.Getenv("ENDPOINT")+":"+os.Getenv("PORT")+")/"+os.Getenv("SCHEMA")+"?multiStatements=true")
	if err != nil {
		return mysql, err
	}
	err = DB.Ping()
	if err != nil {
		return mysql, err
	}
	mysql.DB = DB
	mysql.getStrAry.DB = DB
	mysql.execute.DB = DB
	return mysql, nil
}

//mysql.execute
type mysqlExecute struct {
	DB *sql.DB
}

//Mysql.execute.query
func (mysql *mysqlExecute) query(queryStr string) (err error) {
	_, err = mysql.DB.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

//mysql.getStrAry
type mysqlGetStrAry struct {
	DB *sql.DB
}

//Mysql.getStrAry.query
func (mysql *mysqlGetStrAry) query(queryStr string) (resStrAry []string, err error) {
	rows, err := mysql.DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	return columns, nil
}
