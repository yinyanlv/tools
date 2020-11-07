package time

import "time"

const (
	NanoSecond  time.Duration = 1
	MicroSecond               = NanoSecond * 1000
	MillsSecond               = MicroSecond * 1000
	Second                    = MillsSecond * 1000
	Minute                    = Second * 60
	Hour                      = Minute * 60
	Day                       = Hour * 24
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, nil
	}
	return currentTime.Add(duration), nil
}
