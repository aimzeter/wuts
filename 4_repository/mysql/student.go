package mysql

import (
	zone "github.com/aimzeter/wuts/4_repository"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	ID      uint   `db:"id"`
	NIK     string `db:"nik"`
	Name    string `db:"name"`
	Address string `db:"address"`

	Lat  float64 `db:"latitude"`
	Long float64 `db:"longitude"`
}

// interface validator
var ss zone.StudentStore = &StudentRepo{}

type StudentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) *StudentRepo {
	return &StudentRepo{db}
}

func (r *StudentRepo) Get(nik string) zone.Student {
	var s Student
	r.db.Get(&s, "SELECT * FROM students WHERE nik=?", nik)

	return s.transform()
}

func (r *StudentRepo) Create(s zone.Student) error {
	query := `INSERT INTO students (nik, name, address, latitude, longitude) VALUES (:nik, :name, :address, :latitude, :longitude)`

	values := map[string]interface{}{
		"nik":       s.PersonalInfo.NIK,
		"name":      s.PersonalInfo.Name,
		"address":   s.PersonalInfo.Address,
		"latitude":  s.PersonalInfo.Coord.Lat,
		"longitude": s.PersonalInfo.Coord.Long,
	}

	_, err := r.db.NamedExec(query, values)
	return err
}

func (r *StudentRepo) RemoveAll() error {
	_, err := r.db.Exec("TRUNCATE students")
	return err
}

func (r *StudentRepo) All() []zone.Student {
	var ss []Student
	r.db.Select(&ss, "SELECT * FROM students")

	var res []zone.Student
	for _, s := range ss {
		res = append(res, s.transform())
	}

	return res
}

func (r *StudentRepo) CountAll() int {
	var c int
	r.db.Get(&c, "SELECT COUNT(1) FROM students")
	return c
}

func (r *StudentRepo) Seed(seeds []zone.Student) {
	for _, s := range seeds {
		r.Create(s)
	}
}

func (s Student) transform() zone.Student {
	return zone.Student{
		PersonalInfo: zone.Citizen{
			NIK:     s.NIK,
			Name:    s.Name,
			Address: s.Address,
			Coord:   zone.Coord{Lat: s.Lat, Long: s.Long},
		},
	}
}
