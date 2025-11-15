package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errors.New("incorrect slice length")
	}
	step, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	if step <= 0 {
		return errors.New("the number of steps cannot be zero or negative")
	}
	t.Steps = step
	t.TrainingType = slice[1]
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	} 
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps,t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		runCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, runCalories), nil

	case "Ходьба":
		walkCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, walkCalories), nil
	
	default:
		return "", errors.New("неизвестный тип тренировки")
		} 
}
