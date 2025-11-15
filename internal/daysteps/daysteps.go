package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	Slice := strings.Split(datastring, ",")
	if len(Slice) != 2 {
		return errors.New("Incorrect slice length")
	}
	CountSteps, err := strconv.Atoi(Slice[0])
	if err != nil{
		return err
	}
	if CountSteps <= 0 {
		return errors.New("the number of steps cannot be zero or negative")
	}
	ds.Steps = CountSteps
	Duration, err := time.ParseDuration(Slice[1])
	if err != nil {
		return err
	}
	if Duration <= 0 {
		return errors.New("duration must be greater than zero")
	} 
	ds.Duration = Duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	Dis := spentenergy.Distance(ds.Steps, ds.Height)
	SpentWalk, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, Dis, SpentWalk), nil
}
