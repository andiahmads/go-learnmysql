package connection

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

/*
untuk operasi sql yg tidak membutuhkan hasil, kita bisa gunakan perintah Exec,
namun jika kita membutuhkan hasil kita bisa menggunakan method QueryContext(context,sql,params) */

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	t.Skip("skip testing")

	ctx := context.Background()

	// insertSql := "INSERT INTO customer(name,balance,raiting,birth_date,merried) VALUES('andi',20000,9.0,'1992-09-09',false)"
	insertSql := "INSERT INTO customer(name,email,balance,raiting,birth_date,merried) VALUES('andi','andi.fivesco@gmail.com',20000,9.0,'1992-09-09',false)"
	_, err := db.ExecContext(ctx, insertSql)
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS")
}

func TestQuerySelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	t.Skip("skip testing")
	ctx := context.Background()

	query := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}
}

func TestGetAllCustomers(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name, email sql.NullString
		var balance int
		var raiting float64
		var merried bool
		var birth_date, created_at time.Time

		err = rows.Scan(&id, &name, &email, &balance, &raiting, &merried, &birth_date, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		fmt.Println("email: ", email)
		fmt.Println("balance: ", balance)
		fmt.Println("raiting: ", raiting)
		fmt.Println("crated_at: ", created_at)
		fmt.Println("birth_date: ", birth_date)
		fmt.Println("merried: ", merried)
	}
}

/*
=============================Rows==================================
hasil query function adalah sebuah data  struct sql.Rows.
rows digunakan untuk melakukan iterasi terhadap hasil query.
kita bisa menggunakan function(Rows)(Next)(Boolean) untuk melakukan iterasi terhadap data.
hasil query, jika return data false, artinya sudah tidak ada data lagi didalam result.
untuk membaca tiap data, kita bisa menggunakan (Rows)Scan(Column.....) dan harus menggunakan pointer.
kenapa harus menggunakan pointer?, karena saat proses scan for rows.Next akan set data sesuai parameternya.
dan jangan lupa untuk menutupnya dengan perintah,
Rows.Close() */

/* Null label Type
golang database tidak mengerti dengan tipe data null.
oleh karena itu, khusus kolom yg bisa null akan menjadi masalah jika kita melakukan scan
maka solusinya adalah kita menggunakan (sql.NullString etc.)
*/
