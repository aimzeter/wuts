package mysql_test

import (
	"testing"

	zone "github.com/aimzeter/wuts/four"
	"github.com/aimzeter/wuts/four/mysql"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateStudent(t *testing.T) {
	db, mock, closedb := setupMockDB(t)
	defer closedb()

	repo := mysql.NewStudentRepo(db)

	tests := []struct {
		name       string
		newStudent zone.Student
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
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectExec("INSERT INTO students").WithArgs(
				tc.newStudent.PersonalInfo.NIK, tc.newStudent.PersonalInfo.Name, tc.newStudent.PersonalInfo.Address,
				tc.newStudent.PersonalInfo.Coord.Lat, tc.newStudent.PersonalInfo.Coord.Long,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Create(tc.newStudent)
			if err != nil {
				t.Fatalf("Fail to create student, got error %s", err.Error())
			}
		})
	}
}
