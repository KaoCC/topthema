package util_test

import (
	"testing"
	"time"
	"topthema/util"
)

func TestGetLastTime(t *testing.T) {

	const missing string = "missing_test_file.txt"

	if testTime := util.GetLastTime(missing); !testTime.Equal(time.Time{}) {
		t.Error("Not equal to default")
	}

}
