package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	mdb *sql.DB
	content = context.Background()
)

func init() {
	//mdb, err := mysql.NewConnector(&mysql.Config{
	//	Net: "tcp",
	//	Addr: "10.2.4.58:3306",
	//	User: "root",
	//	Passwd: "admin123",
	//	DBName: "cokdb2",
	//})

	mdb, _ := sql.Open("mysql", "root:admin123@tcp(10.2.4.58:3306)/cokdb2")
	//设置数据库最大连接数
	mdb.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	mdb.SetMaxIdleConns(10)
	//验证连接
	if err := mdb.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}

	fmt.Println("connnect success")
	rows, err := mdb.QueryContext(content, "select name from userprofile where gmFlag = ? order by lastOnlineTime desc limit 10", 1)
	if err != nil {
		fmt.Println("QueryContext error", err)
		return
	}

	defer rows.Close()
	names := make([]string, 1)

	var length int
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		length++
		names = append(names, name)
	}
	fmt.Println(length)
	fmt.Println(names)

	mdb.Close()
}

func TestMysql() {
	fmt.Println()
}
