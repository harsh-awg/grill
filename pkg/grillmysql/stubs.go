package grillmysql

import (
	"fmt"

	"bitbucket.org/swigy/grill"
	"github.com/go-sql-driver/mysql"
)

func (gm *GrillMysql) CreateTable(query string) grill.Stub {
	return grill.StubFunc(func() error {
		_, err := gm.Client().Exec(query)
		return err
	})
}

func (gm *GrillMysql) SeedDataFromCSVFile(tableName string, filePath string) grill.Stub {
	return grill.StubFunc(func() error {
		mysql.RegisterLocalFile(filePath)
		_, err := gm.Client().Exec(fmt.Sprintf(`LOAD DATA LOCAL INFILE  '%s'
		INTO TABLE %s
		FIELDS TERMINATED BY ','
		ENCLOSED BY '"'
		LINES TERMINATED BY '\n'
		IGNORE 1 ROWS;`, filePath, tableName),
		)
		return err
	})
}
