package dbops
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	fmt.Println("Enter `conn.go` init() ...")

	dbConn, err := sql.Open("mysql", "root:root@/streaming_media_service?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("dbConn +%v\n", dbConn)
}
