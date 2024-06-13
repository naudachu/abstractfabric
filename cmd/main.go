package main

import (
	"abstractfabric/fabric"
	"abstractfabric/fabricmethod"
	"abstractfabric/pluggin"
	"fmt"
)

func main() {

	testAbstractFabricMethod()
	fmt.Printf("\n\n---\n\n")

	testFabricMethod()
	fmt.Printf("\n\n---\n\n")

	realAFactorytest()
}

func testAbstractFabricMethod() {

	iosF := fabric.NewIosFactory().SetPrefix("xios")
	aApp := iosF.NewApp(1, "com.dev.app")
	data := aApp.GetApp()

	fmt.Printf(
		"App created.\nLink:\t%s\nID: %d\tName: %s\tBundle: %s\nType: %T\n\n", aApp.GetLink(), data.ID, data.Code, data.Bundle, aApp)

	aF := fabric.NewAndroidFactory().SetPrefix("xand")
	bApp := aF.NewApp(2, "com.dev.app")
	data = bApp.GetApp()

	fmt.Printf("App created.\nLink:\t%s\nID: %d\tName: %s\tBundle: %s\nType: %T\n\n", bApp.GetLink(), data.ID, data.Code, data.Bundle, bApp)
}

func testFabricMethod() {

	iosF := fabricmethod.NewIosFactory().SetPrefix("xios")
	aApp := iosF.NewApp(1, "com.dev.app")
	data := aApp.GetApp()

	fmt.Printf(
		"App created.\nLink:\t%s\nID: %d\tName: %s\tBundle: %s\nType: %T\n\n", aApp.GetLink(), data.ID, data.Code, data.Bundle, aApp)

	aF := fabricmethod.NewAndroidFactory().SetPrefix("xand")
	bApp := aF.NewApp(2, "com.dev.app")
	data = bApp.GetApp()

	fmt.Printf("App created.\nLink:\t%s\nID: %d\tName: %s\tBundle: %s\nType: %T\n\n", bApp.GetLink(), data.ID, data.Code, data.Bundle, bApp)
}

func realAFactorytest() {
	aFactory := pluggin.NewSomeFactory().SetPrefix("2")
	app := aFactory.NewApp(3, "hello, beatches")
	data := app.GetApp()

	fmt.Printf("App created.\nLink:\t%s\nID: %d\tName: %s\tBundle: %s\nType: %T\n\n", app.GetLink(), data.ID, data.Code, data.Bundle, app)
}
