import axios from 'axios'
import { SignupCredTypes } from '../../../types/authCredType'

export const signup = async (signupCredentials: SignupCredTypes) => {

    try {
        
        const response = await axios.post('http://localhost:900/user/signup', signupCredentials)
        if(response.data) return response.data
        
    } catch (error) {
        console.error(error)
    }
}