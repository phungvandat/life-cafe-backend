package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/phungvandat/life-cafe-backend/endpoints"
	serviceHttp "github.com/phungvandat/life-cafe-backend/http"
	"github.com/phungvandat/life-cafe-backend/http/middlewares"
	"github.com/phungvandat/life-cafe-backend/service"
	authSvc "github.com/phungvandat/life-cafe-backend/service/auth"
	userSvc "github.com/phungvandat/life-cafe-backend/service/user"
	"github.com/phungvandat/life-cafe-backend/util/config"
	"github.com/phungvandat/life-cafe-backend/util/config/db/pg"
)

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

	// Setup addr
	port := "3000"
	if config.GetPortEnv() != "" {
		port = config.GetPortEnv()
	}

	httpAddr := fmt.Sprintf(":%v", port)

	// Setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// Setup service
	var (
		pgDB, closeDB = pg.New(config.GetPGDataSourceEnv())
		s             = service.Service{
			UserService: service.Compose(
				userSvc.NewPGService(pgDB, logger),
				userSvc.ValidationMiddleware(),
			).(userSvc.Service),
			AuthService: service.Compose(
				authSvc.NewPGService(pgDB, logger),
				authSvc.ValidationMiddleware(),
			).(authSvc.Service),
		}
	)
	defer closeDB()

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			middlewares.MakeHttpMiddleware(s),
			endpoints.MakeServerEndpoints(s),
			logger,
			os.Getenv("ENV") == "local",
		)
	}

	errs := make(chan error)
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-ch)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)

	}()

	logger.Log("exit", <-errs)
}
