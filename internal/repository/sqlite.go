package repository

import (
	"Ewallet-infotecs/internal/seeds"
	"database/sql"
	"github.com/sirupsen/logrus"
)

func NewSqliteDb(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("failed to ping db: %s", err.Error())
	}
	seeds.WalletSeed(db)
	return db, nil
}
