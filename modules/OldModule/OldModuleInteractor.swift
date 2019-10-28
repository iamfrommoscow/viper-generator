
import Foundation

class OldModuleInteractor: OldModuleInteractorProtocol {
    var presenter: OldModulePresenterProtocol!
    
    required init(presenter: OldModulePresenterProtocol) {
       self.presenter = presenter
   	}    
    
}

