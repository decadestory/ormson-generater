package dbstruct

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Unknwon/goconfig"
)

//GetAtomLines ...
func GetAtomLines(name string) string {

	c, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatal("Config load faild:", err.Error())
	}

	server, _ := c.GetValue("dbconf", "server")
	user, _ := c.GetValue("dbconf", "user")
	pwd, _ := c.GetValue("dbconf", "pwd")
	port, _ := c.GetValue("dbconf", "port")
	database, _ := c.GetValue("dbconf", "database")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", server, user, pwd, port, database)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return ""
	}

	defer db.Close()

	sql := `select 
                cast(ep.value as varchar) [Description],
                (case when  ty.[name] in ('text','ntext' ,'char','nchar', 'varchar', 'nvarchar') then 'string'
                when ty.[name] in ('date' , 'datetime' , 'datetime2') then 'DateTime'
                when ty.[name] in ('bit') then 'bool'
                when ty.[name] in ('smallint') then 'short'
                when ty.[name] in ('bigint') then 'long'
                when ty.[name] in ('real') then 'float'
                when ty.[name] in ('float') then 'double'
                when ty.[name] in ('money') then 'decimal'
                when ty.[name] in ('uniqueidentifier') then 'Guid'
                else ty.[name] end) as TypeName,
                (case c.[is_nullable] when 1 then case when ty.[name] not in('text','ntext' ,'char','nchar', 'varchar', 'nvarchar') then '?' else '' end
                 else '' end) as NullString,c.name+' {set;get;}' as EndString from sys.tables t
                INNER JOIN sys.columns c ON t.object_id = c.object_id
                LEFT JOIN sys.extended_properties ep ON ep.major_id = c.object_id AND ep.minor_id = c.column_id 
                left JOIN sys.types ty on ty.[system_type_id]=c.[user_type_id] and ty.[name]!='sysname'
                WHERE ep.class =1 AND t.name='{0}' or c.object_id=Object_Id('{0}')`
	sql = strings.Replace(sql, "{0}", name, -1)
	text := exec(db, sql)
	if text != "" {
		return text
	}

	return ""
}

func exec(db *sql.DB, cmd string) string {
	rows, err := db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}

	text := ""
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}

		line := ""

		for i := 0; i < len(vals); i++ {
			var t = printValue(vals[i].(*interface{}))

			if i == 0 && t != "" {
				line += "\t\t/// <summary>\n\t\t/// {0}\n\t\t/// </summary>\n\t\tpublic"
				line = strings.Replace(line, "{0}", t, -1)
			} else if i == 0 && t == "" {
				line = "\t\tpublic"
			} else if i == 2 {
				line += t
			} else {
				line += " " + t
			}

		}
		text += line + "\n\n"
	}

	if rows.Err() != nil {
		fmt.Println(rows.Err())
		return ""
	}

	return text
}

func printValue(pval *interface{}) string {
	switch v := (*pval).(type) {
	case nil:
	case time.Time:
	case bool:
		return ""
	case []byte:
		return string(v)
	default:
		return v.(string)
	}
	return ""
}
