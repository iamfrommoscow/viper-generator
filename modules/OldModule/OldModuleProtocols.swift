
import Foundation

protocol OldModuleViewProtocol: class{
    var presenter: OldModulePresenterProtocol! { get set }
   
}

protocol OldModuleConfiguratorProtocol: class {
    func configure(with viewController: OldModuleViewProtocol)
}

protocol OldModulePresenterProtocol: class {
	var view: OldModuleViewProtocol! {get set}
	
	
	var interactor: OldModuleInteractorProtocol! {get set}
	var router: OldModuleRouterProtocol! {get set}
   
}



protocol OldModuleInteractorProtocol: class {
	var presenter: OldModulePresenterProtocol! { get set }
}
	
protocol OldModuleRouterProtocol: class {

}

