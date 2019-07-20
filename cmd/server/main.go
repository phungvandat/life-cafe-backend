package main

import (
	"context"
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
	categorySvc "github.com/phungvandat/life-cafe-backend/service/category"
	orderSvc "github.com/phungvandat/life-cafe-backend/service/order"
	productSvc "github.com/phungvandat/life-cafe-backend/service/product"
	uploadSvc "github.com/phungvandat/life-cafe-backend/service/upload"
	userSvc "github.com/phungvandat/life-cafe-backend/service/user"
	"github.com/phungvandat/life-cafe-backend/util/config"
	"github.com/phungvandat/life-cafe-backend/util/config/db/pg"
	"github.com/phungvandat/life-cafe-backend/util/helper"
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
		spRollback    = helper.NewSagasService()

		userService = service.Compose(
			userSvc.NewPGService(pgDB, spRollback),
			userSvc.ValidationMiddleware(),
		).(userSvc.Service)

		authService = service.Compose(
			authSvc.NewPGService(pgDB),
			authSvc.ValidationMiddleware(),
		).(authSvc.Service)

		uploadService = service.Compose(
			uploadSvc.NewPGService(),
			uploadSvc.ValidationMiddleware(),
		).(uploadSvc.Service)

		categoryService = service.Compose(
			categorySvc.NewPGService(pgDB),
			categorySvc.ValidationMiddleware(),
		).(categorySvc.Service)

		productService = service.Compose(
			productSvc.NewPGService(pgDB, categoryService, spRollback),
			productSvc.ValidationMiddleware(),
		).(productSvc.Service)

		orderService = service.Compose(
			orderSvc.NewPGService(pgDB, userService, productService, spRollback),
			orderSvc.ValidationMiddleware(),
		).(orderSvc.Service)

		s = service.Service{
			UserService:     userService,
			AuthService:     authService,
			UploadService:   uploadService,
			CategoryService: categoryService,
			ProductService:  productService,
			OrderService:    orderService,
		}
	)
	defer closeDB()

	// Create master
	s.UserService.CreateMaster(context.Background())

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			middlewares.MakeHTTPpMiddleware(s),
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
