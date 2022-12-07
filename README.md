## Daemon

Daemon is an easy to start application as a system service.

### Installation
```shell
go get github.com/lawyzheng/daemon
```

### Quick Start
- #### Simple Example
```go
ctl, err := NewController(app, &service.Config{
		Name:             "My Service Name",       
		DisplayName:      "My Service Display Name", 
		Description:      "My Service Discription",           
	})

if err != nil {
    // handler error
	panic(err)
}


// Install Service
ctl.Install()

// Uninstall Service
ctl.Uninstall()

// Start Service
ctl.Start()

// Stop Service
ctl.Stop()

// Restart Service
ctl.Restart()

// Run Service
ctl.Run()
```