package zone_test

import (
	"testing"

	zone "github.com/aimzeter/wuts/three"
	"github.com/aimzeter/wuts/three/testdouble/fake"

	"github.com/stretchr/testify/assert"
)

func TestSetScore(t *testing.T) {
	// Given
	tests := []struct {
		name            string
		nik             string
		score           float32
		participantSeed []zone.Participant

		wantCreatedStudent zone.Student
	}{
		{
			name:  "participant agree to auto register school when pass the exam",
			nik:   "1234567890",
			score: 90.0,
			participantSeed: []zone.Participant{{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
				AutoReg:  true,
				Distance: 0.0,
			}},
			wantCreatedStudent: zone.Student{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
			},
		},
		{
			name:  "participant disagree to auto register school when pass the exam",
			nik:   "1234567890",
			score: 90.0,
			participantSeed: []zone.Participant{{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
				AutoReg:  false,
				Distance: 0.0,
			}},
			wantCreatedStudent: zone.Student{},
		},
		{
			name:  "participant score not pass threshold",
			nik:   "1234567890",
			score: 60.0,
			participantSeed: []zone.Participant{{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
				AutoReg:  false,
				Distance: 0.0,
			}},
			wantCreatedStudent: zone.Student{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// prepare seeded student store
			pStore := fake.NewFakeParticipantStore()
			pStore.Seed(tc.participantSeed)

			// prepare empty student store
			sStore := fake.NewFakeStudentStore()

			// When
			app := zone.AppUsecase{PStore: pStore, SStore: sStore}
			err := app.SetScore(tc.nik, tc.score)

			// Then
			if err != nil {
				t.Fatalf("Fail to set score, got error %s", err.Error())
			}

			currentParticipant := pStore.Get(tc.nik)
			assert.Equal(t, tc.score, currentParticipant.TotalScore, "SetScore should save score")

			savedStudent := sStore.Get(tc.nik)
			assert.Equal(t, tc.wantCreatedStudent, savedStudent, "SetScore did not create student correctly")
		})
	}
}
