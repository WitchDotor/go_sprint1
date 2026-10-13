package daysteps

import (
	"fmt"
	"slices"
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
	if len(data) != 2 {
		ds.Steps = 0
		ds.Duration = 0
		parseErr = fitneserrors.ErrInvalidDataFormat
		return fmt.Errorf("%w", parseErr)
	}

	if slices.Contains(data, "") {
		ds.Steps = 0
		ds.Duration = 0
		parseErr = fitneserrors.ErrInvalidDataFormat
		return fmt.Errorf("%w", parseErr)
	}
	ds.Duration, parseErr = time.ParseDuration(data[1])
	ds.Steps, parseErr = strconv.Atoi(data[0])

	if parseErr != nil {
		ds.Duration = 0
		ds.Steps = 0
	}
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
	var strDist = fmt.Sprintf("Дистанция составила %.2fкм.\n", distance)
	var strEnergy = fmt.Sprintf("Вы сожгли %.2f ккал.\n", energy)
	var strDayStep = strSteps + strDist + strEnergy
	return strDayStep, nil
}

func (ds *DaySteps) Print() {
	ds.Personal.Print()
}
