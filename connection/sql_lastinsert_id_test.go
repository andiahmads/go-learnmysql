package connection

import (
	"context"
	"fmt"
	"testing"
)

func TestLastInserId(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()

	email := "bojo@gmail.com"
	comment := "coooooooooooooooooooooooooookkkk"

	sql := "INSERT INTO comments(email,comment) VALUES(?,?)"

	result, err := db.ExecContext(ctx, sql, email, comment)
	if err != nil {
		defer db.Close()
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("insert: ", insertId)
}
