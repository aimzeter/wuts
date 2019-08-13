// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	zone "github.com/aimzeter/wuts/four"
	"github.com/aimzeter/wuts/four/mysql"
)

func TestIGetParticipant(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

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
			repo.Seed(seeds)
			defer repo.RemoveAll()

			got := repo.Get(tc.nik)
			assert.Equal(t, tc.wantParticipant, got)
		})
	}
}

func TestICreateParticipant(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

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

	tests := []struct {
		name           string
		newParticipant zone.Participant

		wantGotParticipant  zone.Participant
		wantAllParticipants []zone.Participant
		wantTotalCount      int
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
				AutoReg:    true,
				Distance:   5.3,
				TotalScore: 85.5,
			},
			wantGotParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0.1, 0.2},
				},
				AutoReg:    true,
				Distance:   5.3,
				TotalScore: 85.5,
			},
			wantAllParticipants: []zone.Participant{
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
				{
					PersonalInfo: zone.Citizen{
						NIK:     "1234567890",
						Name:    "Foo",
						Address: "Bar",
						Coord:   zone.Coord{0.1, 0.2},
					},
					AutoReg:    true,
					Distance:   5.3,
					TotalScore: 85.5,
				},
			},
			wantTotalCount: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Seed(seeds)
			defer repo.RemoveAll()

			err := repo.Create(tc.newParticipant)
			if err != nil {
				t.Fatalf("Fail to create participant, got error %s", err.Error())
			}

			got := repo.Get(tc.newParticipant.PersonalInfo.NIK)
			assert.Equal(t, tc.wantGotParticipant, got)

			all := repo.All()
			assert.Equal(t, tc.wantAllParticipants, all)

			count := repo.CountAll()
			assert.Equal(t, tc.wantTotalCount, count)
		})
	}
}

func TestIUpdateParticipant(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

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

	tests := []struct {
		name           string
		newParticipant zone.Participant

		wantGotParticipant  zone.Participant
		wantAllParticipants []zone.Participant
		wantTotalCount      int
	}{
		{
			name: "existed participant",
			newParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "1",
					Name:    "NewFoo1",
					Address: "NewBar1",
					Coord:   zone.Coord{12.3, -4.5},
				},
				AutoReg:    false,
				Distance:   20.7,
				TotalScore: 100,
			},
			wantGotParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "1",
					Name:    "NewFoo1",
					Address: "NewBar1",
					Coord:   zone.Coord{12.3, -4.5},
				},
				AutoReg:    false,
				Distance:   20.7,
				TotalScore: 100,
			},
			wantAllParticipants: []zone.Participant{
				{
					PersonalInfo: zone.Citizen{
						NIK:     "1",
						Name:    "NewFoo1",
						Address: "NewBar1",
						Coord:   zone.Coord{12.3, -4.5},
					},
					AutoReg:    false,
					Distance:   20.7,
					TotalScore: 100,
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
			},
			wantTotalCount: 2,
		},
		{
			name: "not existed participant",
			newParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "3",
					Name:    "Foo3",
					Address: "Foo3",
					Coord:   zone.Coord{-7.3, -14.5},
				},
				AutoReg:    false,
				Distance:   14,
				TotalScore: 100,
			},
			wantGotParticipant:  zone.Participant{},
			wantAllParticipants: seeds,
			wantTotalCount:      2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Seed(seeds)
			defer repo.RemoveAll()

			err := repo.Update(tc.newParticipant)
			if err != nil {
				t.Fatalf("Fail to update participant, got error %s", err.Error())
			}

			got := repo.Get(tc.newParticipant.PersonalInfo.NIK)
			assert.Equal(t, tc.wantGotParticipant, got)

			all := repo.All()
			assert.Equal(t, tc.wantAllParticipants, all)

			count := repo.CountAll()
			assert.Equal(t, tc.wantTotalCount, count)
		})
	}
}
