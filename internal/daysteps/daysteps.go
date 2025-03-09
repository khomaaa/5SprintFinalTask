package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

var (
	ErrInvalidData     = errors.New("Не корректный формат данных")
	ErrInvalidDataTime = errors.New("Не корректный формат времени")
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",") // Разделяем входящие данные, исходя по формату: "678,0h50m"
	if len(dataParts) != 2 {
		return ErrInvalidData
	}

	steps, err := strconv.Atoi(strings.TrimSpace(dataParts[0])) // Преобразуем первый элемент в целое число (количество шагов)
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(strings.TrimSpace(dataParts[1])) // Преобразуем третий элемент в time (продолжительность тренировки)
	if err != nil {
		return ErrInvalidDataTime
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("Шаг должен быть больше 0")
	}

	distance := spentenergy.Distance(ds.Steps)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %0.2f\n км.\nВы сожгли %0.2f ккал.", ds.Steps, distance, calories)

	return result, nil
}
