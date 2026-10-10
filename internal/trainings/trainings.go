package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = 0
		return fmt.Errorf("неправильный формат данных, нет поля")
	}

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			t.Steps = 0
			t.TrainingType = ""
			t.Duration = 0
			return fmt.Errorf("неправильный формат данных, пустое поле")

		}
		if i == 1 {
			if data[i] != "бег" && data[i] != "ходьба" {
				t.Steps = 0
				t.TrainingType = ""
				t.Duration = 0
				return fmt.Errorf()
			}
		}
	}

	t.Steps, _ = strconv.Atoi(data[0])
	t.TrainingType = data[1]
	t.Duration, _ = time.ParseDuration(data[2])
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var height = t.Personal.Height
	var weight = t.Personal.Weight
	var speed = spentenergy.MeanSpeed(t.Steps, height, t.Duration)
	var energy float64
	var energyError error

	switch t.TrainingType {
	case "ходьба":
		energy, energyError = spentenergy.WalkingSpentCalories(t.Steps, weight, height, t.Duration)
	case "бег":
		energy, energyError = spentenergy.RunningSpentCalories(t.Steps, weight, height, t.Duration)
	default:
		return "", fmt.Errorf("Неправильный вормат данных, нет поля", t.TrainingType)
	}

	if energyError == nil {

	}
	var distance = float64(t.Duration) * speed
	var strType = fmt.Sprintf("Тип тренировки: %s", t.TrainingType)
	var strDur = fmt.Sprintf("Длительность: %ds", t.Duration)
	var strDist = fmt.Sprintf("Дистанция: %.2f", distance)
	var strSpeed = fmt.Sprintf("Скорость: %.2f", speed)
	var strEnergy = fmt.Sprintf("Сожжено калорий: %.0f", energy)
	return strType + strDur + strDist + strSpeed + strEnergy, nil

}
