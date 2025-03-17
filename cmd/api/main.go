package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/nimilgp/URLcommentary/internal/database"
	"github.com/nimilgp/URLcommentary/internal/env"
	"github.com/nimilgp/URLcommentary/internal/smtp"
	"github.com/nimilgp/URLcommentary/internal/version"

	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL   string
	httpPort  int
	basicAuth struct {
		username       string
		hashedPassword string
	}
	cookie struct {
		secretKey string
	}
	db struct {
		dsn         string
		automigrate bool
	}
	jwt struct {
		secretKey string
	}
	notifications struct {
		email string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		from     string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *slog.Logger
	mailer *smtp.Mailer
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.basicAuth.username = env.GetString("BASIC_AUTH_USERNAME", "admin")
	cfg.basicAuth.hashedPassword = env.GetString("BASIC_AUTH_HASHED_PASSWORD", "$2a$10$jRb2qniNcoCyQM23T59RfeEQUbgdAXfR6S0scynmKfJa5Gj3arGJa")
	cfg.cookie.secretKey = env.GetString("COOKIE_SECRET_KEY", "5jbd43jllpct5anck2pdlnit35mtg75j")
	cfg.db.dsn = env.GetString("DB_DSN", "db.sqlite")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)
	cfg.jwt.secretKey = env.GetString("JWT_SECRET_KEY", "7cw7urefmco3m4ok3j7ivi2ayqthg64u")
	cfg.notifications.email = env.GetString("NOTIFICATIONS_EMAIL", "")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Example Name <no_reply@example.org>")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	mailer, err := smtp.NewMailer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.from)
	if err != nil {
		return err
	}

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
		mailer: mailer,
	}

	return app.serveHTTP()
}
