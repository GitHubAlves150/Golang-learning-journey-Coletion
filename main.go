package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	App     AppConfig     `mapstructure:"app"`
	Message MessageConfig `mapstructure:"message"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type AppConfig struct {
	Mode  string `mapstructure:"mode"`
	Debug bool   `mapstructure:"debug"`
}

type MessageConfig struct {
	Welcome string `mapstructure:"welcome"`
	Health  string `mapstructure:"health"`
}

func loadConfig() (*Config, error) {
	// 1. Configurar o Viper
	v := viper.New()

	// Nome do arquivo de configuração (sem extensão)
	v.SetConfigName("config.default")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	// 2. Definir valores padrão (fallback)
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.name", "Default API")
	v.SetDefault("app.mode", "production")
	v.SetDefault("app.debug", false)
	v.SetDefault("message.welcome", "Bem-vindo!")
	v.SetDefault("message.health", "OK")

	// 3. Ler arquivo de configuração (se existir)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("⚠️ Arquivo de configuração não encontrado, usando padrões")
		} else {
			return nil, fmt.Errorf("erro ao ler arquivo: %w", err)
		}
	}

	// 4. Configurar variáveis de ambiente
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// Mapear variáveis de ambiente para chaves do Viper
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("app.mode", "APP_MODE")
	v.BindEnv("app.debug", "APP_DEBUG")
	v.BindEnv("message.welcome", "WELCOME_MESSAGE")

	// 5. Parsear para struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erro ao parsear config: %w", err)
	}

	return &config, nil
}

func setupServer(cfg *Config) *echo.Echo {
	e := echo.New()

	// Middleware para log das requisições
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Printf("📡 %s %s", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	})

	// Middleware de debug (se ativado)
	if cfg.App.Debug {
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				log.Println("🔧 Modo DEBUG ativado")
				c.Response().Header().Set("X-Debug-Mode", "true")
				return next(c)
			}
		})
	}

	// Rotas
	e.GET("/", welcomeHandler(cfg))
	e.GET("/health", healthHandler(cfg))
	e.GET("/config", configHandler(cfg))

	// Grupo de rotas admin (só disponível em debug)
	if cfg.App.Debug {
		admin := e.Group("/admin")
		admin.GET("/info", adminInfoHandler(cfg))
		admin.GET("/reload", reloadConfigHandler(cfg))
	}

	return e
}

// ============================================================
// Handlers
// ============================================================

func welcomeHandler(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": cfg.Message.Welcome,
			"server":  cfg.Server.Name,
			"mode":    cfg.App.Mode,
		})
	}
}

func healthHandler(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  "ok",
			"message": cfg.Message.Health,
		})
	}
}

func configHandler(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, cfg)
	}
}

func adminInfoHandler(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"config":         cfg,
			"environment":    os.Environ(),
			"viper_all_keys": viper.AllKeys(),
		})
	}
}

func reloadConfigHandler(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Recarregar configuração (simulação)
		return c.JSON(200, map[string]string{
			"message": "Recarga de config seria feita aqui",
		})
	}
}

// ============================================================
// Main
// ============================================================

func main() {
	log.Println("🚀 Iniciando aplicação...")

	// Carregar configuração
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("❌ Erro fatal ao carregar configuração: %v", err)
	}

	log.Printf("📋 Configuração carregada:")
	log.Printf("   Servidor: %s (porta %d)", cfg.Server.Name, cfg.Server.Port)
	log.Printf("   Modo: %s, Debug: %t", cfg.App.Mode, cfg.App.Debug)

	// Configurar e iniciar servidor
	server := setupServer(cfg)

	// Iniciar
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("🌐 Servidor iniciado em http://localhost%s", addr)
	log.Println("📡 Endpoints disponíveis:")
	log.Println("   GET /           - Página inicial")
	log.Println("   GET /health     - Health check")
	log.Println("   GET /config     - Ver configurações atuais")

	if cfg.App.Debug {
		log.Println("   GET /admin/info - Informações de debug")
		log.Println("   GET /admin/reload - Recarregar config")
	}

	if err := server.Start(addr); err != nil {
		log.Fatalf("❌ Servidor falhou: %v", err)
	}
}
