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
	minHeight                  = 0.5
	maxHeight                  = 2.5
	maxDuration                = 24 * time.Hour
)

func Distance(steps int, height float64) float64 {
	return float64(steps) * stepLengthCoefficient * height / 1000
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func checkValidData(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return fitneserrors.ErrStepsValue
	case weight < minWeight || weight > maxWeight:
		return fitneserrors.ErrInvalidWeight
	case height < minHeight || height > maxHeight:
		return fmt.Errorf("%w: дано %.2f, expected range [%.1f, %.1f]",
			fitneserrors.ErrInvalidHeight, height, minHeight, maxHeight)
	case duration <= 0 || duration > maxDuration:
		return fmt.Errorf("%w: дано %s, expected range (0, %s]", fitneserrors.ErrInvalidDuration, duration, maxDuration)
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var dataError = checkValidData(steps, weight, height, duration)
	var result float64
	if dataError == nil {
		result = walkingCaloriesCoefficient * (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / 60
	}
	return result, dataError
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var dataError = checkValidData(steps, weight, height, duration)
	var result float64
	if dataError == nil {
		result = weight * MeanSpeed(steps, height, duration) * duration.Minutes() / 60
	}

	return result, dataError
}
