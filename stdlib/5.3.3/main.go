package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")

	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

var reLine = regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	taskMap := make(map[string]Task, len(lines))
	for _, line := range lines {
		groups := reLine.FindStringSubmatch(line)
		if len(groups) < 4 {
			return nil, fmt.Errorf("can't parse line: %s", line)
		}

		t1, err := time.Parse("15:04", groups[1])
		if err != nil {
			return nil, fmt.Errorf("can't parse time from string: %s, err: %v", groups[1], err)
		}

		t2, err := time.Parse("15:04", groups[2])
		if err != nil {
			return nil, fmt.Errorf("can't parse time from string: %s, err: %v", groups[2], err)
		}

		if !t2.After(t1) {
			return nil, fmt.Errorf("incorrect time range")
		}

		title := groups[3]

		if v, ok := taskMap[title]; !ok {
			taskMap[title] = Task{
				Date:  date,
				Dur:   t2.Sub(t1),
				Title: title,
			}
		} else {
			taskMap[title] = Task{
				Date:  v.Date,
				Dur:   v.Dur + t2.Sub(t1),
				Title: title,
			}
		}
	}

	res := make([]Task, 0, len(taskMap))
	for _, v := range taskMap {
		res = append(res, v)
	}

	return res, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
