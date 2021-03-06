package controller

import (
	"context"
	"errors"
	"sort"

	"cloud.google.com/go/datastore"
	"github.com/tdu-logcation/api/database"
	"github.com/tdu-logcation/api/utils"
)

type User struct {
	database *database.Database
}

func NewUser(ctx *context.Context) (*User, error) {
	database, err := database.New(ctx)
	if err != nil {
		return nil, err
	}

	return &User{
		database: database,
	}, nil
}

// ユーザを追加する
func (c *User) Add(name string) (*database.User, error) {
	id, err := utils.CreateId()
	if err != nil {
		return nil, err
	}
	now := utils.NowTime()

	userInfo := database.User{
		Id:           id,
		Name:         name,
		CreateDate:   *now,
		NumberOfLogs: 0,
	}

	c.set(id, &userInfo)

	return &userInfo, nil
}

// ユーザ情報を取得する
func (c *User) Get(id string) (*database.User, error) {
	key := c.createKey(id)
	entity := new(database.User)

	if err := c.database.Get(key, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

// ユーザを追加or更新する
func (c *User) set(id string, entity *database.User) error {
	key := c.createKey(id)

	return c.database.Put(key, entity)
}

// ログ数をインクリメントする
func (c *User) PlusLog(id string, pulsValue int) error {
	userInfo, err := c.Get(id)
	if err != nil {
		return err
	}

	userInfo.NumberOfLogs += pulsValue

	return c.set(id, userInfo)
}

func (c *User) ChangeName(id string, name string) error {
	userInfo, err := c.Get(id)
	if err != nil {
		return err
	}
	if userInfo == nil {
		return errors.New("no user")
	}

	userInfo.Name = name

	return c.set(id, userInfo)
}

// ユーザを削除する
func (c *User) Delete(id string) error {
	key := c.createKey(id)
	return c.database.Delete(key)
}

// datastoreのkeyを作成
func (c *User) createKey(id string) *datastore.Key {
	return utils.CreateKey("users", id)
}

// ランキング取得
func (c *User) Rank() (database.Ranks, error) {
	query := datastore.NewQuery("users")
	var posts []database.User

	_, err := c.database.GetAll(query, &posts)
	if err != nil {
		return nil, err
	}

	var ranks database.Ranks = []database.Rank{}

	for _, value := range posts {
		ranks = append(ranks, database.Rank{
			Name:         value.Name,
			NumberOfLogs: value.NumberOfLogs,
		})
	}

	sort.Sort(ranks)

	// reverse rank
	for i, j := 0, len(ranks)-1; i < j; i, j = i+1, j-1 {
		ranks[i], ranks[j] = ranks[j], ranks[i]
	}

	return ranks, nil
}
