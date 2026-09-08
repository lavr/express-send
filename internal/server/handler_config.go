package server

import (
	"encoding/json"
	"net/http"

	"github.com/lavr/express-botx/internal/config"
)

func (s *Server) handleBotList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.botEntries)
}

func (s *Server) handleChatsAliasList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	entries := s.chatEntries
	if s.scopedKey(r.Context()) {
		visible := make([]config.ChatEntry, 0, len(entries))
		for _, e := range entries {
			if s.chatAllowed(r.Context(), e.ID) {
				visible = append(visible, e)
			}
		}
		entries = visible
	}
	json.NewEncoder(w).Encode(entries)
}

// WithConfigInfo sets the bot and chat entries for the config endpoints.
func WithConfigInfo(bots []config.BotEntry, chats []config.ChatEntry) Option {
	return func(s *Server) {
		s.botEntries = bots
		s.chatEntries = chats
	}
}
