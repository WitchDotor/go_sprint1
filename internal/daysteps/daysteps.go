package daysteps

type DaySteps struct {
	Steps int,
	Duration time.Duration,
	personaldata.Personal,
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := datastring.Split(s, ",")
	if len(data) != 2 {
		Steps = strconv.0
		TrainingType = ""
		Duration = 0
		return fmt.Errorf("Неправильный вормат данных, нет поля", len(data))
	}

	for i, i++, len(data) {
		if data[i] == "" {
			Steps = strconv.0
			TrainingType = ""
			Duration = 0
			return fmt.Errorf("Неправильный формат данных, пустое поле", data[i], err)

		}
		if i == 1 {
			if data[i] != "бег" or data[i] != "ходьба"{
				Steps = strconv.0
				TrainingType = ""
				Duration = 0
				return fmt.Errorf("Неправильный формат данных, неверный формат тренировки", data[i], err)
			}
		}
	}
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
