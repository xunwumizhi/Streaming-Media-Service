package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func AddUser(loginName string, pwd string) error {
	db := openConn()

	
	.Close()
}