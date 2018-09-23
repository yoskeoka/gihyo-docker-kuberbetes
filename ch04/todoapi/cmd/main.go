package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gihyodocker/todoapi"
)

func main() {

	// (1) 必要な環境変数を格納した構造体を作成
	env, err := todoapi.CreateEnv()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	// (2) MySQL Masterへの接続するための構造体を作成
	masterDB, err := todoapi.CreateDbMap(env.MasterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.MasterURL)
		return
	}

	// (3) MySQL Slaveへの接続するための構造体を作成
	slaveDB, err := todoapi.CreateDbMap(env.SlaveURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.SlaveURL)
		return
	}

	mux := http.NewServeMux()

	// (4) ヘルスチェック用APIのハンドラを作成
	hc := func(w http.ResponseWriter, r *http.Request) {
		log.Println("[GET] /hc")
		w.Write([]byte("OK"))
	}

	// (5) TODO操作APIのハンドラを作成
	todoHandler := todoapi.NewTodoHandler(masterDB, slaveDB)

	// (6) ハンドラをAPIエンドポイントとして登録
	mux.Handle("/todo", todoHandler)
	mux.HandleFunc("/hc", hc)

	// (7) サーバのポートやハンドラを設定し、Listenを開始
	s := http.Server{
		Addr:    env.Bind,
		Handler: mux,
	}
	log.Printf("Listen HTTP Server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
