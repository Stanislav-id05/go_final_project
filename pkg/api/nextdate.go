package api

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Stanislav-id05/go_final_project/pkg/constants"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	// Проверка на пустую строку в repeat
	if repeat == "" {
		return "", errors.New("пустая строка в колонке repeat")
	}
	// Парсинг исходной даты
	startDate, err := time.Parse(constants.DateFormat, date)
	if err != nil {
		return "", errors.New("некорректный формат даты")
	}

	repeatParts := strings.Fields(repeat)

	var nextDate time.Time

	switch repeatParts[0] {
	case "d":
		if len(repeatParts) != 2 {
			return "", errors.New("не указан интервал в днях")
		}
		days, err := strconv.Atoi(repeatParts[1])
		if err != nil || days < 1 || days > 400 {
			return "", errors.New("недопустимое значение дней")
		}
		nextDate = startDate.AddDate(0, 0, days)
		// Цикл для нахождения следующей даты
		for nextDate.Before(now) || nextDate.Equal(now) {
			nextDate = nextDate.AddDate(0, 0, days)
		}
		return nextDate.Format(constants.DateFormat), nil

	case "y":
		nextDate = startDate.AddDate(1, 0, 0)
		// Цикл для нахождения следующей даты
		for nextDate.Before(now) || nextDate.Equal(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
		}
		return nextDate.Format(constants.DateFormat), nil

	default:
		return "", errors.New("неподдерживаемый формат")
	}
}
