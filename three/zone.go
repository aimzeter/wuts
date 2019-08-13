package zone

type AppUsecase struct {
	CStore CitizenStore
	PStore ParticipantStore
	SStore StudentStore
}

//go:generate mockgen -destination=testdouble/mocks/citizen_store.go -package=mocks github.com/aimzeter/wuts/three CitizenStore

type CitizenStore interface {
	Get(string) Citizen
}

//go:generate mockgen -destination=testdouble/mocks/participant_store.go -package=mocks github.com/aimzeter/wuts/three ParticipantStore

type ParticipantStore interface {
	Get(string) Participant
	Create(Participant) error
	Update(Participant) error
}

//go:generate mockgen -destination=testdouble/mocks/student_store.go -package=mocks github.com/aimzeter/wuts/three StudentStore

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
