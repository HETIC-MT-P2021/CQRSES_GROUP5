package database

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// MakeMigrations executes the sql migration files
func MakeMigrations() {
	file, err := ioutil.ReadFile("database/migrations.sql")

	if err != nil {
		fmt.Println("Error while opening migrations file: ", err.Error())
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := DB.Exec(request)
		if err != nil {
			fmt.Println("Error during migrations: ", err.Error())
		}
	}

}
