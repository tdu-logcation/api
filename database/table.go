package database

import "time"

type User struct {
	Id string `datastore:"id" json:"id"`

	// ユーザ名
	Name         string    `datastore:"name" json:"name"`
	CreateDate   time.Time `datastore:"create_date" json:"create_date"`
	NumberOfLogs int       `datastore:"number_of_logs" json:"number_of_logs"`
}

type Log struct {
	Id   string    `datastore:"id" json:"id"`
	Date time.Time `datastore:"date" json:"date"`

	// キャンパス
	// 千住 or 鳩山
	Campus string `datastore:"campus" json:"campus"`

	//ログ定期
	// 詳しくはlogcation/webのソース参照: https://github.com/tdu-logcation/web/blob/0d9feacdd50c5edfc1c28ab5a4510a3370715173/%40types/log.ts#L9-L11
	LogType string `datastore:"log_type" json:"log_type"`

	// ログのラベル
	Label string `datastore:"label" json:"label"`

	// ログデータ
	Code string `datastore:"code" json:"code"`
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u Users) Less(i, j int) bool {
	return u[i].NumberOfLogs < u[j].NumberOfLogs
}
