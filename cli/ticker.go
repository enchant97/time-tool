package cli

import (
	"fmt"
	"time"
)

func commandTicker() error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		fmt.Println(t.Format(time.RFC3339))
	}
}
