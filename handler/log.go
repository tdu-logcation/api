package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tdu-logcation/api/controller"
	"github.com/tdu-logcation/api/utils"
)

// POSTフォームのJSON形式
type Log struct {
	Date string `json:"date"`

	// キャンパス
	// 千住 or 鳩山
	Campus string `json:"campus"`

	//ログ定期
	// 詳しくはlogcation/webのソース参照: https://github.com/tdu-logcation/web/blob/0d9feacdd50c5edfc1c28ab5a4510a3370715173/%40types/log.ts#L9-L11
	LogType string `json:"log_type"`

	// ログのラベル
	Label string ` json:"label"`

	// ログデータ
	Code string `json:"code"`
}

type Logs struct {
	Id   string `json:"id"`
	Logs []Log  `json:"logs"`
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		LogGetHandler(w, r)
	case http.MethodPost:
		LogPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ログ取得
func LogGetHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザのIDを取得
	// クエリパラメータ: ?id=hogehoge
	ctx := r.Context()

	id, err := utils.GetQuery(r, "id")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log, err := controller.NewLog(&ctx, id)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logs, err := log.GetLogs()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := utils.ToJson(*logs)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// ログ追加
func LogPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()

	var logs Logs

	json.NewDecoder(r.Body).Decode(&logs)

	user, err := controller.NewUser(&ctx)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := user.PlusLog(logs.Id, len(logs.Logs)); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, log := range logs.Logs {
		// dateはRFC3339形式を前提とする
		timeDate, err := time.Parse(time.RFC3339, log.Date)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_log, err := controller.NewLog(&ctx, logs.Id)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := _log.Add(log.Campus, timeDate, log.LogType, log.Label, log.Code); err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
