
import Foundation

class OldModulePresenter: OldModulePresenterProtocol {
    required init(view: OldModuleViewProtocol) {
       self.view = view
    }
    
    
    var view: OldModuleViewProtocol!
	
	
	var interactor: OldModuleInteractorProtocol!
	var router: OldModuleRouterProtocol!
    
    
}

