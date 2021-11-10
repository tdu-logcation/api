package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tdu-logcation/api/controller"
	"github.com/tdu-logcation/api/utils"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UserGetHandler(w, r)
	case http.MethodPost:
		UserPostHandler(w, r)
	case http.MethodDelete:
		UserDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// アカウント作成と、ユーザ名変更
func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.ParseForm()

	// ユーザ名
	// アカウント作成、ユーザ名変更どちらも必須
	userName := r.PostFormValue("user_name")

	// ユーザのID
	// ユーザ名変更時のみ使用
	id := r.PostFormValue("id")

	user, err := controller.NewUser(&ctx)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(id) == 0 {
		// アカウント作成
		info, err := user.Add(userName)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := utils.ToJson(*info)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	} else {
		// 名前変更
		if err := user.ChangeName(id, userName); err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// ユーザ情報を取得
//
//	- ユーザID
//	- ユーザ名
//	- アカウント作成日
//	- ログ総数
func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザのIDを取得
	// クエリパラメータ: ?id=hogehoge
	id, err := utils.GetQuery(r, "id")
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

	info, err := user.Get(id)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := utils.ToJson(*info)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// ユーザ削除
func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザのIDを取得
	// クエリパラメータ: ?id=hogehoge
	id, err := utils.GetQuery(r, "id")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ログデータ削除
	log, err := controller.NewLog(&ctx, id)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := log.DeleteAll(); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザ情報削除
	user, err := controller.NewUser(&ctx)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := user.Delete(id); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
