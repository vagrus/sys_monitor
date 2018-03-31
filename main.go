package main

import (
	"logger"
	"monitor/memory"
	"net/http"
	"websock"
)

var addr = "0.0.0.0:8080"

func main() {
	err := logger.Init(logger.LEVEL_DEBUG)
	if err != nil {
		panic(err)
	}
	defer logger.Stop()

	manager := websock.Manager{}
	manager.Start()

	defer manager.Stop()

	_, ch := memory.Start()

	go func() {
		for data := range ch {
			manager.Send(data)
		}
	}()

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		logger.Debug("handle new client")
		manager.NewClient(w, req)
	})
	logger.Error(http.ListenAndServe(addr, nil))

}
