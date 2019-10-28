package main

import "fmt"

const presenterTemplate = `
import Foundation

class %sPresenter: %sPresenterProtocol {
    required init(view: %sViewProtocol) {
       self.view = view
    }
    
    
    var view: %sViewProtocol!
	
	%s
    
    
}

`

const presenterInteractorTemplate = `
	var interactor: %sInteractorProtocol!`

const presenterRouterTemplate = `
	var router: %sRouterProtocol!`

func createPresenterFile(path string, name string, interactor bool, router bool) {
	var extraClassesString = ``
	if interactor {
		extraClassesString += fmt.Sprintf(presenterInteractorTemplate, name)
	}
	if router {
		extraClassesString += fmt.Sprintf(presenterRouterTemplate, name)
	}
	presenterString := fmt.Sprintf(presenterTemplate, name, name, name, name, extraClassesString)
	createFile(path, fmt.Sprintf(`%sPresenter`, name), presenterString)
}
