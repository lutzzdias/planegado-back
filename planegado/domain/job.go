package domain

import log "github.com/sirupsen/logrus"

type Job struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (job Job) CalcRes() {
	log.Info("Calculando resultado...")
}
