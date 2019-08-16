package zone

import (
	"errors"
)

func (app *AppUsecase) Register(nik string, autoReg bool) error {
	c := app.CStore.Get(nik)
	dis := SchoolDistance(c.Coord)

	if dis > ZoneThreshold {
		return errors.New("sorry you're too far away")
	}

	p := Participant{}
	p.PersonalInfo = c
	p.AutoReg = autoReg
	p.Distance = dis

	err := app.PStore.Create(p)
	return err
}

func (app *AppUsecase) RegisterSchool(nik string) error {
	p := app.PStore.Get(nik)
	if !PassScore(p.TotalScore) {
		return errors.New("thou shalt not pass")
	}

	s := Student{}
	s.PersonalInfo = p.PersonalInfo

	err := app.SStore.Create(s)
	if err != nil {
		return err
	}

	return nil
}
