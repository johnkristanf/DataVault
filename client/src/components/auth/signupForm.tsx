import '../../../public/signupFormStyle.css'

import { useForm } from "react-hook-form"
import { SignupCredTypes, emailPattern } from '../../types/authCredType';
import { signup } from '../../services/http/auth/signup';
import { useState } from 'react';


export const SignUpForm = ({ setrenderLoginForm }: any) => {

    const [SignedupUser, setSignedUpUser] = useState()

    const { register, handleSubmit, reset, formState: { errors } } = useForm<SignupCredTypes>();

    const onSubmit = async (signupData: SignupCredTypes) =>{

      const signupSuccess = await signup(signupData)

      if(signupSuccess) reset(); setSignedUpUser(signupSuccess)
      
    } 

    return(

    <> 
        
    <form onSubmit={handleSubmit(onSubmit)} className="form">

       <p className="form-title mb-4">Sign up to get started</p>

        { SignedupUser && <p className='text-lime-500 text-center font-bold mb-3'>Sign up Successfully</p> }

        <div className="input-container">
            <input placeholder="Enter First Name" type="text"
             {...register("firstname", { required: true })} 

            />
         
        </div>

        { errors.firstname && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>First Name is Required</p> }

        <div className="input-container">
            <input placeholder="Enter Last Name" type="text"
              {...register("lastname", { required: true })}

            />

        </div>

        { errors.lastname && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>Last Name is Required</p> }



        <div className="input-container">

            <input placeholder="Enter email" type="email"
             {...register("email", { required: true, pattern: emailPattern })} 
            />

          <span>
            <svg stroke="currentColor" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
            </svg>
          </span>

        </div>

            { errors.email?.type === 'required' && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>Email is Required</p> }
            { errors.email?.type === 'pattern' && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>Invalid Email Address</p>  }


        <div className="input-container">

          <input placeholder="Enter password" type="password"
            {...register("password", { required: true, minLength: 8 })} 
          />

          <span>
            <svg stroke="currentColor" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
              <path d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
            </svg>
          </span>

        </div>

            { errors.password?.type === 'required' && <p className='text-red-700 text-[12px] font-bold mb-3 ml-1'>Password is Required</p> }
            { errors.password?.type === 'minLength' && <p className='text-red-700 text-[12px] font-bold mb-3 ml-1'>Password must have atleast 8 characters</p>  }
        

            <button className="submit text-slate-300 font-bold bg-lime-500" type="submit">Sign Up</button>

        <p className="signup-link">
           Already Have an Account?
           <a className='hover:cursor-pointer' onClick={()=> setrenderLoginForm(true)}>Sign In</a>
        </p>

   </form>

   </>  

    )
}