
import UIKit

class OldModuleViewController: UIViewController, OldModuleViewProtocol  {

    
    var presenter: OldModulePresenterProtocol!
    let configurator: OldModuleConfiguratorProtocol = OldModuleConfigurator()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        configurator.configure(with: self)
        
    }
     

}


