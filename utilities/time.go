package utilities

import (
	"online-store/constants"
	"time"
)

func Times() *string {
	now := time.Now()
	format := constants.DateFromStd
	time := now.Format(format)

	return &time
}
