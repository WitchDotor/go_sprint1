package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/fitneserrors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	var parseErr error
	if len(data) != 2 {
		ds.Steps = 0
		ds.Duration = 0
		parseErr = fitneserrors.ErrInvalidFormat
		return fmt.Errorf("%w", parseErr)
	}

	for i := range data {
		if data[i] == "" {
			ds.Steps = 0
			ds.Duration = 0
			parseErr = fitneserrors.ErrInvalidFormat
			return fmt.Errorf("%w", parseErr)
		}
	}
	ds.Duration, parseErr = time.ParseDuration(data[1])
	ds.Steps, parseErr = strconv.Atoi(data[0])

	if parseErr != nil {
		ds.Duration = 0
		ds.Steps = 0
	}
	return parseErr
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var person = ds.Personal.Height
	var distance = ds.Steps * ds.Personal.Height * stepLengthCoefficient
	fmt.Printf("Количество шагов: %d.\n", ds.Steps)
	fmt.Printf("Дистанция составила %sкм.\n", distance)
	fmt.Printf("Вы сожгли %f ккал.\n", energy)
}
