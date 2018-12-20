package utils

import "time"

func TimeFormat(timeStamp int64) string {

	format := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")

	return format;
}
