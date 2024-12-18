package formatter

import (
	"fmt"
	"time"

	"github.com/Be2Bag/example/pkg/ports"
)

type formatterService struct {
}

func NewFormatterService() ports.FormatterService {
	return &formatterService{}
}

func (f *formatterService) FormatThaiTime(t time.Time) string {
	t = t.In(time.FixedZone("UTC+7", 7*60*60))
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	return fmt.Sprintf("เวลา %02d:%02d:%02d น.", hour, minute, second)
}

func (f *formatterService) FormatThaiDate(t time.Time) string {
	thaiWeekdays := []string{"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"}
	thaiMonths := []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}

	weekday := thaiWeekdays[t.Weekday()]
	day := t.Day()
	month := thaiMonths[int(t.Month())-1]
	year := t.Year() + 543
	return fmt.Sprintf("%sที่ %d %s %d %s", weekday, day, month, year, f.FormatThaiTime(t))
}
