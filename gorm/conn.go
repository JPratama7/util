package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	grm "gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func NewMySQLConnection(connString string, cfg ...grm.Option) (db *grm.DB, err error) {
	dial := mysql.Open(connString)

	if len(cfg) == 0 {
		cfg = append(cfg, &grm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}

	db, err = grm.Open(dial, cfg...)
	if err != nil {
		return
	}

	dbO, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	dbO.SetMaxIdleConns(2)
	dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)

	db.Statement.RaiseErrorOnNotFound = true

	return
}

func NewConnection(dial grm.Dialector, cfg ...grm.Option) (db *grm.DB, err error) {
	if len(cfg) == 0 {
		cfg = append(cfg, &grm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}

	db, err = grm.Open(dial, cfg...)
	if err != nil {
		return
	}

	dbO, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	dbO.SetMaxIdleConns(2)
	dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)
	db.Statement.RaiseErrorOnNotFound = true
	return
}

func NewPostgresConnection(connString string, cfg ...grm.Option) (db *grm.DB, err error) {
	dial := postgres.Open(connString)

	if len(cfg) == 0 {
		cfg = append(cfg, &grm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}

	db, err = grm.Open(dial, cfg...)
	if err != nil {
		return
	}

	dbO, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	dbO.SetMaxIdleConns(2)
	dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)
	db.Statement.RaiseErrorOnNotFound = true

	return
}
