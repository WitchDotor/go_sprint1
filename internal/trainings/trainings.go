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

// Training описывает тренировку: количество шагов, тип и длительность.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "шаги,тип,длительность".
func (t *Training) Parse(dataset string) error {
	t.Steps = 0
	t.TrainingType = ""
	t.Duration = 0

	data := strings.Split(dataset, ",")
	if len(data) != 3 {
		return fitneserrors.ErrInvalidDataFormat
	}

	steps, err := strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return fitneserrors.ErrInvalidDataFormat
	}

	duration, err := time.ParseDuration(strings.TrimSpace(data[2]))
	if err != nil {
		return fitneserrors.ErrInvalidDataFormat
	}

	if steps <= 0 || duration <= 0 {
		return fitneserrors.ErrInvalidDataFormat
	}

	t.Steps = steps
	t.TrainingType = strings.TrimSpace(data[1])
	t.Duration = duration
	return nil
}

// ActionInfo возвращает текстовую сводку по тренировке.
func (t *Training) ActionInfo() (string, error) {
	var energy float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		energy, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		energy, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fitneserrors.ErrTrainingType
	}
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, energy,
	), nil
}

// Print выводит краткую информацию о тренировке.
func (t *Training) Print() {
	fmt.Printf("%d\n%s\n%s\n", t.Steps, t.TrainingType, t.Duration)
}
