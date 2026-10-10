package spentenergy

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/fitneserrors"
	_ "github.com/Yandex-Practicum/tracker/internal/fitneserrors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
	minWeight                  = 20.0
	maxWeight                  = 200.0
	minHeight                  = 50.0
	maxHeight                  = 280.0
	maxSteps                   = 200000
	maxDuration                = 24 * time.Hour
)

func Distance(steps int, height float64) float64 {
	return float64(steps) * stepLengthCoefficient * height
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration == 0 {
		return 0
	}

	return Distance(steps, height) / float64(duration/time.Hour)
}

func checkValidData(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return fmt.Errorf("%w: got %d", fitneserrors.ErrInvalidSteps, steps)
	case steps > maxSteps:
		return fmt.Errorf("steps too large: got %d, max %d", steps, maxSteps)
	case weight < minWeight || weight > maxWeight:
		return fmt.Errorf("%w: got %.2f, expected range [%.1f, %.1f]",
			fitneserrors.ErrInvalidWeight, weight, minWeight, maxWeight)
	case height < minHeight || height > maxHeight:
		return fmt.Errorf("%w: got %.2f, expected range [%.1f, %.1f]",
			fitneserrors.ErrInvalidHeight, height, minHeight, maxHeight)
	case duration <= 0 || duration > maxDuration:
		return fmt.Errorf("%w: got %s, expected range (0, %s]", fitneserrors.ErrInvalidDuration, duration, maxDuration)
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var dataError = checkValidData(steps, weight, height, duration)
	var result float64
	if dataError == nil {
		result = walkingCaloriesCoefficient *
			(weight * MeanSpeed(steps, height, duration) * float64(duration) / float64(time.Minute)) / float64(duration/time.Hour)
	}
	return result, dataError
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var dataError = checkValidData(steps, weight, height, duration)
	var result float64
	if dataError == nil {
		result = weight * MeanSpeed(steps, height, duration) * float64(duration/time.Minute) / float64(duration/time.Hour)
	}

	return result, dataError
}
