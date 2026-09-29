package cli

import (
	"fmt"
	"time"

	"github.com/enchant97/time-tool/core"
)

func commandNow(config core.Config) error {
	var clockOffset time.Duration = 0

	if config.NTPClient.Enable {
		resp, err := core.NTPQuery(core.NTPQueryOptions{
			Server: config.NTPClient.Server,
			Timout: time.Duration(config.NTPClient.Timeout) * time.Second,
		})
		if err != nil {
			return err
		}
		clockOffset = resp.ClockOffset
	}

	timeString, err := core.TimeToHuman(time.Now().Add(clockOffset), config)
	if err != nil {
		return err
	}
	fmt.Println(timeString)
	return nil
}
