package app

import "net/http"

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
func (a *app) runHTTPAPI() error {
	if err := a.server.ListenAndServe(); err != http.ErrServerClosed() {
		return err
	}
	return nil

}
func (a *app) Run() error {
	return a.runHTTPAPI()

}
func New(configFileName string) (App, error) {
	a := &app{}
	if err := a.parseConfig(configFileName); err != nil {
		return nil, err
	}
	return a, nil
}
