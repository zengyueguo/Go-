package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct{
	UserID string
	Info string
}

//插入demo
func insert(user User)(dbErr error) {
	ds := "root:gzy1234@tcp(192.168.254.3:3306)/db4?charset=utf8"

	db, dbErr := sql.Open("mysql", ds)
	if dbErr!=nil{
		return errors.Wrap(dbErr,"main:insert open database fail")
	}

	stmt, dbErr := db.Prepare(`INSERT table1 (UserID,Info) values (?,?)`)
	if dbErr!=nil{
		return errors.Wrap(dbErr,"main:insert Sql Prepare failed")
	}

	_, dbErr = stmt.Exec(user.UserID, user.Info)
	if dbErr!=nil{
		return errors.Wrap(dbErr,"main:insert Sql Exec failed")
	}
	return nil
}

func main() {
	user:=User{"001","This is User001's information"}
	err :=insert(user)
	if err!=nil{
		fmt.Println(errors.Cause(err))
		fmt.Println(err)
	}
}