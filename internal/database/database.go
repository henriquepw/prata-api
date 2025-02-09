package database

import (
	"log"
	"sync"

	"github.com/henriquepw/pobrin-api/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/go-libsql"
)

var (
	dbConn  *sqlx.DB
	dbMutex = &sync.Mutex{}
)

func GetDBConnection(m ...Migration) (*sqlx.DB, error) {
	if dbConn == nil {
		dbMutex.Lock()
		defer dbMutex.Unlock()

		if dbConn == nil {
			db, err := startDB(m...)
			if err != nil {
				return nil, err
			}

			dbConn = db
		}
	}

	return dbConn, nil
}

func startDB(migrations ...Migration) (*sqlx.DB, error) {
	db, err := sqlx.Open("libsql", config.Env().DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %s", err.Error())

		return nil, err
	}

	for _, m := range migrations {
		err := m(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
