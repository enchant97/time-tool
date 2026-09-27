package cli

import (
	"fmt"
	"time"
)

func commandNow(timeLayout string, timeLocation string) error {
	location, err := time.LoadLocation(timeLocation)
	if err != nil {
		return err
	}
	now := time.Now().In(location)
	switch timeLayout {
	case "RFC1123":
		timeLayout = time.RFC1123
	case "RFC3339":
		timeLayout = time.RFC3339
	}
	fmt.Println(now.Format(timeLayout))
	return nil
}
