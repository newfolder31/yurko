package infrastructures

//
//import (
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	_ "github.com/lib/pq"
//	"log"
//)
//
////FOR MORE INFORMATION ABOUT sqlx link: https://github.com/jmoiron/sqlx
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "yurkorole"
//	password = "secret"
//	dbname   = "yurkodb"
//)
//
//type PostgresHandler struct {
//	Conn *sqlx.DB
//}
//
///**
//example:
//	query:	"INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)"
//	domain:	&Person{"Jane", "Citizen", "jane.citzen@example.com"}
//*/
//func (handler *PostgresHandler) Execute(query string, domain interface{}) error {
//	_, err := handler.Conn.NamedExec(query, &domain)
//	return err
//}
//
///**
//example:
//	query:	`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`
//	data:	map[string]interface{}{
//			"first": "Bin",
//			"last":  "Smuth",
//			"email": "bensmith@allblacks.nz",
//		})
//*/
//func (handler *PostgresHandler) NamedExec(query string, data map[string]interface{}) error {
//	_, err := handler.Conn.NamedExec(query, data)
//	return err
//}
//
///**
//example:
//	domain: {}
//	query:	"SELECT * FROM person WHERE first_name=$1"
//	args:	"Jason"
//*/
//func (handler *PostgresHandler) Get(domain interface{}, query string, args ...interface{}) error {
//	err := handler.Conn.Get(&domain, query, args...)
//	return err
//}
//
//func (handler *PostgresHandler) QueryRowx(domain interface{}, query string, args ...interface{}) *sqlx.Row {
//	row := handler.Conn.QueryRowx(query, args...)
//	return row
//}
//
///**
//example:
//	slice:	[]Person{}
//	query:	"SELECT * FROM place ORDER BY first_name ASC"
//*/
//func (handler *PostgresHandler) Select(slice []interface{}, query string) error {
//	err := handler.Conn.Select(&slice, query)
//	return err
//}
//
//func NewPostgresHandler() *PostgresHandler {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	conn, err := sqlx.Connect("postgres", psqlInfo)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	postgresHandler := new(PostgresHandler)
//	postgresHandler.Conn = conn
//	return postgresHandler
//}
