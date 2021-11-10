package utils

import "time"

func NowTime() (*time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	now := time.Now().In(jst)

	return &now, nil
}
