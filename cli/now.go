package cli

import (
	"fmt"
	"time"
)

func commandNow() error {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
	return nil
}
