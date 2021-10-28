package main

import (
	_ "database/sql"
	"flag"
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
	var port int
	flag.IntVar(&port, "port", 0, "Port number")
	flag.Parse()
	if err != nil {
		return
	}
	fullexecpath, err := os.Executable()
	if err != nil {
		return
	}
	log.Info().Msg(fullexecpath)
	if port == 0 {
		cfg, err := LoadConfig()
		if err != nil {
			itrlog.Fatal(err.Error())
			log.Fatal().Msg(err.Error())
		}
		port = cfg.Port
	}

	http.HandleFunc("/members", handler.GetMembers)
	http.HandleFunc("/setMemberInfo", handler.SetMemberStat)
	http.HandleFunc("/version", handler.GetCurrentVersion)
	http.HandleFunc("/logs", handler.GetLogs)
	log.Info().Str("Port", fmt.Sprintf("%v", port)).Msg("Server started")
	itrlog.Info(fmt.Sprintf("Server started. Port: %v", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
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
