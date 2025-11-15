package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be zero or negative")
	}
	if weight <= 0 {
		return 0, errors.New("the weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("growth must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}
	MidSpeed := MeanSpeed(steps, height,duration)
	DurationMiin := duration.Minutes()
	SpentCalories := (weight * MidSpeed * DurationMiin) / minInH
	SpentCaloriesWalk := SpentCalories * walkingCaloriesCoefficient
	return SpentCaloriesWalk, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be zero or negative")
	}
	if weight <= 0 {
		return 0, errors.New("the weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("growth must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}
	AvgSpeed := MeanSpeed(steps,height,duration)
	DurationMin := duration.Minutes()
	SpentCaloriesRun := (weight * AvgSpeed * DurationMin) / minInH
	return SpentCaloriesRun, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	NewDistance := Distance(steps, height)
	AvgSpeed := NewDistance / duration.Hours()
	return AvgSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	StrideLength := height * stepLengthCoefficient // длина шага
	Dis := (float64(steps) * StrideLength) / mInKm
	return Dis
}
