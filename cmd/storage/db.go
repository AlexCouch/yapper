package storage

import (
	"log"
	"os"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"yapper.com/m/yapper/cmd/models"
)

func createSchema(db *pg.DB) error{
    models := []interface{}{
        (*models.User)(nil),
        (*models.ProfileImage)(nil),
    }

    for _, model := range models{
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: true,
            IfNotExists: true,
        })
        if err != nil{
            return err
        }
    }
    return nil
}

func GetDB() (*pg.DB, error){
    err := godotenv.Load()
    if err != nil{
        log.Fatal("ERror loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    db := pg.Connect(&pg.Options{
        User: dbUser,
        Password: dbPass,
        Addr: dbHost + ":" + dbPort,
        Database: dbName,
    })
    db.AddQueryHook(pgdebug.DebugHook{
        Verbose: true,
    })
    //err = createSchema(db)
    //if err != nil{
    //    return nil, err
    //}
    return db, nil
}
