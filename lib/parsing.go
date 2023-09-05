package lib

import (
	"fmt"
	"time"
)

func ParseTime(layout, value string) {
	data, err := time.Parse(layout, value)

	if err != nil {
		fmt.Println(err)
	}

	if data.IsZero() {
		return
	}

	fmt.Printf("%s\t->\t%v\n", value, data)
}
