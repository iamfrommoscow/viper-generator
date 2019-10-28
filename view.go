package main

import "fmt"

const viewTemplate = `
import UIKit

class %sViewController: UIViewController {

    
    var presenter: %sPresenterProtocol!
    let configurator: %sConfiguratorProtocol = %sConfigurator()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        configurator.configure(with: self)
        
    }
     

}


`

func createViewFile(path string, name string) {
	viewString := fmt.Sprintf(viewTemplate, name, name, name, name)
	createFile(path, fmt.Sprintf(`%sViewController`, name), viewString)
}
