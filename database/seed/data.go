package main

import "upcourse/models"

var defaultActivities = []models.Activity{
	{
		ID:          1,
		Name:        "Reading (understand)",
		Description: "130 wpm; 10 pages per hour",
		Metric:      "# of pages",
		Multiplier:  6,
		Custom:      false,
	}, {
		ID:          2,
		Name:        "Reading (study guide)",
		Description: "65 wpm; 5 pages per hour",
		Metric:      "# of pages",
		Multiplier:  12,
		Custom:      false,
	}, {
		ID:          3,
		Name:        "Writing (research)",
		Description: "6 hours per page (500 words, single-spaced)",
		Metric:      "'# of pages",
		Multiplier:  360,
		Custom:      false,
	}, {
		ID:          4,
		Name:        "Writing (reflection)",
		Description: "90 minutes per page (500 words, single-spaced)",
		Metric:      "'# of pages",
		Multiplier:  90,
		Custom:      false,
	}, {
		ID:          5,
		Name:        "Learning Objects (matching/multiple choice)",
		Description: "10 minutes per object",
		Metric:      "# of LOs",
		Multiplier:  10,
		Custom:      false,
	}, {
		ID:          6,
		Name:        "Learning Objects (case study)",
		Description: "20 minutes per object",
		Metric:      "# of LOs",
		Multiplier:  20,
		Custom:      false,
	}, {
		ID:          7,
		Name:        "Lecture",
		Description: "Factor 1.25x the actual lecture runtime",
		Metric:      "# of minutes",
		Multiplier:  1.25,
		Custom:      false,
	}, {
		ID:          8,
		Name:        "Videos",
		Description: "Factor the full length of video",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	}, {
		ID:          9,
		Name:        "Websites",
		Description: "10-20 minutes",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	}, {
		ID:          10,
		Name:        "Discussion Boards",
		Description: "250 words/60 minutes for initial post or 2 replies",
		Metric:      "# of discussion boards",
		Multiplier:  60,
		Custom:      false,
	}, {
		ID:          11,
		Name:        "Quizzes",
		Description: "Average 1.5 minutes per question",
		Metric:      "# of questions",
		Multiplier:  1.5,
		Custom:      false,
	}, {
		ID:          12,
		Name:        "Exams",
		Description: "Average 1.5 minutes per question",
		Metric:      "# of questions",
		Multiplier:  1.5,
		Custom:      false,
	}, {
		ID:          13,
		Name:        "Self Assessments",
		Description: "Average 1 minute per question",
		Metric:      "# of questions",
		Multiplier:  1,
		Custom:      false,
	}, {
		ID:          14,
		Name:        "Miscellaneous",
		Description: "any additional assignments not listed",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	},
}

var sampleCourse = models.Course{
	ID:          1,
	Name:        "Foundations of Nursing",
	Institution: "Colorado Nursing College",
	CreditHours: 3,
	Length:      8,
	Modules: []*models.Module{
		{
			ID:       1,
			Name:     "Module 1",
			Number:   1,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   1,
					Input:      107,
					ActivityId: 1,
				}, {
					ModuleId:   1,
					Input:      6,
					ActivityId: 2,
				}, {
					ModuleId:   1,
					Input:      7,
					ActivityId: 5,
				}, {
					ModuleId:   1,
					Input:      95,
					ActivityId: 8,
				}, {
					ModuleId:   1,
					Input:      1,
					ActivityId: 10,
				}, {
					ModuleId:   1,
					Input:      450,
					ActivityId: 11,
				}, {
					ModuleId:   1,
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			ID:       2,
			Name:     "Module 2",
			Number:   2,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   2,
					Input:      53,
					ActivityId: 1,
				}, {
					ModuleId:   2,
					Input:      5,
					ActivityId: 2,
				}, {
					ModuleId:   2,
					Input:      5,
					ActivityId: 5,
				}, {
					ModuleId:   2,
					Input:      71,
					ActivityId: 8,
				}, {
					ModuleId:   2,
					Input:      1,
					ActivityId: 10,
				}, {
					ModuleId:   2,
					Input:      100,
					ActivityId: 11,
				},
			},
		}, {
			ID:       3,
			Name:     "Module 3",
			Number:   3,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   3,
					Input:      66,
					ActivityId: 1,
				}, {
					ModuleId:   3,
					Input:      4,
					ActivityId: 2,
				}, {
					ModuleId:   3,
					Input:      1,
					ActivityId: 4,
				}, {
					ModuleId:   3,
					Input:      4,
					ActivityId: 5,
				}, {
					ModuleId:   3,
					Input:      2,
					ActivityId: 6,
				}, {
					ModuleId:   3,
					Input:      86,
					ActivityId: 8,
				}, {
					ModuleId:   3,
					Input:      1,
					ActivityId: 10,
				}, {
					ModuleId:   3,
					Input:      240,
					ActivityId: 11,
				}, {
					ModuleId:   3,
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			ID:       4,
			Name:     "Module 4",
			Number:   4,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   4,
					Input:      105,
					ActivityId: 1,
				}, {
					ModuleId:   4,
					Input:      7,
					ActivityId: 2,
				}, {
					ModuleId:   4,
					Input:      2,
					ActivityId: 4,
				}, {
					ModuleId:   4,
					Input:      3,
					ActivityId: 5,
				}, {
					ModuleId:   4,
					Input:      75,
					ActivityId: 8,
				}, {
					ModuleId:   4,
					Input:      390,
					ActivityId: 11,
				}, {
					ModuleId:   4,
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			ID:       5,
			Name:     "Module 5",
			Number:   5,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   5,
					Input:      52,
					ActivityId: 1,
				}, {
					ModuleId:   5,
					Input:      5,
					ActivityId: 2,
				}, {
					ModuleId:   5,
					Input:      1,
					ActivityId: 4,
				}, {
					ModuleId:   5,
					Input:      5,
					ActivityId: 5,
				}, {
					ModuleId:   5,
					Input:      1,
					ActivityId: 6,
				}, {
					ModuleId:   5,
					Input:      62,
					ActivityId: 8,
				}, {
					ModuleId:   5,
					Input:      1,
					ActivityId: 10,
				}, {
					ModuleId:   5,
					Input:      300,
					ActivityId: 11,
				},
			},
		}, {
			ID:       6,
			Name:     "Module 6",
			Number:   6,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   6,
					Input:      36,
					ActivityId: 1,
				}, {
					ModuleId:   6,
					Input:      5,
					ActivityId: 2,
				}, {
					ModuleId:   6,
					Input:      5,
					ActivityId: 5,
				}, {
					ModuleId:   6,
					Input:      1,
					ActivityId: 6,
				}, {
					ModuleId:   6,
					Input:      40,
					ActivityId: 8,
				}, {
					ModuleId:   6,
					Input:      1,
					ActivityId: 10,
				}, {
					ModuleId:   6,
					Input:      90,
					ActivityId: 11,
				}, {
					ModuleId:   6,
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			ID:       7,
			Name:     "Module 7",
			Number:   7,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   7,
					Input:      88,
					ActivityId: 1,
				}, {
					ModuleId:   7,
					Input:      5,
					ActivityId: 2,
				}, {
					ModuleId:   7,
					Input:      4,
					ActivityId: 5,
				}, {
					ModuleId:   7,
					Input:      2,
					ActivityId: 6,
				}, {
					ModuleId:   7,
					Input:      42,
					ActivityId: 8,
				}, {
					ModuleId:   7,
					Input:      240,
					ActivityId: 11,
				},
			},
		}, {
			ID:       8,
			Name:     "Module 8",
			Number:   8,
			CourseId: 1,
			ModuleActivities: []*models.ModuleActivity{
				{
					ModuleId:   8,
					Input:      3,
					ActivityId: 3,
				}, {
					ModuleId:   8,
					Input:      100,
					ActivityId: 13,
				},
			},
		},
	},
}
