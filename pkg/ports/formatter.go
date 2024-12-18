package ports

import "time"

type FormatterService interface {
	FormatThaiTime(t time.Time) string
	FormatThaiDate(t time.Time) string
}
