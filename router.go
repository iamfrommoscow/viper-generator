package main

import "fmt"

const routerTemplate = `
import Foundation
import UIKit

class %sRouter: %sRouterProtocol {
    
    
}`

func createRouterFile(path string, name string) {
	routerString := fmt.Sprintf(routerTemplate, name, name)
	createFile(path, fmt.Sprintf(`%sRouter`, name), routerString)

}
