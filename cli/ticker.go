package cli

import (
	"fmt"
	"time"

	"github.com/enchant97/time-tool/core"
)

func commandTicker(config core.Config) error {
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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		timeString, err := core.TimeToHuman(t.Add(clockOffset), config)
		if err != nil {
			return err
		}
		fmt.Println(timeString)
	}
}
