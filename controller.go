package daemon

import (
	"github.com/kardianos/service"
)

type Controller struct {
	service.Service
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
	if service.Platform() == "linux-systemd" {
		if conf.Option == nil {
			conf.Option = make(service.KeyValue)
		}
		conf.Option["LimitNOFILE"] = 40960
	}

	d := NewDaemon(svc)
	s, err := service.New(d, conf)
	if err != nil {
		return nil, err
	}
	return &Controller{
		Service: s,
	}, nil
}

/**
 * @name: Install
 * @description: 守护进程安装
 * @return {error}
 */
// func (d *Controller) Install() error {
// 	return d.service.Install()
// }

/**
 * @name: Uninstall
 * @description: 守护进程卸载
 * @return {error}
 */
func (d *Controller) Uninstall() error {
	d.Stop()
	return d.Service.Uninstall()
}

/**
 * @name: Start
 * @description: 启动守护进程
 * @return {error}
 */
// func (d *Controller) Start() error {
// 	return d.service.Start()
// }

/**
 * @name: Stop
 * @description: 停止守护进程
 * @return {error}
 */
// func (d *Controller) Stop() error {
// 	return d.service.Stop()
// }

/**
 * @name: Restart
 * @description: 重启守护进程
 * @return {error}
 */
// func (d *Controller) Restart() error {
// 	return d.service.Restart()
// }

/**
 * @name: Run
 * @description: 运行守护进程
 * @return {error}
 */
// func (d *Controller) Run() error {
// 	return d.service.Run()
// }
