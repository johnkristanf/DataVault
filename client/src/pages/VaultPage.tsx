import { SideBar } from "../components/sideBar";

import { useEffect, useState } from "react";
import { checkAuthentication } from "../services/http/auth/isAuthenticated";


const VaultPage = () => {

    const [isAuthenticated, setisAuthenticated] = useState<boolean>();

    useEffect(() => {
      const userData = async () => {
        const data = await checkAuthentication();
        setisAuthenticated(data);
      };

      userData();
    }, []); 

    console.log('isAuthenticated', isAuthenticated);

    
    if(isAuthenticated){

      return(

        <div className="h-screen w-full flex">
          <SideBar />
        </div>
        
      )

    } 

    
};
  

export default VaultPage