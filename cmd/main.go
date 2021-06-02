package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/webitel/sql_for_dialer/internal/handler"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello Windows!")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	http.HandleFunc("/members", handler.GetMembers)
	http.HandleFunc("/setMemberInfo", handler.SetMemberStat)
	log.Info().Str("Port", fmt.Sprintf("%v", cfg.Port)).Msg("Server started")
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
		return
	}
}
