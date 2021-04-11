package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table courses...")
		_, err := db.Exec(`CREATE TABLE courses(
id SERIAL,
name VARCHAR(255),
institution VARCHAR(255),
credit_hours INT,
length INT,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id)
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table courses...")
		_, err := db.Exec(`DROP TABLE courses`)
		return err
	})
}
