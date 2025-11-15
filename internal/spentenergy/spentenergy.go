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
	midSpeed := MeanSpeed(steps, height,duration)
	durationMiin := duration.Minutes()
	spentCalories := (weight * midSpeed * durationMiin) / minInH
	spentCaloriesWalk := spentCalories * walkingCaloriesCoefficient
	return spentCaloriesWalk, nil
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
	avgSpeed := MeanSpeed(steps,height,duration)
	durationMin := duration.Minutes()
	spentCaloriesRun := (weight * avgSpeed * durationMin) / minInH
	return spentCaloriesRun, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	newDistance := Distance(steps, height)
	avgSpeed := newDistance / duration.Hours()
	return avgSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	strideLength := height * stepLengthCoefficient // длина шага
	dis := (float64(steps) * strideLength) / mInKm
	return dis
}
