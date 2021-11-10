package controller_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/tdu-logcation/api/controller"
	"github.com/tdu-logcation/api/utils"
)

func TestLog(t *testing.T) {
	// os.Setenv("DATASTORE_PROJECT_ID", "logcation")

	ctx := context.Background()
	userId := "user1"

	log, err := controller.NewLog(&ctx, userId)
	if err != nil {
		t.Fatal(err)
	}

	campus := "千住"
	date, err := utils.NowTime()
	if err != nil {
		t.Fatal(err)
	}
	logType := "0"
	label := ""
	code := "jp.ac.dendai/2403-2A"

	if err := log.Add(campus, *date, logType, label, code); err != nil {
		t.Fatal(err)
	}

	// 反映されるまでちょっと待つ
	time.Sleep(1 * time.Second)

	logs, err := log.GetLogs()
	if err != nil {
		t.Fatal(err)
	}

	if len(*logs) != 1 {
		t.Fatalf("ログが無いorたくさんあります: %v", *logs)
	}

	getLog := (*logs)[0]

	if getLog.Date.UTC() == (*date).UTC() {
		t.Fatalf("%v != %v", getLog.Date, date)
	}

	if getLog.Campus != campus ||
		getLog.Code != code ||
		getLog.Label != label ||
		getLog.LogType != logType {
		t.Fatalf("Campus: %v\nCode: %v\nLabel: %v\nLogType: %v\nData: %v",
			getLog.Campus != campus,
			getLog.Code != code,
			getLog.Label != label,
			getLog.LogType != logType,
			getLog)
	}

	if err := log.DeleteAll(); err != nil {
		t.Fatal(err)
	}

	emptyLogs, err := log.GetLogs()
	if err != nil {
		t.Fatal(err)
	}

	if len(*emptyLogs) != 0 {
		t.Fatal("ログは削除したはずなのに、、、")
	}
}

func TestEmptyGet(t *testing.T) {
	os.Setenv("DATASTORE_PROJECT_ID", "logcation")

	ctx := context.Background()
	userId := "user2"

	log, err := controller.NewLog(&ctx, userId)
	if err != nil {
		t.Fatal(err)
	}

	logs, err := log.GetLogs()
	if err != nil {
		t.Fatal(err)
	}

	if len(*logs) != 0 {
		t.Fatal("ログはからである必要があります")
	}
}

func TestEmptyDelete(t *testing.T) {
	os.Setenv("DATASTORE_PROJECT_ID", "logcation")

	ctx := context.Background()
	userId := "user3"

	log, err := controller.NewLog(&ctx, userId)
	if err != nil {
		t.Fatal(err)
	}

	if err := log.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}
