package util

import (
	"fmt"
	"strconv"
	"time"
)

func CheckCid(cid string) bool {
	if len(cid) == 13 {
		digits := make([]int, 0, 13)
		for _, c := range cid {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				return false
			}
			digits = append(digits, d)
		}
		lastDigit := digits[len(digits)-1]
		digits = digits[:len(digits)-1]

		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}

		sum := 0
		for k, d := range digits {
			sum += (k + 2) * d
		}
		checkDigit := (11 - (sum % 11)) % 10
		return lastDigit == checkDigit
	}
	return false
}

func FormatThaiDate(t time.Time) string {
	thaiWeekdays := []string{"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"}
	thaiMonths := []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}

	weekday := thaiWeekdays[t.Weekday()]
	day := t.Day()
	month := thaiMonths[int(t.Month())-1]
	year := t.Year() + 543
	return fmt.Sprintf("%sที่ %d %s %d", weekday, day, month, year)
}
