/*!
 * Copyright 2024 Vinicius Benevides
 *
 * Project: github.com/VBenevides/cronexpr
 * File: cronexpr_prev.go
 * Version: 1.0
 * License: pick the one which suits you :
 *   GPL v3 see <https://www.gnu.org/licenses/gpl.html>
 *   APL v2 see <http://www.apache.org/licenses/LICENSE-2.0>
 *
 */

package cronexpr

/******************************************************************************/

import (
	"time"
)

/******************************************************************************/

func (expr *Expression) prevYear(t time.Time) time.Time {
	// Find index at which item in list is greater or equal to
	// candidate year
	i := SearchIntsDescending(expr.yearList, t.Year()-1)
	if i == len(expr.yearList) {
		return time.Time{}
	}
	i = len(expr.yearList) - 1 - i

	// Year changed, need to recalculate actual days of month
	expr.actualDaysOfMonthList = expr.calculateActualDaysOfMonth(expr.yearList[i], expr.monthList[len(expr.monthList)-1])
	if len(expr.actualDaysOfMonthList) == 0 {
		return expr.prevMonth(time.Date(
			expr.yearList[i],
			time.Month(expr.monthList[len(expr.monthList)-1]),
			1,
			expr.hourList[len(expr.hourList)-1],
			expr.minuteList[len(expr.minuteList)-1],
			expr.secondList[len(expr.secondList)-1],
			0,
			t.Location()))
	}
	return time.Date(
		expr.yearList[i],
		time.Month(expr.monthList[len(expr.monthList)-1]),
		expr.actualDaysOfMonthList[len(expr.actualDaysOfMonthList)-1],
		expr.hourList[len(expr.hourList)-1],
		expr.minuteList[len(expr.minuteList)-1],
		expr.secondList[len(expr.secondList)-1],
		0,
		t.Location())
}

/******************************************************************************/

func (expr *Expression) prevMonth(t time.Time) time.Time {
	// Find index at which item in list is lower or equal to
	// candidate month
	i := SearchIntsDescending(expr.monthList, int(t.Month())-1)
	if i == len(expr.monthList) {
		return expr.prevYear(t)
	}
	i = len(expr.monthList) - 1 - i

	// Month changed, need to recalculate actual days of month
	expr.actualDaysOfMonthList = expr.calculateActualDaysOfMonth(t.Year(), expr.monthList[i])
	if len(expr.actualDaysOfMonthList) == 0 {
		return expr.prevMonth(time.Date(
			t.Year(),
			time.Month(expr.monthList[i]),
			1,
			expr.hourList[len(expr.hourList)-1],
			expr.minuteList[len(expr.minuteList)-1],
			expr.secondList[len(expr.secondList)-1],
			0,
			t.Location()))
	}

	return time.Date(
		t.Year(),
		time.Month(expr.monthList[i]),
		expr.actualDaysOfMonthList[len(expr.actualDaysOfMonthList)-1],
		expr.hourList[len(expr.hourList)-1],
		expr.minuteList[len(expr.minuteList)-1],
		expr.secondList[len(expr.secondList)-1],
		0,
		t.Location())
}

/******************************************************************************/

func (expr *Expression) prevDayOfMonth(t time.Time) time.Time {
	// Find index at which item in list is lower or equal to
	// candidate day of month
	i := SearchIntsDescending(expr.actualDaysOfMonthList, t.Day()-1)
	if i == len(expr.actualDaysOfMonthList) {
		return expr.prevMonth(t)
	}
	i = len(expr.actualDaysOfMonthList) - 1 - i

	return time.Date(
		t.Year(),
		t.Month(),
		expr.actualDaysOfMonthList[i],
		expr.hourList[len(expr.hourList)-1],
		expr.minuteList[len(expr.minuteList)-1],
		expr.secondList[len(expr.secondList)-1],
		0,
		t.Location())
}

/******************************************************************************/

func (expr *Expression) prevHour(t time.Time) time.Time {
	// Find index at which item in list is lower or equal to
	// candidate hour
	i := SearchIntsDescending(expr.hourList, t.Hour()-1)
	if i == len(expr.hourList) {
		return expr.prevDayOfMonth(t)
	}
	i = len(expr.hourList) - 1 - i

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		expr.hourList[i],
		expr.minuteList[len(expr.minuteList)-1],
		expr.secondList[len(expr.secondList)-1],
		0,
		t.Location())
}

/******************************************************************************/

func (expr *Expression) prevMinute(t time.Time) time.Time {
	// Find index at which item in list is lower or equal to
	// candidate minute
	i := SearchIntsDescending(expr.minuteList, t.Minute()-1)
	if i == len(expr.minuteList) {
		return expr.prevHour(t)
	}
	i = len(expr.minuteList) - 1 - i
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		expr.minuteList[i],
		expr.secondList[len(expr.secondList)-1],
		0,
		t.Location())
}

/******************************************************************************/

func (expr *Expression) prevSecond(t time.Time) time.Time {
	// prevSecond() assumes all other fields are exactly matched
	// to the cron expression

	// Find index at which item in list is lower or equal to
	// candidate second
	i := SearchIntsDescending(expr.secondList, t.Second()-1)
	if i == len(expr.secondList) {
		return expr.prevMinute(t)
	}
	i = len(expr.secondList) - 1 - i

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		expr.secondList[i],
		0,
		t.Location())
}

/******************************************************************************/

func SearchIntsDescending(a []int, x int) int {
	// Let's suppose the list is ordered, but we don't actually care about the order
	// we care how many elements are greater than x (so we know x's index in a
	// descending array)
	count := 0
	for _, val := range a {
		if val > x {
			count++
		}
	}
	return count
}
