package main

import "fmt"

const protocolsTemplate = `
import Foundation

protocol %sViewProtocol: class{
    var presenter: %sPresenterProtocol! { get set }
   
}

protocol %sConfiguratorProtocol: class {
    func configure(with viewController: %sViewProtocol)
}

protocol %sPresenterProtocol: class {
	var view: %sViewProtocol! {get set}
	
	%s
   
}

%s

`

const protocolsInteractorVar = `
	var interactor: %sInteractorProtocol! {get set}`

const protocolsRouterVar = `
	var router: %sRouterProtocol! {get set}`

const protocolsInteractor = `

protocol %sInteractorProtocol: class {
	var presenter: %sPresenterProtocol! { get set }
}`

const protocolsRouter = `
	
protocol %sRouterProtocol: class {

}`

func createProtocolsFile(path string, name string, interactor bool, router bool) {
	var extraPresenterVars = ``
	var extraProtocols = ``

	if interactor {
		extraPresenterVars += fmt.Sprintf(protocolsInteractorVar, name)
		extraProtocols += fmt.Sprintf(protocolsInteractor, name, name)
	}

	if router {
		extraPresenterVars += fmt.Sprintf(protocolsRouterVar, name)
		extraProtocols += fmt.Sprintf(protocolsRouter, name)
	}

	protocolsString := fmt.Sprintf(protocolsTemplate, name, name, name, name, name, name, extraPresenterVars, extraProtocols)
	createFile(path, fmt.Sprintf(`%sProtocols`, name), protocolsString)
}
