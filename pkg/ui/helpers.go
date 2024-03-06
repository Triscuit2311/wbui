package ui

import (
	"fmt"
)

type integer interface {
	int8 | int16 | int32 | int64 | int
}
type fpnumber interface {
	float32 | float64
}

func itostr[T integer](n T) string {
	return fmt.Sprintf("%d", n)
}

func ftostr[T fpnumber](n T) string {
	return fmt.Sprintf("%.2f", n)
}
