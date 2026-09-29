package cli

import (
	"os"
	"time"

	"github.com/enchant97/time-tool/core"
)

func commandNtpV4Query(server string, enableNTS bool, timeout uint16) error {
	opt := core.NTPQueryOptions{
		Server:    server,
		EnableNTS: enableNTS,
		Timout:    time.Duration(timeout) * time.Second,
	}
	response, err := core.NTPQuery(opt)
	if err == nil {
		response.Log(os.Stdout)
	}
	return err
}
