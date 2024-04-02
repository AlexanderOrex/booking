package model

import "time"

type Date time.Time

func DateFromTime(t time.Time) Date {
	return Date(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC))
}

func (d Date) ToTime() time.Time {
	t := time.Time(d)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func (d Date) After(date Date) bool {
	return time.Time(d).After(time.Time(date))
}

func (d Date) Before(date Date) bool {
	return time.Time(d).Before(time.Time(date))
}

func (d Date) AddDay() Date {
	return Date(time.Time(d).AddDate(0, 0, 1))
}
