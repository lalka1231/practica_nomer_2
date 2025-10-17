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
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInHours := duration.Hours()

	met := 2.5
	calories := met * weight * durationInHours * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInHours := duration.Hours()

	met := 8.0
	calories := met * weight * durationInHours
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	hours := duration.Hours()
	return dist / hours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepDist := height * stepLengthCoefficient
	dist := (float64(steps) * stepDist) / mInKm
	return dist
}
