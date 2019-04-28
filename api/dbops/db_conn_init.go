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

//mysql连接测试成功后`go test db_conn_test.go`，再将`初始化连接`过程放入init()
func init() {
	fmt.Println("Enter `conn.go` init() ...")

	dbConn, err = sql.Open("mysql", "root:root@/streaming_media_service?charset=utf8") //
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("dbConn + %v\n", dbConn) //Println不能格式化输出，直接拼接；格式化输出使用Printf
}
