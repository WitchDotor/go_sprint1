package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/fitneserrors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(dataset string) (err error) {
	data := strings.Split(dataset, ",")
	if len(data) != 3 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = 0
		return fmt.Errorf("%w", fitneserrors.ErrTrainingType)
	}

	for i := range data {
		if data[i] == "" {
			t.Steps = 0
			t.TrainingType = ""
			t.Duration = 0
			return fmt.Errorf("%w", fitneserrors.ErrTrainingType)

		}
		if i == 1 {
			if data[i] != "Бег" && data[i] != "Ходьба" {
				t.Steps = 0
				t.TrainingType = ""
				t.Duration = 0
				return fmt.Errorf("%w", fitneserrors.ErrTrainingType)
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
	case "Ходьба":
		energy, energyError = spentenergy.WalkingSpentCalories(t.Steps, weight, height, t.Duration)
	case "Бег":
		energy, energyError = spentenergy.RunningSpentCalories(t.Steps, weight, height, t.Duration)
	default:
		return "", fmt.Errorf("%w", fitneserrors.ErrTrainingType)
	}

	if energyError != nil {
		return "", energyError
	}
	var distance = float64(t.Duration) * speed
	var strType = fmt.Sprintf("Тип тренировки: %s.\n", t.TrainingType)
	var strDur = fmt.Sprintf("Длительность: %ds.\n", t.Duration)
	var strDist = fmt.Sprintf("Дистанция: %.2f.\n", distance)
	var strSpeed = fmt.Sprintf("Скорость: %.2f.\n", speed)
	var strEnergy = fmt.Sprintf("Сожжено калорий: %.0f.\n", energy)
	var strTraining = strType + strDur + strDist + strSpeed + strEnergy
	return fmt.Sprintf("%s", strTraining), nil
}

func (t *Training) Print() (string, error) {
	var strSteps = fmt.Sprintf("%i\n", t.Steps)
	var strTrainingType = fmt.Sprintf("%s\n", t.TrainingType)
	var strDuration = fmt.Sprintf("%d\n", t.Duration)
	var strTraining = strSteps + strTrainingType + strDuration
	return fmt.Sprintf("%s", strTraining), nil
}
