package main

import "fmt"

const interactorTemplate = `
import Foundation

class %sInteractor: %sInteractorProtocol {
    var presenter: %sPresenterProtocol!
    
    required init(presenter: %sPresenterProtocol) {
       self.presenter = presenter
   	}    
    
}

`

func createInteractorFile(path string, name string) {
	interactorString := fmt.Sprintf(viewTemplate, name, name, name, name)
	createFile(path, fmt.Sprintf(`%sInteractor`, name), interactorString)
}
