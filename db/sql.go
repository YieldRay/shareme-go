package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type mysql_db struct {
	db        *sql.DB
	tableName string
}

func (this mysql_db) Get(namespace string) (content string, err error) {
	DB := this.db
	rows, err := DB.Query("SELECT data FROM "+this.tableName+" WHERE namespace=?", namespace)
	defer rows.Close()
	if err != nil {
		return
	}
	if hasNext := rows.Next(); !hasNext {
		return "", nil
	}
	err = rows.Scan(&content)
	return
}

func (this mysql_db) Set(namespace, content string) bool {
	DB := this.db
	_, err := DB.Query("INSERT INTO "+this.tableName+" VALUES (?, ?) ON DUPLICATE KEY UPDATE data=?", namespace, content, content)

	if err != nil {
		return false
		// panic(err.Error())
	}
	return true
}

func MySQL(username, password, ip, port, dbName, tableName string) IDB {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbName)
	db, err := sql.Open("mysql", uri)

	if err != nil {
		fmt.Println("failed to connect mysql")
		panic(err.Error())
	}

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (namespace VARCHAR(32),data TEXT,PRIMARY KEY (namespace));", tableName))

	if err != nil {
		fmt.Println("failed to init table, please check if the database already exists")
		panic(err.Error())
	}

	return mysql_db{db, tableName}

}
