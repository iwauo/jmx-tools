package jmxclient

import "testing"

func TestConfig(t *testing.T) {
	expect := "JMX/PATH"
	actual := Config()
	if actual != expect {
		t.Errorf("Invalid Config, got %s, should be %s", actual, expect)
	}
}
