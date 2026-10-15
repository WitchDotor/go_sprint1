package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/fitneserrors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит информацию о шагах за день.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "шаги,длительность".
func (ds *DaySteps) Parse(dataset string) error {
	ds.Steps = 0
	ds.Duration = 0

	data := strings.Split(dataset, ",")
	if len(data) != 2 || strings.Contains(dataset, " ") {
		return fitneserrors.ErrInvalidDataFormat
	}

	dataSteps, err := strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return fitneserrors.ErrInvalidDataFormat
	}

	dataDuration, err := time.ParseDuration(strings.TrimSpace(data[1]))
	if err != nil {
		return fitneserrors.ErrInvalidDataFormat
	}

	if dataSteps <= 0 || dataDuration <= 0 {
		return fitneserrors.ErrInvalidDataFormat
	}

	ds.Steps = dataSteps
	ds.Duration = dataDuration
	return nil
}

// ActionInfo возвращает информацию о расстоянии и потраченных калориях.
func (ds *DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	energy, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, energy,
	), nil
}

// Print выводит персональные данные.
func (ds *DaySteps) Print() {
	ds.Personal.Print()
}
