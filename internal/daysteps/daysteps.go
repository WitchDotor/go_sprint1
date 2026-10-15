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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(dataset string) (err error) {
	data := strings.Split(dataset, ",")
	var parseErr error
	var stepsErr error
	var durationErr error
	if len(data) != 2 {
		ds.Steps = 0
		ds.Duration = 0
		parseErr = fitneserrors.ErrInvalidDataFormat
		return fmt.Errorf("%w", parseErr)
	}

	dataSteps, stepsErr := strconv.Atoi(data[0])
	dataDuration, durationErr := time.ParseDuration(data[1])
	if dataSteps <= 0 || dataDuration <= 0 || stepsErr != nil || durationErr != nil {
		ds.Duration = 0
		ds.Steps = 0
		parseErr = fitneserrors.ErrInvalidDataFormat
		return fmt.Errorf("%w", parseErr)
	}

	ds.Duration = dataDuration
	ds.Steps = dataSteps

	return parseErr
}

func (ds *DaySteps) ActionInfo() (string, error) {

	var height = ds.Personal.Height
	var weight = ds.Personal.Weight
	var distance = spentenergy.Distance(ds.Steps, height)
	var energy, actionError = spentenergy.WalkingSpentCalories(ds.Steps, weight, height, ds.Duration)
	if actionError != nil {
		return "", actionError
	}

	var strSteps = fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	var strDist = fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	var strEnergy = fmt.Sprintf("Вы сожгли %.2f ккал.\n", energy)
	var strDayStep = strSteps + strDist + strEnergy
	return strDayStep, nil
}

func (ds *DaySteps) Print() {
	ds.Personal.Print()
}
