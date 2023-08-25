package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('eko', 'Eko')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date,  married, created_at FROM golang.customer"
	rows, err := db.QueryContext(ctx, script)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("===========================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("email: ", email)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birthDate: ", birthDate)
		}
		fmt.Println("createdAt: ", createdAt)
		fmt.Println("married: ", married)
	}
}

func TestSqlInjection(t *testing.T) {

	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM golang.user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1 "
	rows, err := db.QueryContext(ctx, script)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login", username)
	}
}

func TestSqlInjectionSafe(t *testing.T) {

	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM golang.user WHERE username = ? AND password = ? LIMIT 1 "
	rows, err := db.QueryContext(ctx, script, username, password)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login", username)
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "Eko"
	password := "Eko"

	script := "INSERT INTO customer(id, name) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	email := "Eko"
	comment := "Eko"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with comment id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Commit()

	// Rollback
	// err = tx.Rollback()

	if err != nil {
		panic(err)
	}
}
