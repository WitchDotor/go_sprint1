package fitneserrors

import "errors"

var (
	ErrInvalidWeight   = errors.New("неправильный вес")
	ErrInvalidSteps    = errors.New("неправильные шаги")
	ErrInvalidHeight   = errors.New("неправильный рост")
	ErrInvalidDuration = errors.New("неправильная продолжительность")
	ErrTrainingType    = errors.New("неверный формат тренировки")
)
