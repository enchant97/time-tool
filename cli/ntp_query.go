package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
	"github.com/beevik/nts"
)

func commandNtpV4Query(server string, enableNTS bool, timeout uint16) error {
	NtpOptions := ntp.QueryOptions{
		Version:                  4,
		RequestSupportedVersions: true,
		Timeout:                  time.Duration(timeout) * time.Second,
	}
	if enableNTS {
		session, err := nts.NewSessionWithOptions(server, &nts.SessionOptions{
			Timeout: time.Duration(timeout) * time.Second,
		})
		if err != nil {
			return err
		}
		response, err := session.QueryWithOptions(&NtpOptions)
		if err != nil {
			return err
		}
		if err := response.Validate(); err != nil {
			fmt.Println(err)
		}
		response.Log(os.Stdout)
	} else {
		response, err := ntp.QueryWithOptions(server, NtpOptions)
		if err != nil {
			return err
		}
		response.Log(os.Stdout)
	}
	return nil
}
