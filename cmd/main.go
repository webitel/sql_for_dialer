package main

import (
	_ "database/sql"
	"fmt"
	"github.com/itrepablik/itrlog"
	_ "github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/webitel/sql_for_dialer/internal/handler"
	"net/http"
	"os"
)

func main() {
	err := setLoggers()
	if err != nil {
		return
	}
	fullexecpath, err := os.Executable()
	if err != nil {
		return
	}
	log.Info().Msg(fullexecpath)
	cfg, err := LoadConfig()
	if err != nil {
		itrlog.Fatal(err.Error())
		log.Fatal().Msg(err.Error())
	}
	http.HandleFunc("/members", handler.GetMembers)
	http.HandleFunc("/setMemberInfo", handler.SetMemberStat)
	http.HandleFunc("/version", handler.GetVersion)
	log.Info().Str("Port", fmt.Sprintf("%v", cfg.Port)).Msg("Server started")
	itrlog.Info(fmt.Sprintf("Server started. Port: %v", cfg.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
		return
	}
}

func setLoggers() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	itrlog.SetLogInit(50, 90, "logs", "run_log_")
	return nil
}
