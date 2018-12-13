package main

var (
	appName = "eventor"
	version = "2.6.0"
	log     = zaplog.Prefix(appName)
	cfgFile = "config.json"
	cfg     = defaultConfig()
)

func main() {
	app.Prepare(app.PrepareOption{
		Name:        appName,
		Version:     version,
		CfgFile:     cfgFile,
		WriteConfig: true,
	}, &cfg)
	// init redis
	redispool.MustCNewClients(cfg.Redis.To0ptions())
	log.Info("init-redis")
}
