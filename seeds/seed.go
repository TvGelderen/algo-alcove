package seeds

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TvGelderen/algo-alcove/database"
)

func Seed(dbConnectionString string) {
	fmt.Println("Seeding database")

	connection, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}

	db := database.New(connection)

    seedSortingAlgorithms(db)
    seedPathFindingAlgorithms(db)

    seedBubbleSort(db)
}
