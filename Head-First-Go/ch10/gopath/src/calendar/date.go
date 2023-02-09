package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

// setter 方法，用来设置字段或者基础类型中的其他值。
// 依照惯例，Go 的 setter 方法名为 SetX 的形式，X 是你想要设置的东西的名称。
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 12 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// getter 方法
// 通常getter方法命名为X，X就是获取的值的字段名称。
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
