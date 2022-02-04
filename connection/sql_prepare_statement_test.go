package connection

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

/*
query atau exec dengan parameter
saat kita menggunakan  function query atau exec dengan parameter, sebenarnya kita menggunakan prepare statement.
jadi tahapan pertama statementnya disiapkan terlebih dahulu, setelah itu baru diisi dengan parameter.
kadang kasus kita ingin melakukan beberapa hal sekaligus, hanya berbeda parameternya, misalnya data lansung banyak.
pembuatan prepare statement bisa dilakukan dengan manual, tanpa harus menggunakan query atau exec dengan parameter.

prepare statement mengenali koneksi database yang digunakan secara otomatis.
untuk prepare statement kita bisa menggunakan function (DB)Prepare(context,sql)
prepare statement direpresentasikan dalam struct database/sql.Stmt.
prepare statement harus di Close() ketika tidak digunakan lagi.
*/

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	sql := "INSERT INTO comments(email,comment) VALUES (?,?)"
	statement, err := db.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "andi" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke" + strconv.Itoa(i)
		resutl, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := resutl.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("inser id ke = ", id)
	}

}
