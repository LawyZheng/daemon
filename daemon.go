package daemon

import (
	"os"

	"github.com/kardianos/service"
)

type Service interface {
	Run() error
	HandleError(error)
}

type Daemon struct {
	svc  Service
	errs chan error
}

func NewDaemon(svc Service) *Daemon {
	return &Daemon{
		errs: make(chan error, 100),
		svc:  svc,
	}
}

func (d *Daemon) Start(s service.Service) error {
	go d.run()
	return nil
}

func (d *Daemon) run() {
	// 运行逻辑
	err := d.svc.Run()
	if err != nil {
		d.svc.HandleError(err)
		return
	}
}

func (d *Daemon) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}
