package handler

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tdu-logcation/api/controller"
	"github.com/tdu-logcation/api/utils"
)

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

	id := r.PostFormValue("id")
	date := r.PostFormValue("date")
	campus := r.PostFormValue("campus")
	logType := r.PostFormValue("log_type")
	label := r.PostFormValue("label")
	code := r.PostFormValue("code")

	// dateはRFC3339Nano形式を前提とする
	timeDate, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := controller.NewUser(&ctx)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := user.PlusLog(id); err != nil {
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

	if err := log.Add(campus, timeDate, logType, label, code); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
