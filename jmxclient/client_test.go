package jmxclient

import "testing"

func TestConfig(t *testing.T) {
	config, err := GetConfig("config.yml")
	if err != nil {
		t.Error(err)
	}
	expect := "http://localhost:8080/manager/jmxproxy"
	actual := config.Endpoint
	if actual != expect {
		t.Errorf("Invalid Config, got %s, should be %s", actual, expect)
	}
	//Start(*config)
}
