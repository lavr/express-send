package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"

	"github.com/lavr/express-send/internal/auth"
	"github.com/lavr/express-send/internal/config"
	"github.com/lavr/express-send/internal/input"
	"github.com/lavr/express-send/internal/secret"
	"github.com/lavr/express-send/internal/sender"
	"github.com/lavr/express-send/internal/token"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var flags config.Flags
	var noCache bool

	flag.StringVar(&flags.ConfigPath, "config", "", "path to config file (default: ~/.config/express-send/config.yaml)")
	flag.StringVar(&flags.Host, "host", "", "eXpress server host")
	flag.StringVar(&flags.BotID, "bot-id", "", "bot UUID")
	flag.StringVar(&flags.Secret, "secret", "", "bot secret (literal, env:VAR, or vault:path#key)")
	flag.StringVar(&flags.ChatID, "chat-id", "", "target chat UUID")
	flag.StringVar(&flags.MessageFrom, "message-from", "", "read message from file")
	flag.BoolVar(&noCache, "no-cache", false, "disable token caching")
	flag.Usage = usage
	flag.Parse()
	flags.NoCache = noCache

	// Load config (YAML → env → flags)
	cfg, err := config.Load(flags)
	if err != nil {
		return err
	}

	// Read message
	isTerminal := term.IsTerminal(int(os.Stdin.Fd()))
	message, err := input.ReadMessage(flags.MessageFrom, flag.Args(), os.Stdin, isTerminal)
	if err != nil {
		return err
	}

	// Resolve secret
	secretKey, err := secret.Resolve(cfg.Secret)
	if err != nil {
		return fmt.Errorf("resolving secret: %w", err)
	}

	// Build signature
	signature := auth.BuildSignature(cfg.BotID, secretKey)

	// Setup token cache
	cache := newCache(cfg.Cache)

	ctx := context.Background()
	ttl := time.Duration(cfg.Cache.TTL) * time.Second
	cacheKey := cfg.BotID

	// Try cached token
	tok, _ := cache.Get(ctx, cacheKey)

	if tok == "" {
		tok, err = auth.GetToken(ctx, cfg.Host, cfg.BotID, signature)
		if err != nil {
			return fmt.Errorf("getting token: %w", err)
		}
		cache.Set(ctx, cacheKey, tok, ttl)
	}

	// Send message
	err = sender.Send(ctx, cfg.Host, tok, cfg.ChatID, message)
	if err != nil {
		// Retry once on 401 with fresh token
		if errors.Is(err, sender.ErrUnauthorized) {
			tok, err = auth.GetToken(ctx, cfg.Host, cfg.BotID, signature)
			if err != nil {
				return fmt.Errorf("refreshing token: %w", err)
			}
			cache.Set(ctx, cacheKey, tok, ttl)
			err = sender.Send(ctx, cfg.Host, tok, cfg.ChatID, message)
		}
		if err != nil {
			return fmt.Errorf("sending message: %w", err)
		}
	}

	return nil
}

func newCache(cfg config.CacheConfig) token.Cache {
	switch cfg.Type {
	case "file":
		path := cfg.FilePath
		if path == "" {
			if dir, err := os.UserCacheDir(); err == nil {
				path = dir + "/express-send/tokens.json"
			} else {
				home, _ := os.UserHomeDir()
				path = home + "/.cache/express-send/tokens.json"
			}
		}
		return &token.FileCache{Path: path}
	case "vault":
		return &token.VaultCache{
			URL:   cfg.VaultURL,
			Path:  cfg.VaultPath,
			Token: os.Getenv("VAULT_TOKEN"),
		}
	default:
		return token.NoopCache{}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: express-send [options] [message]

Send messages to eXpress messenger via BotX API.

Message sources (in priority order):
  --message-from FILE   Read message from file
  [message]             Positional argument
  stdin                 Pipe input (auto-detected)

Options:
`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
Examples:
  express-send "Hello, world!"
  express-send --message-from report.txt
  echo "Deploy OK" | express-send
  express-send --secret env:BOT_SECRET "Build passed"
  express-send --secret "vault:secret/data/express#bot_secret" "Alert"

Configuration:
  Config file: ~/.config/express-send/config.yaml
  Environment: EXPRESS_HOST, EXPRESS_BOT_ID, EXPRESS_SECRET, EXPRESS_CHAT_ID
  Priority: config file < env vars < CLI flags
`)
}
