package mysql

import (
	zone "github.com/aimzeter/wuts/4_repository"
	"github.com/jmoiron/sqlx"
)

type Participant struct {
	ID      uint   `db:"id"`
	NIK     string `db:"nik"`
	Name    string `db:"name"`
	Address string `db:"address"`

	Lat  float64 `db:"latitude"`
	Long float64 `db:"longitude"`

	AutoReg    bool    `db:"auto_reg"`
	Distance   float64 `db:"distance"`
	TotalScore float32 `db:"total_score"`
}

// interface validator
var ps zone.ParticipantStore = &ParticipantRepo{}

type ParticipantRepo struct {
	db *sqlx.DB
}

func NewParticipantRepo(db *sqlx.DB) *ParticipantRepo {
	return &ParticipantRepo{db}
}

func (r *ParticipantRepo) Get(nik string) zone.Participant {
	var p Participant
	r.db.Get(&p, "SELECT * FROM participants WHERE nik=?", nik)

	return p.transform()
}

func (r *ParticipantRepo) Create(p zone.Participant) error {
	query := `INSERT INTO participants (nik, name, address, latitude, longitude, auto_reg, distance, total_score)
								VALUES (:nik, :name, :address, :latitude, :longitude, :auto_reg, :distance, :total_score)`

	values := map[string]interface{}{
		"nik":         p.PersonalInfo.NIK,
		"name":        p.PersonalInfo.Name,
		"address":     p.PersonalInfo.Address,
		"latitude":    p.PersonalInfo.Coord.Lat,
		"longitude":   p.PersonalInfo.Coord.Long,
		"auto_reg":    p.AutoReg,
		"distance":    p.Distance,
		"total_score": p.TotalScore,
	}

	_, err := r.db.NamedExec(query, values)
	return err
}

func (r *ParticipantRepo) Update(p zone.Participant) error {
	query := `UPDATE participants SET
				nik=:nik, name=:name, address=:address, latitude=:latitude, longitude=:longitude,
				auto_reg=:auto_reg, distance=:distance, total_score=:total_score WHERE nik=:nik`

	values := map[string]interface{}{
		"nik":         p.PersonalInfo.NIK,
		"name":        p.PersonalInfo.Name,
		"address":     p.PersonalInfo.Address,
		"latitude":    p.PersonalInfo.Coord.Lat,
		"longitude":   p.PersonalInfo.Coord.Long,
		"auto_reg":    p.AutoReg,
		"distance":    p.Distance,
		"total_score": p.TotalScore,
	}

	_, err := r.db.NamedExec(query, values)
	return err
}

func (r *ParticipantRepo) RemoveAll() error {
	_, err := r.db.Exec("TRUNCATE participants")
	return err
}

func (r *ParticipantRepo) All() []zone.Participant {
	var ps []Participant
	r.db.Select(&ps, "SELECT * FROM participants")

	var res []zone.Participant
	for _, p := range ps {
		res = append(res, p.transform())
	}

	return res
}

func (r *ParticipantRepo) CountAll() int {
	var c int
	r.db.Get(&c, "SELECT COUNT(1) FROM participants")
	return c
}

func (r *ParticipantRepo) Seed(seeds []zone.Participant) {
	for _, p := range seeds {
		r.Create(p)
	}
}

func (p Participant) transform() zone.Participant {
	return zone.Participant{
		PersonalInfo: zone.Citizen{
			NIK:     p.NIK,
			Name:    p.Name,
			Address: p.Address,
			Coord:   zone.Coord{Lat: p.Lat, Long: p.Long},
		},
		AutoReg:    p.AutoReg,
		Distance:   p.Distance,
		TotalScore: p.TotalScore,
	}
}
