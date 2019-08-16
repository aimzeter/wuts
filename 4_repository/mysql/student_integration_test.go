// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	zone "github.com/aimzeter/wuts/4_repository"
	"github.com/aimzeter/wuts/4_repository/mysql"
)

func TestICreateStudent(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	repo := mysql.NewStudentRepo(db)
	seeds := []zone.Student{
		{
			PersonalInfo: zone.Citizen{
				NIK:     "1",
				Name:    "Foo1",
				Address: "Bar1",
				Coord:   zone.Coord{1, 1},
			},
		},
		{
			PersonalInfo: zone.Citizen{
				NIK:     "2",
				Name:    "Foo2",
				Address: "Bar2",
				Coord:   zone.Coord{2, 2},
			},
		},
	}

	tests := []struct {
		name       string
		newStudent zone.Student

		wantGotStudent  zone.Student
		wantAllStudents []zone.Student
		wantTotalCount  int
	}{
		{
			name: "create new student",
			newStudent: zone.Student{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0.1, 0.2},
				},
			},
			wantGotStudent: zone.Student{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0.1, 0.2},
				},
			},
			wantAllStudents: []zone.Student{
				{
					PersonalInfo: zone.Citizen{
						NIK:     "1",
						Name:    "Foo1",
						Address: "Bar1",
						Coord:   zone.Coord{1, 1},
					},
				},
				{
					PersonalInfo: zone.Citizen{
						NIK:     "2",
						Name:    "Foo2",
						Address: "Bar2",
						Coord:   zone.Coord{2, 2},
					},
				},
				{
					PersonalInfo: zone.Citizen{
						NIK:     "1234567890",
						Name:    "Foo",
						Address: "Bar",
						Coord:   zone.Coord{0.1, 0.2},
					},
				},
			},
			wantTotalCount: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Seed(seeds)
			defer repo.RemoveAll()

			err := repo.Create(tc.newStudent)
			if err != nil {
				t.Fatalf("Fail to create student, got error %s", err.Error())
			}

			got := repo.Get(tc.newStudent.PersonalInfo.NIK)
			assert.Equal(t, tc.wantGotStudent, got)

			all := repo.All()
			assert.Equal(t, tc.wantAllStudents, all)

			count := repo.CountAll()
			assert.Equal(t, tc.wantTotalCount, count)
		})
	}
}
