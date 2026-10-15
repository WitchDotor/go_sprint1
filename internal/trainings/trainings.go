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
	var stepsErr error
	var durationErr error
	if len(data) != 3 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = 0
		return fmt.Errorf("%w", fitneserrors.ErrTrainingType)
	}
	steps, stepsErr := strconv.Atoi(data[0])
	trainingType := data[1]
	trainingDuration, durationErr := time.ParseDuration(data[2])

	if steps <= 0 || trainingDuration <= 0 || stepsErr != nil || durationErr != nil {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = 0
		return fmt.Errorf("%w", fitneserrors.ErrTrainingType)
	}

	if trainingType != "Бег" && trainingType != "Ходьба" {
		trainingType = "Неизвестный тип тренировки"
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = trainingDuration
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
	var distance = spentenergy.Distance(t.Steps, height)
	var strType = fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	var strDur = fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	var strDist = fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	var strSpeed = fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	var strEnergy = fmt.Sprintf("Сожгли калорий: %.2f\n", energy)
	var strTraining = strType + strDur + strDist + strSpeed + strEnergy
	return fmt.Sprintf("%s", strTraining), nil
}

func (t *Training) Print() (string, error) {
	var strSteps = fmt.Sprintf("%d\n", t.Steps)
	var strTrainingType = fmt.Sprintf("%s\n", t.TrainingType)
	var strDuration = fmt.Sprintf("%d\n", t.Duration)
	var strTraining = strSteps + strTrainingType + strDuration
	return fmt.Sprintf("%s", strTraining), nil
}
