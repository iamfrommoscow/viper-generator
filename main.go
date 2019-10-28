package main

import (
	"flag"
)

func main() {
	interactor := flag.Bool("interactor", true, "interactor")
	router := flag.Bool("router", true, "router")
	moduleName := flag.String("moduleName", "NeModule", "moduleName")
	path := flag.String("path", "/../", "path")

	flag.Parse()

	*path = *path + *moduleName + "/"

	createDirectory(*path)

	createViewFile(*path, *moduleName)
	if *interactor {
		createInteractorFile(*path, *moduleName)
	}
	createPresenterFile(*path, *moduleName, *interactor, *router)
	if *router {
		createRouterFile(*path, *moduleName)
	}
	createConfiguratorFile(*path, *moduleName, *interactor, *router)
	createProtocolsFile(*path, *moduleName, *interactor, *router)

}
