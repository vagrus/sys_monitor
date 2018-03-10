package memory

import (
	"monitor"
	"time"
	"os/exec"
	"strings"
	"strconv"
	"logger"
)

const NAME = base.RAM

var timeout = 1 * time.Second

type Data struct {
	Total int32	`json:"total"`
	Used int32	`json:"used"`
	Free int32	`json:"free"`
}

func Start() (string, <-chan base.MonitoringData) {
	ch := make(chan base.MonitoringData, 100)

	go func() {
		for {
			d := base.MonitoringData{NAME, time.Now(), getData()}
			ch <- d
			time.Sleep(timeout)
		}

		defer close(ch)
	}()

	return NAME, ch
}

func getData() Data {
	out, err := exec.Command("bash", "-c", "free -m | sed -n 2p | sed 's/  */ /g'").Output()

	outArr := strings.Split(string(out), " ")

	if err != nil {
		logger.Error(NAME, err.Error())
	}

	total, _ := strconv.ParseInt(outArr[1], 10, 32)
	used, _ := strconv.ParseInt(outArr[2], 10, 32)
	free, _ := strconv.ParseInt(outArr[3], 10, 32)

	return Data{int32(total), int32(used), int32(free)}
}
