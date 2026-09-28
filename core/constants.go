package core

import "time"

var TimeLayoutMappings = map[string]string{
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Date":        time.DateOnly,
	"Time":        time.TimeOnly,
	"Time12":      "03:04:05 PM",
	"DateTime":    time.DateTime,
	"DateTime12":  "2006-01-02 03:04:05 PM",
	"Unix":        "Unix",
}
