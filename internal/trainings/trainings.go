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
	Slice := strings.Split(datastring, ",")
	if len(Slice) != 3 {
		return errors.New("incorrect slice length")
	}
	Step, err := strconv.Atoi(Slice[0])
	if err != nil {
		return err
	}
	if Step <= 0 {
		return errors.New("the number of steps cannot be zero or negative")
	}
	t.Steps = Step
	t.TrainingType = Slice[1]
	Duration, err := time.ParseDuration(Slice[2])
	if err != nil {
		return err
	}
	if Duration <= 0 {
		return errors.New("duration must be greater than zero")
	} 
	t.Duration = Duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	Distance := spentenergy.Distance(t.Steps,t.Height)
	AvgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		RunCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), Distance, AvgSpeed, RunCalories), nil

	case "Ходьба":
		WalkCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), Distance, AvgSpeed, WalkCalories), nil
	
	default:
		return "", errors.New("неизвестный тип тренировки")
		} 
}
