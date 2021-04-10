package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table activities...")
		_, err := db.Exec(`CREATE TABLE activities(
id SERIAL,
name VARCHAR(255),
description VARCHAR(255),
metric VARCHAR(255),
multiplier INT,
custom BOOLEAN,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id)
)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table activities...")
		_, err := db.Exec(`DROP TABLE activities`)
		return err
	})
}
