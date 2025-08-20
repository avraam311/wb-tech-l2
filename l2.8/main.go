package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func currentTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения текущего времени:", err.Error())
		os.Exit(1)
	}

	return time
}

func main() {
	curTime := currentTime()
	fmt.Println(curTime)
}
