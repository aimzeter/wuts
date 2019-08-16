package mysql_test

import (
	"database/sql"
	"testing"

	zone "github.com/aimzeter/wuts/4_repository"
	"github.com/aimzeter/wuts/4_repository/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetParticipant(t *testing.T) {
	db, mock, closedb := setupMockDB(t)
	defer closedb()

	repo := mysql.NewParticipantRepo(db)
	seeds := []zone.Participant{
		{
			PersonalInfo: zone.Citizen{
				NIK:     "1",
				Name:    "Foo1",
				Address: "Bar1",
				Coord:   zone.Coord{1, 1},
			},
			AutoReg:    true,
			Distance:   10.5,
			TotalScore: 70,
		},
		{
			PersonalInfo: zone.Citizen{
				NIK:     "2",
				Name:    "Foo2",
				Address: "Bar2",
				Coord:   zone.Coord{2, 2},
			},
			AutoReg:    false,
			Distance:   20.3,
			TotalScore: 100,
		},
	}

	mockRows := sqlmock.NewRows([]string{"nik", "name", "address", "latitude", "longitude", "auto_reg", "distance", "total_score"})
	seedParticipant(mockRows, seeds)
	mockRows.RowError(1, sql.ErrNoRows)

	tests := []struct {
		name string
		nik  string

		wantParticipant zone.Participant
	}{
		{
			name: "existed doc",
			nik:  "1",
			wantParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "1",
					Name:    "Foo1",
					Address: "Bar1",
					Coord:   zone.Coord{1, 1},
				},
				AutoReg:    true,
				Distance:   10.5,
				TotalScore: 70,
			},
		},
		{
			name:            "not existed doc",
			nik:             "99",
			wantParticipant: zone.Participant{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery("^SELECT (.+) FROM participants (.+)").WillReturnRows(mockRows)
			got := repo.Get(tc.nik)
			assert.Equal(t, tc.wantParticipant, got)
		})
	}
}

func TestCreateParticipant(t *testing.T) {
	db, mock, closedb := setupMockDB(t)
	defer closedb()

	repo := mysql.NewParticipantRepo(db)

	tests := []struct {
		name           string
		newParticipant zone.Participant
	}{
		{
			name: "create new participant",
			newParticipant: zone.Participant{
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
			mock.ExpectExec("INSERT INTO participants").WithArgs(
				tc.newParticipant.PersonalInfo.NIK, tc.newParticipant.PersonalInfo.Name, tc.newParticipant.PersonalInfo.Address,
				tc.newParticipant.PersonalInfo.Coord.Lat, tc.newParticipant.PersonalInfo.Coord.Long,
				tc.newParticipant.AutoReg, tc.newParticipant.Distance, tc.newParticipant.TotalScore,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Create(tc.newParticipant)
			if err != nil {
				t.Fatalf("Fail to create participant, got error %s", err.Error())
			}
		})
	}
}

func TestUpdateParticipant(t *testing.T) {
	db, mock, closedb := setupMockDB(t)
	defer closedb()

	repo := mysql.NewParticipantRepo(db)

	tests := []struct {
		name           string
		newParticipant zone.Participant
	}{
		{
			name: "update new participant",
			newParticipant: zone.Participant{
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
			mock.ExpectExec("UPDATE participants SET").WithArgs(
				tc.newParticipant.PersonalInfo.NIK, tc.newParticipant.PersonalInfo.Name, tc.newParticipant.PersonalInfo.Address,
				tc.newParticipant.PersonalInfo.Coord.Lat, tc.newParticipant.PersonalInfo.Coord.Long,
				tc.newParticipant.AutoReg, tc.newParticipant.Distance, tc.newParticipant.TotalScore,
				tc.newParticipant.PersonalInfo.NIK,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Update(tc.newParticipant)
			if err != nil {
				t.Fatalf("Fail to update participant, got error %s", err.Error())
			}
		})
	}
}

func seedParticipant(rows *sqlmock.Rows, seeds []zone.Participant) {
	for _, p := range seeds {
		rows.AddRow(
			p.PersonalInfo.NIK, p.PersonalInfo.Name, p.PersonalInfo.Address, p.PersonalInfo.Coord.Lat, p.PersonalInfo.Coord.Long,
			p.AutoReg, p.Distance, p.TotalScore,
		)
	}
}
