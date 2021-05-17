package app

type App interface {
	Run()
}

type app struct {
	config  *Config
	httpAPI *httpapi.API
	server  *http.Server
}

func (a *app) parseConfig(configFileName string) error {
	config, err := NewConfig(configFileName)
	if err != nil {
		return err
	}
	a.config = config
	return nil
}
