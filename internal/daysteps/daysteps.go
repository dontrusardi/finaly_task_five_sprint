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
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("incorrect slice length")
	}
	countSteps, err := strconv.Atoi(slice[0])
	if err != nil{
		return err
	}
	if countSteps <= 0 {
		return errors.New("the number of steps cannot be zero or negative")
	}
	ds.Steps = countSteps
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	} 
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dis := spentenergy.Distance(ds.Steps, ds.Height)
	spentWalk, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dis, spentWalk), nil
}
