package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config := LoadCustomize("../../config/users.json")

	if len(config.Databases) <= 0 {
		t.Fail()
	}

	if config.PublicUrl == "" {
		t.Fail()
	}

	if config.ListenAddress == "" {
		t.Fail()
	}
}
