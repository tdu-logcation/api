package controller

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/tdu-logcation/api/database"
	"github.com/tdu-logcation/api/utils"
)

type Log struct {
	userId    string
	userIdKey string
	database  *database.Database
}

func NewLog(ctx *context.Context, userId string) (*Log, error) {
	if len(userId) == 0 {
		return nil, errors.New("user id is none")
	}

	database, err := database.New(ctx)
	if err != nil {
		return nil, err
	}

	userIdKey := utils.ConvertUserIdKey(userId)

	return &Log{
		userId:    userId,
		userIdKey: userIdKey,
		database:  database,
	}, nil
}

// ログの追加
func (c *Log) Add(campus string, date time.Time, logType string, label string, code string) error {
	logId, err := utils.CreateId()
	if err != nil {
		return err
	}

	key := utils.CreateKey(c.userIdKey, logId)

	entry := database.Log{
		Id:      logId,
		Date:    date,
		Campus:  campus,
		LogType: logType,
		Label:   label,
		Code:    code,
	}

	return c.database.Put(key, &entry)
}

// 全ログ取得
func (c *Log) GetLogs() (*[]database.Log, error) {
	query := datastore.NewQuery(c.userIdKey)
	var posts []database.Log

	_, err := c.database.GetAll(query, &posts)
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

// 全ログ削除
func (c *Log) DeleteAll() error {
	logs, err := c.GetLogs()
	if err != nil {
		return err
	}

	keys := []*datastore.Key{}

	for _, log := range *logs {
		keys = append(keys, utils.CreateKey(c.userIdKey, log.Id))
	}

	return c.database.DeleteMulti(keys)
}
