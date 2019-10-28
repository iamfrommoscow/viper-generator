
import Foundation

class OldModuleConfigurator: OldModuleConfiguratorProtocol {
    func configure(with viewController: OldModuleViewProtocol) {
		let presenter = OldModulePresenter(view: viewController)
		viewController.presenter = presenter
		
		
	let interactor = OldModuleInteractor(presenter: presenter)
	presenter.interactor = interactor
	let router = OldModuleRouter()
	presenter.router = router
        
    }
}