package timer

import "time"

// 获取当前时间
func GetNowTime() time.Time {
	localtion, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(localtion)
}

// 时间推算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
