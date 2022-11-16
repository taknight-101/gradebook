package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Ahmed",
			LastName:  "Ibrahim",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Diaa",
			LastName:  "Mohamed",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 100,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
	}
}
