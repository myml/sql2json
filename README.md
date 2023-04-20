# sql2json

sql database query result export to json lines

## Install

`go install -x github.com/myml/sql2json@latest`

## Command Args

```sh
  -db string
        database type, mysql,postgres,sqlite,sqlserver,clickhouse (default "mysql")
  -dns string
        see https://gorm.io/docs/connecting_to_the_database.html#SQLite
  -sql string
        select * from mysql
  -format
        output format json
```

## Example

```sh
sql2json -db mysql -dns "user:password@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local" -sql "select * from user"
```

Output:

```json
{"Alter_priv":"N","Alter_routine_priv":"N","Create_priv":"N","Create_routine_priv":"N","Create_tmp_table_priv":"N","Create_view_priv":"N","Db":"performance_schema","Delete_priv":"N","Drop_priv":"N","Event_priv":"N","Execute_priv":"N","Grant_priv":"N","Host":"localhost","Index_priv":"N","Insert_priv":"N","Lock_tables_priv":"N","References_priv":"N","Select_priv":"Y","Show_view_priv":"N","Trigger_priv":"N","Update_priv":"N","User":"mysql.session"}
{"Alter_priv":"N","Alter_routine_priv":"N","Create_priv":"N","Create_routine_priv":"N","Create_tmp_table_priv":"N","Create_view_priv":"N","Db":"sys","Delete_priv":"N","Drop_priv":"N","Event_priv":"N","Execute_priv":"N","Grant_priv":"N","Host":"localhost","Index_priv":"N","Insert_priv":"N","Lock_tables_priv":"N","References_priv":"N","Select_priv":"N","Show_view_priv":"N","Trigger_priv":"Y","Update_priv":"N","User":"mysql.sys"}
```
