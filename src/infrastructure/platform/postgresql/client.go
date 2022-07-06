package postgresql

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func Init(url string) Client {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(dao.Route{})

	return Client{db}
}
