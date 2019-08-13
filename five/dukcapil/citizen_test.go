package dukcapil_test

import (
	"testing"

	zone "github.com/aimzeter/wuts/five"
	"github.com/aimzeter/wuts/five/dukcapil"

	"github.com/stretchr/testify/assert"
)

func TestGetCitizen(t *testing.T) {
	tests := []struct {
		name      string
		nik       string
		dkcplResp string
		dkcplCode int

		wantResult zone.Citizen
	}{
		{
			name: "success",
			nik:  "1234567890",
			dkcplResp: `
				{
					"Identitas": {
						"NomerKependudukan": "1234567890",
						"Nama": "Foo Bar"
					},
					"Alamat": {
						"Alamat": "Jalan Diponegoro II No. 31 Kel. Bagan",
						"Kecamatan": "Bogor Barat",
						"Kota": "Bogor",
						"KodePos": "16610",
						"Lat": 123.45,
						"Long": -87.12
					}
				}
			`,
			dkcplCode: 200,
			wantResult: zone.Citizen{
				NIK:     "1234567890",
				Name:    "Foo Bar",
				Address: "Jalan Diponegoro II No. 31 Kel. Bagan Kec. Bogor Barat, Bogor, 16610",
				Coord: zone.Coord{
					Lat: 123.45, Long: -87.12,
				},
			},
		},
		{
			name: "invalid response body",
			nik:  "1234567890",
			dkcplResp: `
				{
					invalid_field: "1"
				}
			`,
			dkcplCode:  200,
			wantResult: zone.Citizen{},
		},
		{
			name:       "non ok status code",
			nik:        "1234567890",
			dkcplCode:  500,
			wantResult: zone.Citizen{},
		},
	}

	h := DukcapilHandler{Endpoint: dukcapil.CitizenPath}
	s := SetupDukcapilServer(&h)
	defer s.Close()

	c := dukcapil.NewClient()
	d := dukcapil.NewDukcapil(c, s.URL)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h.Stub(tc.dkcplResp, tc.dkcplCode)

			res := d.Get(tc.nik)
			assert.Equal(t, tc.wantResult, res)
		})
	}
}
