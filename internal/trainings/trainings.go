package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

var (
	ErrInvalidData     = errors.New("Не корректный формат данных")
	ErrInvalidActivity = errors.New("Не корректный тип тренировки")
	ErrInvalidDataTime = errors.New("Не корректный формат времени")
)

// создайте метод Parse()
// Принимает данные формата - "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) (err error) {

	dataParts := strings.Split(datastring, ",") // Разделяем входящие данные, исходя по формату: "3456,Ходьба,3h00m"
	if len(dataParts) != 3 {
		return ErrInvalidData
	}

	steps, err := strconv.Atoi(strings.TrimSpace(dataParts[0])) // Преобразуем первый элемент в целое число (количество шагов)
	if err != nil {
		return err
	}
	t.Steps = steps

	// 2 элемент - строка (активность)
	if strings.ToLower(dataParts[1]) != "ходьба" && strings.ToLower(dataParts[1]) != "бег" {
		return ErrInvalidActivity
	}
	t.TrainingType = dataParts[1]

	duration, err := time.ParseDuration(strings.TrimSpace(dataParts[2])) // Преобразуем третий элемент в time (продолжительность тренировки)
	if err != nil {
		return ErrInvalidDataTime
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	cal := 0.0
	if strings.ToLower(t.TrainingType) == "ходьба" {
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		if err != nil {
			return "", err
		}
		cal = calories
	} else if strings.ToLower(t.TrainingType) == "бег" {
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		cal = calories
	} else {
		return "", ErrInvalidActivity
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f", t.TrainingType, t.Duration.Hours(), distance, speed, cal)
	return result, nil
}
