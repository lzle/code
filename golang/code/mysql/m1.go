package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"runtime"
	"time"
)

type T struct {
	id int
	k string
}


func query(db *sql.DB) T {
	row  := db.QueryRow("SELECT * FROM lzle.t where id = 1;")
	t := T{}
	row.Scan(&t.id,&t.k)
	return t
}

func update(db *sql.DB) {
	db.Exec("UPDATE lzle.t set k = k+1 where id = 1;")
	fmt.Println("update")
}


func txQuery(db *sql.Tx) T {
	row  := db.QueryRow("SELECT * FROM lzle.t where id = 1;")
	t := T{}
	row.Scan(&t.id,&t.k)
	return t
}

func txUpdate(db *sql.Tx) {
	db.Exec("UPDATE lzle.t set k = k+1 where id = 1;")
	fmt.Println("update")
}

func trans1() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.188.105:3306)/lzle?autocommit=true")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	tx,_:= db.Begin()
	fmt.Println("start transaction1")
	txQuery(tx)
	time.Sleep(time.Second*10)
	t := txQuery(tx)
	fmt.Println("trans1 k",t.k)
	tx.Commit()

}


func trans2() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.188.105:3306)/lzle?autocommit=true")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	tx,_:= db.Begin()
	fmt.Println("start transaction2")
	txQuery(tx)
	t := txQuery(tx)
	fmt.Println("trans2 k",t.k)
	time.Sleep(time.Second*2)
	txUpdate(tx)
	t = txQuery(tx)
	fmt.Println("trans2 k",t.k)
	time.Sleep(time.Second*10)
	tx.Commit()
}

func trans3() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.188.105:3306)/lzle?autocommit=true")
	if err != nil {
		log.Print(err)
	}
	tx,_:= db.Begin()
	txUpdate(tx)
	fmt.Println("trans3 txUpdate")
	time.Sleep(time.Second*6)
	fmt.Println("trans3 txUpdate done")
	tx.Commit()
}

func main()  {
	go trans1()
	time.Sleep(time.Second)
	go trans2()
	go trans3()
	runtime.Goexit()
}