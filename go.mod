module github.com/myml/sql2json

go 1.16

require gorm.io/driver/mysql v1.5.0

require gorm.io/driver/clickhouse v0.5.1

require (
	gorm.io/driver/postgres v1.5.0
	gorm.io/driver/sqlite v1.5.0
	gorm.io/driver/sqlserver v1.4.3
	gorm.io/gorm v1.25.0
)
