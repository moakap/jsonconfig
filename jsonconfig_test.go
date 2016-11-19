package jsonconfig_test

import (
	"fmt"
	"git.coding.net/bb_rd/EAgent/jsonconfig"
	"testing"
	"time"
)

func TestJsonConfig(t *testing.T) {
	err := jsonconfig.InitConf("conf.json")
	if err != nil {
		fmt.Println("Initialize config failed: ", err.Error())
		return
	}

	jsonconfig.Set("param1", 100)
	jsonconfig.Set("param2", "stringValue")
	jsonconfig.Set("param3", time.Now())

	jsonconfig.Set("param1", 100)
	if jsonconfig.Get("param1") != 100 {
		t.Fail()
	}

	jsonconfig.Set("param2", "stringValue")
	if jsonconfig.Get("param2") != "stringValue" {
		t.Fail()
	}

	jsonconfig.Set("param2", 101)
	if jsonconfig.Get("param2") != 101 {
		t.Fail()
	}
}
