package websock

import (
	"logger"
	"monitor"
	"net/http"
	"sync"
)

type Manager struct {
	// main agg chan
	ch          chan base.MonitoringData
	clients     []Client
	clientsLock sync.RWMutex
	isRunning   bool
}

func (this *Manager) NewClient(w http.ResponseWriter, req *http.Request) {
	if this.isRunning != true {
		logger.Error("manager", "trying to accept new client, but manager is stopped")
		return
	}

	client := Client{}
	client.Run(w, req)

	this.clientsLock.Lock()
	defer this.clientsLock.Unlock()
	this.clients = append(this.clients, client)
}

func (this *Manager) Start() {
	this.ch = make(chan base.MonitoringData, 100)
	this.isRunning = true

	go func() {
		var data base.MonitoringData
		var ok = true
		for this.isRunning && ok {
			data, ok = <-this.ch

			if ok {
				this.sendToClients(data)
			}
		}

		logger.Debug("manager", "stopping base goro")
	}()
}

func (this *Manager) Stop() {
	logger.Debug("stopping manager")

	this.isRunning = false

	this.clientsLock.Lock()
	defer this.clientsLock.Unlock()

	for _, client := range this.clients {
		client.Stop()
	}

	this.clients = []Client{}
	close(this.ch)
}

func (this *Manager) sendToClients(data base.MonitoringData) {
	this.clientsLock.RLock()
	defer this.clientsLock.RUnlock()

	for _, client := range this.clients {
		client.Send(data)
	}
}

func (this *Manager) Send(data base.MonitoringData) {
	if this.isRunning {
		this.ch <- data
	}
}
