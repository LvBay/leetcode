package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//关联主体
type RelateInfo struct {
	RelateType      int
	CheckObjectType int
}

func toString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

type Person struct {
	list []int
}

func (p *Person) Shift() {
	if len(p.list) > 1 {
		p.list = p.list[1:]
	} else {
		p.list = nil
	}
}

func main() {
	p := &Person{list: []int{1, 2, 3}}
	for i, v := range p.list {
		fmt.Println(i, v)
		p.Shift()
	}
	fmt.Println(p)
}

func Swt(t int) {
	switch {
	case t == 1:
		fmt.Println("1")
	case t == 2:
		fmt.Println("2")
	case t == 3:
	case t >= 3:
		fmt.Println(">=3")
	default:
		fmt.Println("default")
	}
}

// days for t1-t2
func SubForDay(t1, t2 time.Time) (days int) {
	i := t2.Sub(t1) / time.Hour / 24
	return int(i)
}

func NormalRpDateList(y, m, d, d2, stage int) (list []time.Time) {
	now := time.Now()
	t1 := Date(y, m, d)
	t2 := Date(now.Year(), int(now.Month()), d2)
	if SubForDay(t1, t2) < 15 {
		t2 = t2.AddDate(0, 1, 0)
	} else if SubForDay(t1, t2) > 40 {
		t2 = t2.AddDate(0, -1, 0)
	}
	for i := 0; i <= stage; i++ {
		date := t2.AddDate(0, i, 0)
		list = append(list, date)
		fmt.Println(date)
	}
	return list
}

// 根据签约日期，还款日，期数生成应还款日数组
func GetRpDateList(y, m, d, stage int) (list []time.Time) {
	if d != 1 { // 还款日为1，从当月开始计算，否则从下月开始计算
		m++
	}
	list = make([]time.Time, 0, stage)
	for i := 0; i < stage; i++ {
		// date := Max30Days(y, m)
		date := LastDay(y, m)
		date = Max30Days(date)
		list = append(list, date)
		y, m = NextMonth(y, m)
		fmt.Println(date)
	}
	return list
}

// 返回下个月的年份yy，月份mm
func NextMonth(y, m int) (yy, mm int) {
	m++
	if m == 13 {
		y++
		m = 1
	}
	return y, m
}

// y年m月的最后一天
func LastDay(y, m int) time.Time {
	d := Count(y, m)      // 计算y年m月的天数
	date := Date(y, m, d) // y年m月dt日
	return date
}

func Max30Days(t time.Time) time.Time {
	d := t.Day()
	if d != 31 {
		return t
	}
	return t.AddDate(0, 0, -1)
}

// Count 返回y年m月的天数
func Count(y int, m int) (days int) {
	t1 := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(y, time.Month(m+1), 1, 0, 0, 0, 0, time.Local)
	return int(t2.Sub(t1) / time.Hour / 24)
}

// 是否为闰年
func IsLeap(y int) bool {
	if y%100 == 0 {
		return y%400 == 0
	}
	return y%4 == 0
}

// 返回y年m月d日
func Date(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
}
