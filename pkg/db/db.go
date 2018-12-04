package db

import (
	"log"
	"sync"

	"git.zam.io/microservices/customer-api/pkg/config"
	"github.com/go-pg/pg"
)

var db *pg.DB
var onceInit sync.Once

func DB() *pg.DB {
	onceInit.Do(func() {
		conf := config.Config()
		dbOpts := &pg.Options{
			Addr:     conf.GetString("database.address"),
			User:     conf.GetString("database.user"),
			Password: conf.GetString("database.password"),
			Database: conf.GetString("database.database"),
		}
		db = pg.Connect(dbOpts)
		_, err := db.Exec("SELECT 1")
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}

func CheckDB() error {
	_, err := DB().Exec("SELECT 1")
	return err
}
