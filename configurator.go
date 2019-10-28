package main

import "fmt"

const configuratorTemplate = `
import Foundation

class %sConfigurator: %sConfiguratorProtocol {
    func configure(with viewController: %sViewProtocol) {
		let presenter = %sPresenter(view: viewController)
		viewController.presenter = presenter
		
		%s
        
    }
}`

const configuratorRouterTemplate = `
	let router = %sRouter()
	presenter.router = router`

const configuratorInteractorTemplate = `
	let interactor = %sInteractor(presenter: presenter)
	presenter.interactor = interactor`

func createConfiguratorFile(path string, name string, interactor bool, router bool) {
	var extraClassesString = ``
	if interactor {
		extraClassesString += fmt.Sprintf(configuratorInteractorTemplate, name)
	}
	if router {
		extraClassesString += fmt.Sprintf(configuratorRouterTemplate, name)
	}
	configuratorString := fmt.Sprintf(configuratorTemplate, name, name, name, name, extraClassesString)
	createFile(path, fmt.Sprintf(`%sConfigurator`, name), configuratorString)
}
