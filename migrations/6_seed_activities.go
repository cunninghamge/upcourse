package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("seeding databases...")
		_, err := db.Exec(`DELETE FROM activities;
		INSERT INTO activities(id, name, description, metric, multiplier, custom)
		VALUES (1, 'Reading (understand)', '130 wpm; 10 pages per hour', '# of pages', 6, FALSE),
			(2, 'Reading (study guide)', '65 wpm; 5 pages per hour', '# of pages', 12, FALSE),
			(3, 'Writing (research)', '6 hours per page (500 words, single-spaced)', '# of pages', 360, FALSE),
			(4, 'Writing (reflection)', '90 minutes per page (500 words, single-spaced)', '# of pages', 90, FALSE),
			(5, 'Learning Objects (matching/multiple choice)', '10 minutes per object', '# of LOs', 10, FALSE),
			(6, 'Learning Objects (case study)', '20 minutes per object', '# of LOs', 20, FALSE),
			(7, 'Lecture', 'Factor 1.25x the actual lecture runtime', '# of minutes', 1.25, FALSE),
			(8, 'Videos', 'Factor the full length of video', '# of minutes', 1, FALSE),
			(9, 'Websites', '10-20 minutes', '', 1, FALSE),
			(10, 'Discussion Boards', '250 words/60 minutes for initial post or 2 replies', '# of discussion boards', 60, FALSE),
			(11, 'Quizzes', 'Average 1.5 minutes per question', '# of questions', 1.5, FALSE),
			(12, 'Exams', 'Average 1.5 minutes per question', '# of questions', 1.5, FALSE),
			(13, 'Self Assessments', 'Average 1 minute per question', '# of questions', 1, FALSE),
			(14, 'Miscellaneous', 'any additional assignments not listed', '', 1, FALSE);
			`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("removing seed data...")
		_, err := db.Exec(`DELETE FROM activities;`)
		return err
	})
}
