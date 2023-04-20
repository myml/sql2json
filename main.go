package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var openMap = map[string]func(string) gorm.Dialector{
	"mysql":      mysql.Open,
	"postgres":   postgres.Open,
	"sqlite":     sqlite.Open,
	"sqlserver":  sqlserver.Open,
	"clickhouse": clickhouse.Open,
}

func main() {
	var dbType, dns, sql string
	var format, array bool
	flag.StringVar(&dbType, "db", "mysql", "database type, mysql,postgres,sqlite,sqlserver,clickhouse")
	flag.StringVar(&dns, "dns", "", "see https://gorm.io/docs/connecting_to_the_database.html#SQLite")
	flag.StringVar(&sql, "sql", "", "select * from mysql")
	flag.BoolVar(&format, "format", false, "output format json")
	flag.BoolVar(&array, "array", false, "output json array. default output json lines")
	flag.Parse()
	open := openMap[dbType]
	if open == nil || len(dns) == 0 || len(sql) == 0 {
		flag.PrintDefaults()
		return
	}
	db, err := gorm.Open(open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	out := os.Stdout
	encoder := json.NewEncoder(out)
	if format {
		encoder.SetIndent("", "\t")
	}
	first := true
	for rows.Next() {
		var m map[string]interface{}
		err = db.ScanRows(rows, &m)
		if err != nil {
			log.Fatal(err)
		}
		if array {
			if first {
				out.Write([]byte{'['})
				if format {
					out.Write([]byte{'\n'})
				}
				first = false
			} else {
				out.Write([]byte{','})
				if format {
					out.Write([]byte{'\n'})
				}
			}
		}
		err = encoder.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if array {
		out.Write([]byte{']'})
	}
}
