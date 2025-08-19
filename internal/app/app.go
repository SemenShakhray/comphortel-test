package app

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"comphortel-test/internal/config"
	"comphortel-test/internal/delivery/handlers"
	"comphortel-test/internal/delivery/router"
	"comphortel-test/internal/repository"
	"comphortel-test/internal/service"
	"comphortel-test/utils/logger"
	"comphortel-test/utils/logger/sl"
)

func RunServer() {
	cfg := config.MustLoad()

	log := logger.SetupLogger()

	repo, err := repository.NewRepository(cfg, log)
	if err != nil {
		log.Error("failed to create repository", sl.Err(err))

		os.Exit(1)
	}
	defer repo.Close()

	service := service.NewService(repo, log)

	handler := handlers.NewHandler(service)

	router := router.NewRouter(&handler)

	srv := &http.Server{
		Addr:         net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.Port)),
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Info("starting server", slog.String("address", srv.Addr))

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("failed to start server", sl.Err(err))
			os.Exit(1)
		}
	}()

	sig := <-sigint
	log.Info("received signal", slog.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Info("failed to stop server", sl.Err(err))
	}
}
