package core

import (
	"time"
)

func DefaultIfUnset[T comparable](v, d, u T) T {
	if v == u {
		return d
	}
	return v
}

func LayoutAsGoTimeLayout(v string) string {
	switch v {
	case "RFC1123":
		return time.RFC1123
	case "RFC3339":
		return time.RFC3339
	}
	return v
}

func TimeToHuman(t time.Time, config Config) (string, error) {
	location, err := time.LoadLocation(config.Location)
	if err != nil {
		return "", err
	}
	timeLayout := LayoutAsGoTimeLayout(config.Layout)
	return t.In(location).Format(timeLayout), nil
}
