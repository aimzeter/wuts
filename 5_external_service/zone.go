package zone

type AppUsecase struct {
	CStore CitizenStore
	PStore ParticipantStore
	SStore StudentStore
}

type CitizenStore interface {
	Get(string) Citizen
}

type ParticipantStore interface {
	Get(string) Participant
	Create(Participant) error
	Update(Participant) error
}

type StudentStore interface {
	Create(Student) error
}

type Citizen struct {
	NIK     string
	Name    string
	Address string
	Coord   Coord
}

type Participant struct {
	PersonalInfo Citizen
	AutoReg      bool
	Distance     float64
	TotalScore   float32
}

type Student struct {
	PersonalInfo Citizen
}

type Coord struct {
	Lat, Long float64
}
