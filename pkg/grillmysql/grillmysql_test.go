package grillmysql

import (
	"testing"

	"bitbucket.org/swigy/grill"
)

const createTaskTableQuery = `CREATE TABLE task (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		name VARCHAR(20),
		completed bool,
		description text,
		PRIMARY KEY (id)
	)`

func Test_GrillMysql(t *testing.T) {
	helper, err := Start()
	if err != nil {
		t.Errorf("error starting mysql grill, error=%v", err)
		return
	}

	tests := []grill.TestCase{
		{
			Name: "Test_SeedAndCount",
			Stubs: []grill.Stub{
				helper.CreateTable(createTaskTableQuery),
				helper.SeedDataFromCSVFile("task", "test_data/example.csv"),
			},
			Action: func() interface{} {
				return nil
			},
			Assertions: []grill.Assertion{
				helper.AssertCount("task", 2),
			},
			Cleaners: []grill.Cleaner{
				helper.DeleteTable("task"),
			},
		},
	}

	grill.Run(t, tests)
}
