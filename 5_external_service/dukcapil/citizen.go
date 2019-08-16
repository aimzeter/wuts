package dukcapil

import (
	"fmt"

	"github.com/aimzeter/wuts/5_external_service"
)

const (
	CitizenPath = "/citizens/"
)

type CitizenResponse struct {
	Identity IdentityResponse `json:"Identitas"`
	Address  AddressResponse  `json:"Alamat"`
}

type IdentityResponse struct {
	NIK  string `json:"NomerKependudukan"`
	Name string `json:"Nama"`
}

type AddressResponse struct {
	Address  string  `json:"Alamat"`
	Area     string  `json:"Kecamatan"`
	City     string  `json:"Kota"`
	PostCode string  `json:"KodePos"`
	Lat      float64 `json:"Lat"`
	Long     float64 `json:"Long"`
}

func (d *Dukcapil) Get(nik string) zone.Citizen {
	res := CitizenResponse{}
	path := CitizenPath + nik

	err := d.get(path, &res)
	if err != nil {
		return zone.Citizen{}
	}

	return res.transform()
}

const addressFmt = "%s Kec. %s, %s, %s"

func (c CitizenResponse) transform() zone.Citizen {
	address := fmt.Sprintf(addressFmt, c.Address.Address, c.Address.Area, c.Address.City, c.Address.PostCode)
	return zone.Citizen{
		NIK:     c.Identity.NIK,
		Name:    c.Identity.Name,
		Address: address,
		Coord: zone.Coord{
			Lat:  c.Address.Lat,
			Long: c.Address.Long,
		},
	}
}
