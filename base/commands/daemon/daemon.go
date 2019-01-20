package daemon

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/kardianos/service"
)

type Daemon struct {
	config *service.Config
	errs   chan error
}

func NewDaemon() *Daemon {

	config := &service.Config{
		Name:             "box-qrcode",//服务显示名称
		DisplayName:      "box-qrcode service",//服务名称
		Description:      "生成箱码PDF", //服务描述
		Arguments:        os.Args[1:],
	}

	return &Daemon{
		config: config,
		errs:   make(chan error, 100),
	}
}

func (d *Daemon) Config() *service.Config {
	return d.config
}
func (d *Daemon) Start(s service.Service) error {

	go d.Run()
	beego.Info("Service Start!")
	return nil
}

func (d *Daemon) Run() {
	beego.Run()
}

func (d *Daemon) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

func Install() {
	d := NewDaemon()
	//d.config.Arguments = os.Args[3:]

	s, err := service.New(d, d.config)

	if err != nil {
		beego.Error("Create service error => ", err)
		os.Exit(1)
	}
	err = s.Install()
	if err != nil {
		beego.Error("Install service error:", err)
		os.Exit(1)
	} else {
		beego.Info("Service installed!")
	}

	os.Exit(0)
}

func Uninstall() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		beego.Error("Create service error => ", err)
		os.Exit(1)
	}
	err = s.Uninstall()
	if err != nil {
		beego.Error("Install service error:", err)
		os.Exit(1)
	} else {
		beego.Info("Service uninstalled!")
	}
	os.Exit(0)
}

func Restart() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		beego.Error("Restart service error => ", err)
		os.Exit(1)
	}
	err = s.Restart()
	if err != nil {
		beego.Error("Restart service error:", err)
		os.Exit(1)
	} else {
		beego.Info("Service Restart!")
	}
	os.Exit(0)
}

func Stop() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		beego.Error("Stop service error => ", err)
		os.Exit(1)
	}
	err = s.Stop()
	if err != nil {
		beego.Error("Stop service error:", err)
		os.Exit(1)
	} else {
		beego.Info("Service Stop!")
	}
	os.Exit(0)
}

func Status() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		beego.Error("Get service Status error => ", err)
		os.Exit(1)
	}
	status, err := s.Status()
	if err != nil {
		beego.Error("Get service Status error:", err)
		os.Exit(1)
	} else {
		beego.Info(status)
	}
	os.Exit(0)
}


