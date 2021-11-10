package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/tdu-logcation/api/controller"
)

func createUser(t *testing.T) controller.User {
	ctx := context.Background()

	user, err := controller.NewUser(&ctx)
	if err != nil {
		t.Fatal(err)
	}

	return *user
}

func TestUser(t *testing.T) {
	name := "Cateiru"

	user := createUser(t)

	info, err := user.Add(name)
	if err != nil {
		t.Fatal(err)
	}

	if info.Name != name {
		t.Fatalf("名前が違います: %v != %v", info.Name, name)
	}

	id := info.Id

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	// ------
	// ログ数カウント
	for i := 0; 10 > i; i++ {
		if err := user.PlusLog(id); err != nil {
			t.Fatal(err)
		}
	}

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	info2, err := user.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if info2.NumberOfLogs != 10 {
		t.Fatalf("ログ数が10ではありません: %v", info2.NumberOfLogs)
	}

	// -----
	// 名前変更
	newName := "Nya"

	if err := user.ChangeName(id, newName); err != nil {
		t.Fatal(err)
	}

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	info3, err := user.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if info3.Name != newName {
		t.Fatalf("名前が変更されていません: %v != %v", info3.Name, name)
	}

	// -----
	// 削除
	if err := user.Delete(id); err != nil {
		t.Fatal(err)
	}

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	deletedInfo, err := user.Get(id)
	if err == nil {
		t.Fatalf("ユーザが削除されていません: %v", deletedInfo)
	}
}

func TestEmptyUserGet(t *testing.T) {
	id := "nyanya"
	user := createUser(t)

	deletedInfo, err := user.Get(id)
	if err == nil {
		t.Fatalf("謎なユーザが取得できてしまっています: %v", deletedInfo)
	}
}

func TestEmptyUserDelete(t *testing.T) {
	user := createUser(t)
	id := "hoge"

	if err := user.Delete(id); err != nil {
		t.Fatal(err)
	}
}

func TestEmptyChangeName(t *testing.T) {
	user := createUser(t)
	id := "asdad"
	name := "tomato"

	if err := user.ChangeName(id, name); err == nil {
		t.Fatal("不明なユーザの名前変更でエラーが起きません")
	}
}

func TestRank(t *testing.T) {
	names := []string{
		"Okayu",
		"EVA",
		"Godzilla",
	}

	user1 := createUser(t)
	user2 := createUser(t)
	user3 := createUser(t)

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	info1, err := user1.Add(names[0])
	if err != nil {
		t.Fatal(err)
	}
	info2, err := user2.Add(names[1])
	if err != nil {
		t.Fatal(err)
	}
	info3, err := user3.Add(names[2])
	if err != nil {
		t.Fatal(err)
	}

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	// ユーザ1(Okayu): ログ数5
	for i := 0; 5 > i; i++ {
		if err := user1.PlusLog(info1.Id); err != nil {
			t.Fatal(err)
		}
	}

	// ユーザ2(EVA): ログ数3
	for i := 0; 3 > i; i++ {
		if err := user2.PlusLog(info2.Id); err != nil {
			t.Fatal(err)
		}
	}

	// ユーザ3(Godzilla): ログ数7
	for i := 0; 7 > i; i++ {
		if err := user3.PlusLog(info3.Id); err != nil {
			t.Fatal(err)
		}
	}

	// 反映される待つ！
	time.Sleep(1 * time.Second)

	rank, err := user1.Rank()
	if err != nil {
		t.Fatal(err)
	}

	if rank[0] != names[2] {
		t.Fatalf("ランキング1位が違います: %v", rank)
	}
	if rank[1] != names[0] {
		t.Fatalf("ランキング2位が違います: %v", rank)
	}
	if rank[2] != names[1] {
		t.Fatalf("ランキング3位が違います: %v", rank)
	}
}
