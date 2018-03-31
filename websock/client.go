package websock

import (
	"github.com/gorilla/websocket"
	"logger"
	"monitor"
	"net/http"
)

type Client struct {
	isRunning bool
	// individual client chan
	ch     chan base.MonitoringData
	wsConn *websocket.Conn
}

func (this *Client) Run(w http.ResponseWriter, req *http.Request) {
	logger.Debug("client running")
	this.ch = make(chan base.MonitoringData, 20)
	this.createWS(w, req)
	this.isRunning = true

	go this.sendToWs()
}

func (this *Client) createWS(w http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// allow all connections by default
			return true
		},
	}

	var err error
	this.wsConn, err = upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("create ws", err)
		return
	}
}

func (this *Client) Stop() {
	logger.Debug("closing client")

	this.isRunning = false
	close(this.ch)
	this.wsConn.Close()
}

func (this *Client) Send(data base.MonitoringData) {
	if this.isRunning {
		this.ch <- data
	}
}

func (this *Client) sendToWs() {
	logger.Debug("client is running:", this.isRunning)
	var data base.MonitoringData
	for this.isRunning {
		data = <-this.ch
		logger.Debug("data received")

		err := this.wsConn.WriteJSON(data)
		if err != nil {
			logger.Error("ws write", err)
		}
		logger.Debug("data sended")
	}
}
