package connection

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlParameter(t *testing.T) {
	db := GetConnection()

	defer db.Close()
	ctx := context.Background()

	username := "andi"
	password := "admin123"

	sql := "SELECT username FROM user WHERE username=? AND password=? limit 1"

	//tambahkan parameter
	rows, err := db.QueryContext(ctx, sql,username,password)
	if err != nil {
		defer rows.Close()
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("suuccesss login", username)
	} else {
		fmt.Println("gagal login", username)

	}

}
