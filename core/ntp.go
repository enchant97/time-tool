package core

import (
	"time"

	"github.com/beevik/ntp"
	"github.com/beevik/nts"
)

type NTPQueryOptions struct {
	Server    string
	EnableNTS bool
	Timout    time.Duration
}

func NTPQuery(opt NTPQueryOptions) (*ntp.Response, error) {
	ntpOptions := ntp.QueryOptions{
		Version:                  4,
		RequestSupportedVersions: true,
		Timeout:                  opt.Timout,
	}
	if opt.EnableNTS {
		// TODO this makes a new session every request,
		//      can this be saved for multiple executions
		session, err := nts.NewSessionWithOptions(opt.Server, &nts.SessionOptions{
			Timeout: opt.Timout,
		})
		if err != nil {
			return nil, err
		}
		response, err := session.QueryWithOptions(&ntpOptions)
		if err != nil {
			return nil, err
		}
		return response, response.Validate()
	}
	response, err := ntp.QueryWithOptions(opt.Server, ntpOptions)
	if err != nil {
		return nil, err
	}
	return response, response.Validate()
}
