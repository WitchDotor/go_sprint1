package spentenergy

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/fitneserrors"
)

// Коэффициенты для расчётов.
const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Границы допустимых значений.
const (
	minWeight   = 20.0
	maxWeight   = 200.0
	minHeight   = 0.5
	maxHeight   = 2.5
	maxDuration = 24 * time.Hour
)

// Distance возвращает пройденное расстояние в километрах.
func Distance(steps int, height float64) float64 {
	return float64(steps) * stepLengthCoefficient * height / mInKm
}

// MeanSpeed возвращает среднюю скорость в км/ч.
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
		return fmt.Errorf("%w: получено %.2f, ожидаемый диапазон [%.1f, %.1f]",
			fitneserrors.ErrInvalidHeight, height, minHeight, maxHeight)
	case duration <= 0 || duration > maxDuration:
		return fmt.Errorf("%w: получено %s, ожидаемый диапазон (0, %s]",
			fitneserrors.ErrInvalidDuration, duration, maxDuration)
	}
	return nil
}

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkValidData(steps, weight, height, duration); err != nil {
		return 0, err
	}
	return walkingCaloriesCoefficient * weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkValidData(steps, weight, height, duration); err != nil {
		return 0, err
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}
