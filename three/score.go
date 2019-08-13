package zone

func (app *AppUsecase) SetScore(nik string, score float32) error {
	p := app.PStore.Get(nik)
	p.TotalScore = score

	err := app.PStore.Update(p)
	if err != nil {
		return err
	}

	if PassScore(p.TotalScore) {
		if p.AutoReg {
			app.RegisterSchool(nik)
		}
	}

	return nil
}
