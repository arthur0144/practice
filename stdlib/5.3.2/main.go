package main

import (
	"fmt"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour int
	min  int
	sec  int
	loc  *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.min
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.sec
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.hour, t.min, t.sec, t.loc.String())
}

// Location возвращает строкове представление таймзоны
func (t TimeOfDay) Location() string {
	return t.loc.String()
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	return t.Location() == other.Location() && t.hour == other.hour && t.min == other.min && t.sec == other.sec
}

var errNotEqualLocation = fmt.Errorf("not equal location")

// seconds возвращает количество секунд с начала дня
func (t TimeOfDay) seconds() int {
	return t.sec + t.min*60 + t.hour*60*60
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.Location() != other.Location() {
		return false, errNotEqualLocation
	}
	return t.seconds() < other.seconds(), nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.Location() != other.Location() {
		return false, errNotEqualLocation
	}
	return t.seconds() > other.seconds(), nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{hour, min, sec, loc}
}

// конец решения

func main() {

}
