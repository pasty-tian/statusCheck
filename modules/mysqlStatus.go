package modules

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func MysqlStatusCheck(c *gin.Context) {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-east-1.prod.aws.tidbcloud.com",
	})

	db, err := sql.Open("mysql", "2bGNBbVNcaX62MH.root:DCoO9ftOflRLOCLS@tcp(gateway01.us-east-1.prod.aws.tidbcloud.com:4000)/status?tls=tidb")
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()

	var dbName string
	err = db.QueryRow("SELECT DATABASE();").Scan(&dbName)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
	fmt.Println(dbName)
}
