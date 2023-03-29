package main

import (
	"github.com/egorgasay/grpc-storage/internal/cli/config"
	"github.com/egorgasay/grpc-storage/internal/cli/handler"
	"github.com/egorgasay/grpc-storage/internal/cli/usecase"
	"github.com/egorgasay/grpc-storage/pkg/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httplog"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	cfg := config.New()

	logic := usecase.New(cfg)

	e := echo.New()
	lg := httplog.NewLogger("cli", httplog.Options{
		Concise: true,
	})
	e.Use(echo.WrapMiddleware(httplog.RequestLogger(lg)))
	e.Use(echo.WrapMiddleware(middleware.Recoverer))
	h := handler.New(cfg, logic, logger.New(lg))
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/html/*.html")),
	}
	e.Renderer = t
	h.PublicRoutes(e)
	e.Static("/static", "static")

	//router.Use(gzip.Gzip(gzip.BestSpeed))
	go func() {
		lg.Info().Msg("Stating CLI-Server: " + cfg.Host)
		err := http.ListenAndServe(cfg.Host, e)
		if err != nil {
			lg.Error().Msg(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	lg.Info().Msg("Shutdown CLI-Server ...")
}
