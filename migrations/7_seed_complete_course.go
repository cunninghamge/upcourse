package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("seeding database...")
		_, err := db.Exec(`INSERT INTO courses(id, name, institution, credit_hours, length)
			VALUES (1, 'Foundations of Nursing', 'Colorado Nursing College', 3, 8);
		INSERT INTO modules(id, name, number, course_id)
			VALUES (1, 'Module 1', 1, 1),
			(2, 'Module 2', 2, 1),
			(3, 'Module 3', 3, 1),
			(4, 'Module 4', 4, 1),
			(5, 'Module 5', 5, 1),
			(6, 'Module 6', 6, 1),
			(7, 'Module 7', 7, 1),
			(8, 'Module 8', 8, 1);
		INSERT INTO module_activities(input, module_id, activity_id)
			VALUES 
			(107, 1, 1),
			(6, 1, 2),
			(7, 1, 5),
			(95, 1, 8),
			(1, 1, 10),
			(450, 1, 11),
			(50, 1, 13),
			( 53, 2, 1),
			( 5, 2, 2),
			( 5, 2, 5),
			( 71, 2, 8),
			( 1, 2, 10),
			( 100, 2, 11),
			( 66, 3, 1),
			( 4, 3, 2),
			( 1, 3, 4),
			( 4, 3, 5),
			( 2, 3, 6),
			( 86, 3, 8),
			( 1, 3, 10),
			( 240, 3, 11),
			( 50, 3, 13),
			( 105, 4, 1),
			( 7, 4, 2),
			( 2, 4, 4),
			( 3, 4, 5),
			( 75, 4, 8),
			( 390, 4, 11),
			( 50, 4, 13),
			( 52, 5, 1),
			( 5, 5, 2),
			( 1, 5, 4),
			( 5, 5, 5),
			( 1, 5, 6),
			( 62, 5, 8),
			( 1, 5, 10),
			( 300, 5, 11),
			( 36, 6, 1),
			( 5, 6, 2),
			( 5, 6, 5),
			( 1, 6, 6),
			( 40, 6, 8),
			( 1, 6, 10),
			( 90, 6, 11),
			( 50, 6, 13),
			( 88, 7, 1),
			( 5, 7, 2),
			( 4, 7, 5),
			( 2, 7, 6),
			( 42, 7, 8),
			( 240, 7, 11),
			( 3, 8, 3),
			( 100, 8, 13);
			`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("deleting seed data...")
		_, err := db.Exec("")
		return err
	})
}
