package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/skripov-ds-ai/highload_course/internal/config"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/dialog"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/friend"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/login"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/post"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/user"
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func NewModule() fx.Option {
	return fx.Module(
		"httpsrv",
		fx.Provide(
			// handlers
			dialog.NewHandler,
			friend.NewHandler,
			login.NewHandler,
			post.NewHandler,
			user.NewHandler,

			config.NewHttpConfig,

			fx.Annotate(
				NewHandler,
				fx.As(new(generated.ServerInterface)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, cfg *config.HttpConfig, srv generated.ServerInterface, log *zap.Logger) {
				//log.Info("!")
				server := &http.Server{
					Addr:    fmt.Sprintf(":%d", cfg.Port),
					Handler: service(srv),
				}

				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
								log.Error("server error: ", zap.Error(err))
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						err := server.Shutdown(ctx)
						return err
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("httpsrv")
		}),
	)
}

func service(si generated.ServerInterface) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	opts := generated.ChiServerOptions{
		BaseRouter: r,
	}

	return generated.HandlerWithOptions(si, opts)
}

//func main() {
//	// config
//	cfg := config.Config{}
//
//	srv := &v1.Handler{}
//
//	// http server
//	server := &http.Server{
//		Addr:    fmt.Sprintf(":%d", cfg.Http.Port),
//		Handler: service(srv),
//	}
//
//	// Server run context
//	serverCtx, serverStopCtx := context.WithCancel(context.Background())
//
//	// Listen for syscall signals for process to interrupt/quit
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//	go func() {
//		<-sig
//
//		// Shutdown signal with grace period of 30 seconds
//		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
//		defer cancel()
//
//		go func() {
//			<-shutdownCtx.Done()
//			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
//				log.Fatal("graceful shutdown timed out.. forcing exit.")
//			}
//		}()
//
//		// Trigger graceful shutdown
//		err := server.Shutdown(shutdownCtx)
//		if err != nil {
//			log.Fatal(err)
//		}
//		serverStopCtx()
//	}()
//
//	// Run the server
//	err := server.ListenAndServe()
//	if err != nil && !errors.Is(err, http.ErrServerClosed) {
//		log.Fatal(err)
//	}
//
//	// Wait for server context to be stopped
//	<-serverCtx.Done()
//}
//
//func service(si generated.ServerInterface) http.Handler {
//	r := chi.NewRouter()
//	r.Use(middleware.RequestID)
//	r.Use(middleware.Logger)
//	r.Use(middleware.Recoverer)
//
//	opts := generated.ChiServerOptions{
//		BaseRouter: r,
//	}
//
//	return generated.HandlerWithOptions(si, opts)
//}
