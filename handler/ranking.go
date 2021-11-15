package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tdu-logcation/api/controller"
	"github.com/tdu-logcation/api/utils"
)

// ランキングの情報取得
// 認証なし
func RnakingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := controller.NewUser(&ctx)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rank, err := users.Rank()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ranks を byteに変換
	body, err := utils.ToJson(rank)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
