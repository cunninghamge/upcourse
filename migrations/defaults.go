package main

import "upcourse/models"

var defaultActivities = []models.Activity{
	{
		Name:        "Reading (understand)",
		Description: "130 wpm; 10 pages per hour",
		Metric:      "# of pages",
		Multiplier:  6,
		Custom:      false,
	}, {
		Name:        "Reading (study guide)",
		Description: "65 wpm; 5 pages per hour",
		Metric:      "# of pages",
		Multiplier:  12,
		Custom:      false,
	}, {
		Name:        "Writing (research)",
		Description: "6 hours per page (500 words, single-spaced)",
		Metric:      "'# of pages",
		Multiplier:  360,
		Custom:      false,
	}, {
		Name:        "Writing (reflection)",
		Description: "90 minutes per page (500 words, single-spaced)",
		Metric:      "'# of pages",
		Multiplier:  90,
		Custom:      false,
	}, {
		Name:        "Learning Objects (matching/multiple choice)",
		Description: "10 minutes per object",
		Metric:      "# of LOs",
		Multiplier:  10,
		Custom:      false,
	}, {
		Name:        "Learning Objects (case study)",
		Description: "20 minutes per object",
		Metric:      "# of LOs",
		Multiplier:  20,
		Custom:      false,
	}, {
		Name:        "Lecture",
		Description: "Factor 1.25x the actual lecture runtime",
		Metric:      "# of minutes",
		Multiplier:  1.25,
		Custom:      false,
	}, {
		Name:        "Videos",
		Description: "Factor the full length of video",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	}, {
		Name:        "Websites",
		Description: "10-20 minutes",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	}, {
		Name:        "Discussion Boards",
		Description: "250 words/60 minutes for initial post or 2 replies",
		Metric:      "# of discussion boards",
		Multiplier:  60,
		Custom:      false,
	}, {
		Name:        "Quizzes",
		Description: "Average 1.5 minutes per question",
		Metric:      "# of questions",
		Multiplier:  1.5,
		Custom:      false,
	}, {
		Name:        "Exams",
		Description: "Average 1.5 minutes per question",
		Metric:      "# of questions",
		Multiplier:  1.5,
		Custom:      false,
	}, {
		Name:        "Self Assessments",
		Description: "Average 1 minute per question",
		Metric:      "# of questions",
		Multiplier:  1,
		Custom:      false,
	}, {
		Name:        "Miscellaneous",
		Description: "any additional assignments not listed",
		Metric:      "# of minutes",
		Multiplier:  1,
		Custom:      false,
	},
}

var sampleCourse = models.Course{
	Name:        "Foundations of Nursing",
	Institution: "Colorado Nursing College",
	CreditHours: 3,
	Length:      8,
	Modules: []models.Module{
		{
			Name:   "Module 1",
			Number: 1,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      107,
					ActivityId: 1,
				}, {
					Input:      6,
					ActivityId: 2,
				}, {
					Input:      7,
					ActivityId: 5,
				}, {
					Input:      95,
					ActivityId: 8,
				}, {
					Input:      1,
					ActivityId: 10,
				}, {
					Input:      450,
					ActivityId: 11,
				}, {
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			Name:   "Module 2",
			Number: 2,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      53,
					ActivityId: 1,
				}, {
					Input:      5,
					ActivityId: 2,
				}, {
					Input:      5,
					ActivityId: 5,
				}, {
					Input:      71,
					ActivityId: 8,
				}, {
					Input:      1,
					ActivityId: 10,
				}, {
					Input:      100,
					ActivityId: 11,
				},
			},
		}, {
			Name:   "Module 3",
			Number: 3,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      66,
					ActivityId: 1,
				}, {
					Input:      4,
					ActivityId: 2,
				}, {
					Input:      1,
					ActivityId: 4,
				}, {
					Input:      4,
					ActivityId: 5,
				}, {
					Input:      2,
					ActivityId: 6,
				}, {
					Input:      86,
					ActivityId: 8,
				}, {
					Input:      1,
					ActivityId: 10,
				}, {
					Input:      240,
					ActivityId: 11,
				}, {
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			Name:   "Module 4",
			Number: 4,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      105,
					ActivityId: 1,
				}, {
					Input:      7,
					ActivityId: 2,
				}, {
					Input:      2,
					ActivityId: 4,
				}, {
					Input:      3,
					ActivityId: 5,
				}, {
					Input:      75,
					ActivityId: 8,
				}, {
					Input:      390,
					ActivityId: 11,
				}, {
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			Name:   "Module 5",
			Number: 5,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      52,
					ActivityId: 1,
				}, {
					Input:      5,
					ActivityId: 2,
				}, {
					Input:      1,
					ActivityId: 4,
				}, {
					Input:      5,
					ActivityId: 5,
				}, {
					Input:      1,
					ActivityId: 6,
				}, {
					Input:      62,
					ActivityId: 8,
				}, {
					Input:      1,
					ActivityId: 10,
				}, {
					Input:      300,
					ActivityId: 11,
				},
			},
		}, {
			Name:   "Module 6",
			Number: 6,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      36,
					ActivityId: 1,
				}, {
					Input:      5,
					ActivityId: 2,
				}, {
					Input:      5,
					ActivityId: 5,
				}, {
					Input:      1,
					ActivityId: 6,
				}, {
					Input:      40,
					ActivityId: 8,
				}, {
					Input:      1,
					ActivityId: 10,
				}, {
					Input:      90,
					ActivityId: 11,
				}, {
					Input:      50,
					ActivityId: 13,
				},
			},
		}, {
			Name:   "Module 7",
			Number: 7,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      88,
					ActivityId: 1,
				}, {
					Input:      5,
					ActivityId: 2,
				}, {
					Input:      4,
					ActivityId: 5,
				}, {
					Input:      2,
					ActivityId: 6,
				}, {
					Input:      42,
					ActivityId: 8,
				}, {
					Input:      240,
					ActivityId: 11,
				},
			},
		}, {
			Name:   "Module 8",
			Number: 8,
			ModuleActivities: []models.ModuleActivity{
				{
					Input:      3,
					ActivityId: 3,
				}, {
					Input:      100,
					ActivityId: 13,
				},
			},
		},
	},
}
