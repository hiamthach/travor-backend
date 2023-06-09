package db

import (
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetInstance(path string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		DB, err = gorm.Open(postgres.Open(path), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return DB, nil
}
