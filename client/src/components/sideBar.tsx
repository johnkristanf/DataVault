import { AddFileBtn } from "./ui/button";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faVault, faTrash, faGear, faRightFromBracket } from '@fortawesome/free-solid-svg-icons';

const link = [
    {name: "My Vault", icon : <FontAwesomeIcon icon={faVault} /> },
    {name: "Trash", icon : <FontAwesomeIcon icon={faTrash} />},
    {name: "Settings", icon : <FontAwesomeIcon icon={faGear} />},
    {name: "Sign Out", icon : <FontAwesomeIcon icon={faRightFromBracket} />},
]


const Logo = () => {
    return  <h1 className="font-semibold text-2xl flex items-center gap-3 mb-3"> 
              <img src="../../public/img/vault_logo.png" className="rounded-full" width={65} alt="" />
              DataVault
            </h1>
}


const SidebarLinks = () => {
    return (


        <ul className="mt-10 w-[70%]">
            {
                link.map((item) => (

                    <li 
                    key={item.name}
                    className="font-semibold rounded-md text-lg p-3 mb-10 hover:bg-gray-200 hover:cursor-pointer"
                    > 

                    { item.icon } {item.name} 

                    </li>
                ))
            }

        </ul>
    )
}

export const SideBar = () => {
    
    return(
        <div className="h-full bg-gray-300 w-[20%] flex flex-col items-start p-5">
            <Logo />
            <AddFileBtn />

            <SidebarLinks />
        </div>
    )
}

