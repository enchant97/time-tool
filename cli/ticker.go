package cli

import (
	"fmt"
	"time"

	"github.com/enchant97/time-tool/core"
)

func commandTicker(config core.Config) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		timeString, err := core.TimeToHuman(t, config)
		if err != nil {
			return err
		}
		fmt.Println(timeString)
	}
}
