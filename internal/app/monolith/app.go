package monolith

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/skripov-ds-ai/highload_course/internal/config"
	v1 "github.com/skripov-ds-ai/highload_course/internal/controller/http/v1"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/login"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/user"
	"github.com/skripov-ds-ai/highload_course/internal/db/postgres"
	user_repository "github.com/skripov-ds-ai/highload_course/internal/db/repository/user"
	auth_service "github.com/skripov-ds-ai/highload_course/internal/service/auth"
	user_service "github.com/skripov-ds-ai/highload_course/internal/service/user"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		CreateApp(),
	)
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewProduction,
			// config
			config.NewConfig,
			config.NewDBConfig,

			// TODO: add shutdowner for db!
			// postgres
			postgres.NewDB,
			// repos
			fx.Annotate(
				user_repository.NewRepository,
				fx.As(new(user_repository.Repository)),
				fx.As(new(auth_service.Repository)),
				fx.As(new(user_service.Repository)),
			),
			// service
			fx.Annotate(
				auth_service.NewAuthService,
				fx.As(new(auth_service.AuthService)),
				fx.As(new(login.AuthService)),
			),

			fx.Annotate(
				user_service.NewUserService,
				fx.As(new(user_service.UserService)),
				fx.As(new(user.UserService)),
			),
		),
		// server
		v1.NewModule(),
		// logger
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		// supplements
		fx.Invoke(
			// ....
			func(diGraph fx.DotGraph, log *zap.Logger) {
				// ....
				log.Info(string(diGraph))
				// ....
			},
		),
	)
}

//	return fx.Options(
//		// fx.Supply(),
//		fx.Provide(
//			fx.Annotate(
//				v1.NewHandler,
//				fx.As(new(generated.ServerInterface)),
//			),
//			// TODO: move to interface with Annotate?
//			dialog.NewHandler,
//			friend.NewHandler,
//			login.NewHandler,
//			post.NewHandler,
//			user.NewHandler,
//		),
//		fx.Invoke(
//			// ....
//			func(diGraph fx.DotGraph) {
//				// ....
//				log.Println(diGraph)
//				// ....
//			},
//		),
//	)
//}

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
