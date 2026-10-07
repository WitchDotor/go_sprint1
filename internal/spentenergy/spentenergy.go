package spentenergy

import (
	"time"
	"fmt"
	"strconv"
	"strings"	
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
	minWeight = 20.0
	maxWeight = 200.0
	minHeight = 50.0
	maxHeight = 280.0
	maxSteps  = 200000
	maxDuration = 24 * time.Hour
)

func Distance(steps int, height float64) float64 {
	return steps * stepLengthCoefficient * height
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration == 0 {
		return 0
	} else { 
		return Distance / (duration / time.Hour)
	}
}

func checkValidData(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return fmt.Errorf("%w: got %d", ErrInvalidSteps, steps)
	case steps > maxSteps:
		return fmt.Errorf("steps too large: got %d, max %d", steps, maxSteps)
	case weight < minWeight || weight > maxWeight:
		return fmt.Errorf("%w: got %.2f, expected range [%.1f, %.1f]",
			ErrInvalidWeight, weight, minWeight, maxWeight)
	case height < minHeight || height > maxHeight:
		return fmt.Errorf("%w: got %.2f, expected range [%.1f, %.1f]",
			ErrInvalidHeight, height, minHeight, maxHeight)
	case duration <= 0 || duration > maxDuration:
		return fmt.Errorf("%w: got %s, expected range (0, %s]",
			ErrInvalidDuration, duration, maxDuration)
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var data_error := checkValidData(steps, weight, height, duration)
	val result int
	if data_error == nil{
		result = (walkingCaloriesCoefficient * 
		(weight * MeanSpeed(steps, height, duration) * duration/time.Minute) / duration/time.Hour)
	}
	
	return result, data_error
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var data_error := checkValidData(steps, weight, height, duration)
	val result int
	if data_error == nil{
		result =  ((weight * MeanSpeed(steps, height, duration) * duration/time.Minute) 
		/ duration/time.Hour)
	}
	
	return result, data_error
}



