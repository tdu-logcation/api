package utils

import "time"

func NowTime() *time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	return &now
}
