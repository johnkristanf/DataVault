import { useState } from "react";
import { LoginForm } from "../components/auth/loginForm";
import { SignUpForm } from "../components/auth/signupForm";
import { DataVaultDesc } from "../components/dataVaultDesc";


const LandingPage = () => {

  const [renderLoginForm, setrenderLoginForm] = useState<boolean>(true)
  

    return(
        <div className="h-screen bg-slate-800 flex items-center justify-evenly">
          <DataVaultDesc />

  
            { 
              renderLoginForm 
                ? <LoginForm setrenderLoginForm={setrenderLoginForm} /> 
                : <SignUpForm setrenderLoginForm={setrenderLoginForm}/>
            }


        </div>
    )
}

export default LandingPage
