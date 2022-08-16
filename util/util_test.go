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

func TestReadToken(t *testing.T) {

	const missing string = "missing_config.json"

	if _, err := util.ReadToken(missing); err == nil {
		t.Error("Should not be able to read the token", err)
	}

}
