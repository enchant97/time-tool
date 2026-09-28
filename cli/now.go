package cli

import (
	"fmt"
	"time"

	"github.com/enchant97/time-tool/core"
)

func commandNow(config core.Config) error {
	timeString, err := core.TimeToHuman(time.Now(), config)
	if err != nil {
		return err
	}
	fmt.Println(timeString)
	return nil
}
