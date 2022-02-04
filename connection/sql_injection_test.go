package connection

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlInjections(t *testing.T) {
	db := GetConnection()

	defer db.Close()
	ctx := context.Background()

	username := "andi';#"
	password := "admin123"

	sql := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' limit 1"

	//contoh query manipulasi dari sql injection
	fmt.Println(sql)
	/* SELECT username FROM user WHERE username='andi';#' AND password='admin123' limit 1
	DILIHAT PADA QUERY DIATAS, BAHWA datanda # merupakan perintah command dimysql, sehingga  field password diabaikan

	solusi:
	jangan membuat querysql secara manual dengan menggabungkan string secara bulat2
	jika kita membutuhkan parameter, kita bisa menggunakan function execute atau query dengan parameter.
	*/
	rows, err := db.QueryContext(ctx, sql)
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
