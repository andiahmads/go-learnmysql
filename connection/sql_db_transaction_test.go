package connection

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

/* Transaction GOlang

secara default, semua perintah SQL yg kita kirim akan otomatis di commit.
dengan fitur transaction,fitur auto commit tidak akan dieksekusi.
untuk memulai transaction, kita bisa menggunakan function (DB) Begin() yg menghasilkan struct Tx yang merupakan representasi transaction.
struct Tx digunakan sebagai pengganti DB. setelah proses transaction suucces kita bisa menggunakan function (Tx) untuk melakukan commit / rollback.

*/

func TestDbTransctions(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO comments (email, comment) VALUES(?,?)"
	// do transaction
	for i := 1; i < 10; i++ {
		email := "dbtx" + strconv.Itoa(i) + "@gmail.com"
		comment := "transaction ke" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sql, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("transaction with id: ", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
