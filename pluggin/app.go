package pluggin

import (
	"abstractfabric/fabric"
	"fmt"
	"strconv"
)

type SomeFactory struct {
	prefix int
}

func NewSomeFactory() fabric.AbstractFactory {
	return &SomeFactory{}
}

func (sf *SomeFactory) SetPrefix(prefix string) fabric.AbstractFactory {
	i, _ := strconv.Atoi(prefix)
	//fixme ofc нужна проверка;
	sf.prefix = i
	return sf
}

func (sf *SomeFactory) NewApp(id int64, bundle string) fabric.IApp {
	return &SomeApp{
		Code: fmt.Sprintf("#%d-%d", sf.prefix, id),
	}
}

type SomeApp struct {
	Code string
}

func (sa *SomeApp) GetApp() *fabric.App {
	return &fabric.App{}
}

func (sa *SomeApp) GetLink() string {
	return "https://some.store/" + sa.Code
}
