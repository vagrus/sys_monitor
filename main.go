package main

import (
	"monitor"
	"monitor/memory"
	"logger"

	"encoding/json"
)


func main() {
	err := logger.Init(logger.LEVEL_DEBUG)
	if err != nil {
		panic(err)
	}

	defer logger.Stop()

	name, ch := memory.Start()

	logger.Debug(name)

	worker(ch)
}

func worker(ch <-chan base.MonitoringData) {
	for m := range ch {
		j, err := json.Marshal(m)
		println(string(j))
		if err != nil {
			println(err.Error())
		}
	}
}
