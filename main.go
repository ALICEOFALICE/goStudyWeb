package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMysql() {
	//初始化启动信息
	dsn := "zhang:123456789@tcp(127.0.0.1:3306)/goweb"
	//初始化错误变量
	var err error
	//创建数据库连接
	db, err = sql.Open("mysql", dsn)
	//错误检查，如果对象为空就抛出错误
	if err != nil {
		fmt.Println("错误详细", err)
		panic("发送一个致命错误，数据库初始化失败")
	}
	//ping并处理错误
	err = db.Ping()
	if err != nil {
		fmt.Println("错误详细：", err)
		panic("连接数据库失败")
	}
}
func main() {
	//初始化数据库连接
	initMysql()
	//主函数结束后关不连接
	defer db.Close()

}
