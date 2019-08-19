package zone_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	zone "github.com/aimzeter/wuts/3_usecase"
	"github.com/aimzeter/wuts/3_usecase/testdouble/mocks"
)

type CStoreMock struct {
	GetRet zone.Citizen
}

func (s *CStoreMock) Get(string) zone.Citizen {
	return s.GetRet
}

type PStoreMock struct {
	CreateErr error
	GetRet    zone.Participant
	UpdateErr error

	CreatedParticipant zone.Participant
}

func (s *PStoreMock) Create(c zone.Participant) error {
	s.CreatedParticipant = c
	return s.CreateErr
}
func (s *PStoreMock) Get(string) zone.Participant     { return s.GetRet }
func (s *PStoreMock) Update(c zone.Participant) error { return s.UpdateErr }

func TestRegisterUsingStub(t *testing.T) {
	tests := []struct {
		name          string
		nik           string
		autoReg       bool
		cStoreGetResp zone.Citizen

		wantError              bool
		wantCreatedParticipant zone.Participant
	}{
		{
			name:    "register participant with in range coord",
			nik:     "1234567890",
			autoReg: true,
			cStoreGetResp: zone.Citizen{
				NIK:     "1234567890",
				Name:    "Foo",
				Address: "Bar",
				Coord:   zone.Coord{0, 0},
			},
			wantError: false,
			wantCreatedParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "1234567890",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
				AutoReg: true,
			},
		},
		{
			name:    "register participant with out range coord",
			nik:     "1234567890",
			autoReg: true,
			cStoreGetResp: zone.Citizen{
				NIK:     "1234567890",
				Name:    "Foo",
				Address: "Bar",
				Coord:   zone.Coord{100, 100},
			},
			wantError:              true,
			wantCreatedParticipant: zone.Participant{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cStore := &CStoreMock{}
			pStore := &PStoreMock{}

			// stub cStore.Get
			cStore.GetRet = tc.cStoreGetResp

			app := zone.AppUsecase{CStore: cStore, PStore: pStore}
			err := app.Register(tc.nik, tc.autoReg)

			assert.Equal(t, tc.wantError, err != nil)
			assert.Equal(t, tc.wantCreatedParticipant, pStore.CreatedParticipant)
		})
	}
}

func TestRegisterUsingMock(t *testing.T) {
	tests := []struct {
		name          string
		nik           string
		autoReg       bool
		cStoreGetResp zone.Citizen

		wantError              bool
		wantCreatedParticipant zone.Participant
	}{
		{
			name:    "register participant with in range coord",
			nik:     "1234567890",
			autoReg: true,
			cStoreGetResp: zone.Citizen{
				NIK:     "1234567890",
				Name:    "Foo",
				Address: "Bar",
				Coord:   zone.Coord{0, 0},
			},
			wantError: false,
			wantCreatedParticipant: zone.Participant{
				PersonalInfo: zone.Citizen{
					NIK:     "12345678901",
					Name:    "Foo",
					Address: "Bar",
					Coord:   zone.Coord{0, 0},
				},
				AutoReg: true,
			},
		},
		{
			name:    "register participant with out range coord",
			nik:     "1234567890",
			autoReg: true,
			cStoreGetResp: zone.Citizen{
				NIK:     "1234567890",
				Name:    "Foo",
				Address: "Bar",
				Coord:   zone.Coord{100, 100},
			},
			wantError:              true,
			wantCreatedParticipant: zone.Participant{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			// defer c.Finish()

			cStore := mocks.NewMockCitizenStore(c)
			pStore := mocks.NewMockParticipantStore(c)

			cStore.EXPECT().Get(tc.nik).Return(tc.cStoreGetResp)
			pStore.EXPECT().Create(tc.wantCreatedParticipant)

			app := zone.AppUsecase{CStore: cStore, PStore: pStore}
			err := app.Register(tc.nik, tc.autoReg)

			assert.Equal(t, tc.wantError, err != nil)
		})
	}
}

// c.Finish() hanya verify invoked, jadi meskipun di comment kalau expect nya salah, tetep bakal error
// jadi untuk solusi kasus yg ada kemungkinan suatu fungsi tak dipanggil, bisa diremove c.Finish nya (kalo gamau validasi semua fungsi masing2 dipanggil atau gak)

// for _, tc := range tests {
// 	t.Run(tc.name, func(t *testing.T) {
// 		c := gomock.NewController(t)
// 		defer c.Finish()

// 		cStore := mocks.NewMockCitizenStore(c)
// 		pStore := mocks.NewMockParticipantStore(c)

// 		cStore.EXPECT().Get("234").Return(tc.cStoreGetResp)

// 		createdParticipant := zone.Participant{}
// 		pStore.EXPECT().Create(gomock.Any()).AnyTimes().Do(func(p zone.Participant) {
// 			createdParticipant = p
// 		})

// 		app := zone.AppUsecase{CStore: cStore, PStore: pStore}
// 		err := app.Register(tc.nik, tc.autoReg)

// 		assert.Equal(t, tc.wantError, err != nil)
// 		assert.Equal(t, tc.wantCreatedParticipant, createdParticipant)
// 	})
// }
