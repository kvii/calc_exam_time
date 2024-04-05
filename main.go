package main

import (
	"fmt"
	"time"
)

func main() {
	var used time.Duration
	for _, vs := range [][]time.Time{
		{t("21:24"), t("22:36")},
		{t("10:35"), t("11:08")},
	} {
		begin, end := vs[0], vs[1]
		duration := end.Sub(begin)
		used += duration
		fmt.Printf("used:\t%v\n", duration)
	}
	fmt.Printf("total:\t%v\n", used)

	const total = 2*time.Hour + 30*time.Minute
	left := total - used
	fmt.Printf("left:\t%v\n", left)
}

func t(s string) time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		panic(err)
	}
	return t
}
