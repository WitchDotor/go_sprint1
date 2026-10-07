package trainings

import (
	"time"
	"fmt"
	"strconv"
	"strings"
	"spentenergy"	
)

type Training struct {
	Steps int,
	TrainingType string,
	Duration time.Duration,
	personaldata.Personal ,
}

func (t *Training) Parse(datastring string) (err error) {
	data := datastring.Split(s, ",")
	if len(data) != 3 {
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
	
	&Steps = strconv.Atoi(data[0])
	&TrainingType = data[1]
	&Duration = time.ParseDuration(data[2])
	return nill
}

func (t Training) ActionInfo() (string, error) {
	height = personaldata.Height
	weight = personaldata.Weight
	speed = spentenergy.MeanSpeed(steps=Steps, height=personaldata.Height, duration=Duration)

	switch TrainingType{
		case "ходьба":
			energy = spentenergy.WalkingSpentCalories(steps=Steps, weight=weight, height=height, duration=Duration)
		case "бег":
			energy = spentenergy.RunningSpentCalories(steps=Steps, weight=weight, height=height, duration=Duration)
		default:
			return "", fmt.Errorf("Неправильный вормат данных, нет поля", TrainingType)
		}

	
	// TODO: реализовать функцию
	str_type="Тип тренировки: %s", TrainingType
	str_dur="Длительность: %sd", Duration
	str_dist="Дистанция: %sf", Distance
	str_speed="Скорость: %sf", speed
	str_energy"Сожгли калорий: %sf", energy

	return str_type+str_dur+str_dist+str_speed+str_energy, nill
	
}
