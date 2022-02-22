## Daemon

Daemon is an easy to start application as a system service.


### Quick Start
- #### Simple Example
```go
con, err := NewDaemonController(app, &service.Config{
		Name:             "My Service Name",       
		DisplayName:      "My Service Display Name", 
		Description:      "My Service Discription",           
	})

if err != nil {
    // handler error
    ...
}


// Install Service
con.Install()

// Uninstall Service
con.Uninstall()

// Start Service
con.Start()

// Stop Service
con.Stop()

// Restart Service
con.Restart()

// Run Service
con.Run()
```