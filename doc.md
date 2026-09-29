## Layouts
The time can be displayed in different layouts. Use the below table to find the desired format. You can also specify standard go time layout values.

| Name | Example |
| :--- | :------ |
| RFC822 | '02 Jan 06 15:04 MST' |
| RFC822Z | '02 Jan 06 15:04 -0700' |
| RFC850 | 'Monday, 02-Jan-06 15:04:05 MST' |
| RFC1123 | 'Mon, 02 Jan 2006 15:04:05 MST' |
| RFC1123Z | 'Mon, 02 Jan 2006 15:04:05 -0700' |
| RFC3339 | '2006-01-02T15:04:05Z07:00' |
| RFC3339Nano | '2006-01-02T15:04:05.999999999Z07:00' |
| Date | '2006-01-02' |
| Time | '15:04:05' |
| Time12 | '03:04:05 PM' |
| DateTime | '2006-01-02 15:04:05' |
| DateTime12 | '2006-01-02 03:04:05 PM' |
| Unix | 'Unix' |

## Locations
The time can be displayed in different locations. Provide a IANA Time Zone location name. The location data will either be loaded from the embedded database or from the system.

Example Name:

    Europe/London

## Config File
The application will store defaults (set via the TUI) in a config file stored the user config directory under `time-tool/config.toml`.
