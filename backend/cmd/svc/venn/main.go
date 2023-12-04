package main

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/database"
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/router"
)

func main() {
	// db
	mysqlConf := &database.MySQLConfig{}
	mysqlConf.SetDefaults("venn")
	database.ConfigMySQL(mysqlConf)

	db, err := database.New(mysqlConf)

	if err != nil {
		panic(err)
	}
	loader := database.NewLoader(db)
	err = loader.Migrate()

	if err != nil {
		panic(err)
	}

	r := router.Setup(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
