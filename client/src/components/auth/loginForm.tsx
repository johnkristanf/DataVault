import '../../../public/loginFormStyle.css'

import { useForm } from 'react-hook-form'
import { LoginCredTypes } from '../../types/authCredType'
import { login } from '../../services/http/auth/login'
import { useNavigate } from 'react-router-dom'

export const LoginForm = ({ setrenderLoginForm }: any) => {

    const navigate = useNavigate();

    const { register, handleSubmit, formState: { errors }, reset } = useForm<LoginCredTypes>()

    const onSubmit = async (loginCredentials : LoginCredTypes) => {

      reset()

      if(await login(loginCredentials)) navigate('/vault')

    }

    return(
        
    <form onSubmit={handleSubmit(onSubmit)} className="form h-[54%]">

       <p className="form-title mb-4">Login in to your account</p>

        <div className="input-container">

          <input placeholder="Enter email" type="email" 
              {...register("email", { required: true })} 
          />

          <span>
            <svg stroke="currentColor" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
            </svg>
          </span>

        </div>

        { errors.email && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>Email is Required</p> }

        <div className="input-container">

            <input placeholder="Enter password" type="password" 
               {...register("password", { required: true })} 
            />

            <span>
                <svg stroke="currentColor" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
                  <path d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" strokeWidth="2" strokeLinejoin="round" strokeLinecap="round"></path>
                </svg>
            </span>

        </div>

          { errors.password && <p className='text-red-700 text-[12px] font-bold mb-2 ml-1'>Password is Required</p> }

          <button className="submit text-slate-300 font-bold bg-lime-500" type="submit">Sign in</button>


      <p className="signup-link">
        No account?
        <a className='hover:cursor-pointer' onClick={() => setrenderLoginForm(false)}>Sign up</a>
      </p>

    </form>

    )
}

