package timer

import "time"

func GetTime() string {
	return time.Now().String()[0:19]
}

func GetTimeDurationMillisecond(ms uint64) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func GetTimeDurationSecond(s uint64) time.Duration {
	return time.Duration(s) * time.Second
}

func GetTimeDurationMinute(m uint64) time.Duration {
	return time.Duration(m) * time.Minute
}

func GetTimeDurationHour(h uint64) time.Duration {
	return time.Duration(h) * time.Hour
}
