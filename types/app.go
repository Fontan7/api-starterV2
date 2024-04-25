package types

import (
	"api-starterV2/storage"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	_   = godotenv.Load(".env")
	Env = os.Getenv("ENV")
)

type App interface {
	Env() string
	Port() string
	SupaKey() string
	SupaProjectReference() string
	ClientKey() string
	JwtKey() string
	Host() string
	DB() *storage.DB
}

type app struct {
	env                  string
	port                 string
	supaKey              string
	supaProjectReference string
	clientKey            string
	jwtKey               string
	host                 string
	db                   *storage.DB
}

func NewApp(db *storage.DB) (App, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("ENV variable is empty")
	}
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT variable is empty")
	}

	supaProjectReference := os.Getenv("SUPA_PROJECT_REFERENCE")
	if supaProjectReference == "" {
		return nil, fmt.Errorf("SUPA_PROJECT_REFERENCE variable is empty")
	}

	supaKey := os.Getenv("SUPA_ANON_KEY")
	if supaKey == "" {
		return nil, fmt.Errorf("SUPA_ANON_KEY variable is empty")
	}

	clientKey := os.Getenv("CLIENT_KEY")
	if clientKey == "" {
		return nil, fmt.Errorf("CLIENT_KEY variable is empty")
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return nil, fmt.Errorf("JWT_KEY variable is empty")
	}

	host := os.Getenv("SERVER_HOST")
	if host == "" {
		return nil, fmt.Errorf("SERVER_HOST variable is empty")
	}

	fmt.Println("Successfully loaded app environment variables")
	return &app{
		env:                  env,
		port:                 port,
		supaKey:              supaKey,
		supaProjectReference: supaProjectReference,
		clientKey:            clientKey,
		jwtKey:               jwtKey,
		host:                 host,
		db:                   db,
	}, nil
}

func (a *app) Env() string                  { return a.env }
func (a *app) Port() string                 { return a.port }
func (a *app) SupaKey() string              { return a.supaKey }
func (a *app) SupaProjectReference() string { return a.supaProjectReference }
func (a *app) ClientKey() string            { return a.clientKey }
func (a *app) JwtKey() string               { return a.jwtKey }
func (a *app) Host() string                 { return a.host }
func (a *app) DB() *storage.DB              { return a.db }
