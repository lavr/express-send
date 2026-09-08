package config

import (
	"strings"
	"testing"
)

func TestValidate_APIKeyChatScope(t *testing.T) {
	base := func(chats []string) *Config {
		return &Config{
			Bots:  map[string]BotConfig{"b": {Host: "h", ID: "i", Secret: "s"}},
			Chats: map[string]ChatConfig{"known": {ID: "bcb715a2-e8d3-57a8-ab3b-6a14c044dd22"}},
			Server: ServerConfig{
				APIKeys: []APIKeyConfig{{Name: "k", Key: "v", Chats: chats}},
			},
		}
	}

	tests := []struct {
		name    string
		chats   []string
		wantErr bool
	}{
		{"absent scope", nil, false},
		{"known alias", []string{"known"}, false},
		{"bare uuid", []string{"7ee8aaa9-c6cb-5ee6-8445-7d654819b285"}, false},
		{"empty list", []string{}, true},
		{"unknown alias", []string{"nope"}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var hasErr bool
			for _, r := range base(tc.chats).Validate(nil) {
				if r.Level == ValidationError && strings.Contains(r.Path, "api_keys") {
					hasErr = true
				}
			}
			if hasErr != tc.wantErr {
				t.Errorf("error = %v, want %v", hasErr, tc.wantErr)
			}
		})
	}
}
