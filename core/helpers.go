package core

import (
	"fmt"
	"time"
)

func DefaultIfUnset[T comparable](v, d, u T) T {
	if v == u {
		return d
	}
	return v
}

func LayoutAsGoTimeLayout(v string) string {
	if v, ok := TimeLayoutMappings[v]; ok {
		return v
	}
	return v
}

func TimeToHuman(t time.Time, config Config) (string, error) {
	switch config.Layout {
	case "Unix":
		return fmt.Sprintf("%d", t.Unix()), nil
	}
	location, err := time.LoadLocation(config.Location)
	if err != nil {
		return "", err
	}
	timeLayout := LayoutAsGoTimeLayout(config.Layout)
	return t.In(location).Format(timeLayout), nil
}
