package daemon

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

type Application interface {
	Run() error
}

type Daemon struct {
	config *service.Config
	app    Application
	errs   chan error
}

func NewDaemon(app Application, config *service.Config) *Daemon {
	return &Daemon{
		config: config,
		app:    app,
		errs:   make(chan error, 100),
	}
}

func (d *Daemon) Start(s service.Service) error {
	go d.run()
	return nil
}

func (d *Daemon) run() {
	// 运行逻辑
	err := d.app.Run()
	if err != nil {
		fmt.Println("Start Server Error -> ", err)
		os.Exit(1)
		return
	}
}

func (d *Daemon) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}
