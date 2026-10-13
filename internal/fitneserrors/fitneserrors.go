package fitneserrors

import (
	"fmt"
)

func New(text string) error {
	return &validationError{text}
}

type validationError struct {
	s string
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%s", e.s)
}

var (
	ErrInvalidWeight     = New("Неправильно указан вес")
	ErrInvalidHeight     = New("Неправильно указан рост")
	ErrInvalidDuration   = New("Некорректное значение продолжительности")
	ErrTrainingType      = New("Неверный формат тренировки")
	ErrStepsValue        = New("Неправильное значение шагов")
	ErrInvalidDataFormat = New("Полученный формат данных невозможно обработать")
)
