package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@/streaming_media_service")
	defer db.Close()

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}
