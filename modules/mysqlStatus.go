package modules

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func MysqlStatusCheck(c *gin.Context) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-east-1.prod.aws.tidbcloud.com",
	})
	db, err := sql.Open("mysql", "2bGNBbVNcaX62MH.root:DCoO9ftOflRLOCLS@tcp(gateway01.us-east-1.prod.aws.tidbcloud.com:4000)/status?tls=tidb")
	//db, err := sql.Open("mysql", "2bGNBbV23123NcaX62MH.sadsaroot:DsdadCoO9ftOflRLOCLS@tcp(gateway01.prod.aws.tidbcloud.com:4000)/status?tls=tidb")
	if err == nil {
		//fmt.Println(err)
		//log.Fatal("failed to connect database", err)
		c.JSON(200, "success")

	}
	defer db.Close()

	var dbName string
	err = db.QueryRow("SELECT DATABASE();").Scan(&dbName)
	if err != nil {
		//log.Fatal("failed to execute query", err)
		fmt.Println("failed to execute query", err)
		c.JSON(500, "mysql database not found")
		return
	}
	fmt.Println(dbName)
	c.JSON(200, "mysql health ok")
	return
}
