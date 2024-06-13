package fabricmethod

import "fmt"

type App struct {
	ID     int64
	Code   string
	Bundle string
}

type iapp interface {
	GetApp() *App
	GetLink() string
}

type iosApp struct {
	app        App
	appStoreID string
}

func (ia *iosApp) GetApp() *App {
	return &ia.app
}

func (ia *iosApp) GetLink() string {
	return "https://apps.apple.com/app/" + ia.appStoreID
}

type IosFactory struct {
	prefix string
}

func NewIosFactory() *IosFactory {
	return &IosFactory{}
}

func (af *IosFactory) SetPrefix(prefix string) *IosFactory {
	af.prefix = prefix
	return af
}

func (iof *IosFactory) NewApp(id int64, bundle string) iapp {
	return &iosApp{
		app: App{
			ID:     id,
			Code:   fmt.Sprintf("%s-%d", iof.prefix, id),
			Bundle: bundle,
		},
	}
}

type androidApp struct {
	app App
}

func (aa *androidApp) GetLink() string {
	return "https://play.google.com/store/apps/details?id=" + aa.app.Bundle
}

func (ia *androidApp) GetApp() *App {
	return &ia.app
}

type AndroidFactory struct {
	prefix string
}

func NewAndroidFactory() *AndroidFactory {
	return &AndroidFactory{}
}

func (af *AndroidFactory) SetPrefix(prefix string) *AndroidFactory {
	af.prefix = prefix
	return af
}

func (af *AndroidFactory) NewApp(id int64, bundle string) iapp {
	return &androidApp{
		app: App{
			ID:     id,
			Code:   fmt.Sprintf("%s-%d", af.prefix, id),
			Bundle: bundle,
		},
	}
}
