package daemon

import (
	"github.com/kardianos/service"
)

type Controller struct {
	service service.Service
}

/**
 * @name: NewController
 * @description: 守护进程控制器
 * @param {*service.Config} conf->服务配置
 * @param {Service} svc->服务接口
 * @return {*Controller} 守护进程控制器
 * @return {error} 错误
 */
func NewController(conf *service.Config, svc Service) (*Controller, error) {
	d := NewDaemon(conf, svc)
	if service.Platform() == "linux-systemd" {
		d.config.Option = service.KeyValue{
			"LimitNOFILE": 40960,
		}
	}

	s, err := service.New(d, d.config)
	if err != nil {
		return nil, err
	}
	return &Controller{
		service: s,
	}, nil
}

/**
 * @name: Install
 * @description: 守护进程安装
 * @return {error}
 */
func (d *Controller) Install() error {
	return d.service.Install()
}

/**
 * @name: Uninstall
 * @description: 守护进程卸载
 * @return {error}
 */
func (d *Controller) Uninstall() error {
	d.service.Stop()
	return d.service.Uninstall()
}

/**
 * @name: Start
 * @description: 启动守护进程
 * @return {error}
 */
func (d *Controller) Start() error {
	return d.service.Start()
}

/**
 * @name: Stop
 * @description: 停止守护进程
 * @return {error}
 */
func (d *Controller) Stop() error {
	return d.service.Stop()
}

/**
 * @name: Restart
 * @description: 重启守护进程
 * @return {error}
 */
func (d *Controller) Restart() error {
	return d.service.Restart()
}

/**
 * @name: Run
 * @description: 运行守护进程
 * @return {error}
 */
func (d *Controller) Run() error {
	return d.service.Run()
}
