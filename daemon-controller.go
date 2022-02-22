package daemon

import (
	"fmt"

	"github.com/kardianos/service"
)

type DaemonController struct {
	Service service.Service
}

// NewDaemonController 实例化守护进程控制器
func NewDaemonController(app Application, config *service.Config) (*DaemonController, error) {
	d := NewDaemon(app, config)
	if service.Platform() == "linux-systemd" {
		d.config.Option = service.KeyValue{
			"LimitNOFILE": 40960,
		}
	}

	s, err := service.New(d, d.config)
	if err != nil {
		return nil, err
	}
	return &DaemonController{
		Service: s,
	}, nil
}

// Install 安装守护进程
func (d *DaemonController) Install() {
	err := d.Service.Install()
	if err != nil {
		fmt.Println("Install Error -> ", err)
		return
	}
	fmt.Println("Installed Successfully!")
}

// Uninstall 卸载守护进程
func (d *DaemonController) Uninstall() {
	d.Service.Stop()
	err := d.Service.Uninstall()
	if err != nil {
		fmt.Println("Uninstall Error -> ", err)
		return
	}
	fmt.Println("Uninstalled Successfully!")
}

// Start 开始守护进程
func (d *DaemonController) Start() {
	err := d.Service.Start()
	if err != nil {
		fmt.Println("Start Error -> ", err)
		return
	}
	fmt.Println("Started Successfully!")
}

// Stop 停止守护进程
func (d *DaemonController) Stop() {
	err := d.Service.Stop()
	if err != nil {
		fmt.Println("Stop Error -> ", err)
		return
	}
	fmt.Println("Stopped Successfully!")
}

// restart 重启守护进程
func (d *DaemonController) Restart() {
	err := d.Service.Restart()
	if err != nil {
		fmt.Println("Restart Error -> ", err)
		return
	}
	fmt.Println("Restarted Successfully!")
}

// Run 命令行运行守护进程
func (d *DaemonController) Run() {
	err := d.Service.Run()
	if err != nil {
		fmt.Println("Run Error -> ", err)
		return
	}
}
