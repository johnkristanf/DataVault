import axios from 'axios'
import { LoginCredTypes } from '../../../types/authCredType'

export const login = async (loginCredentials: LoginCredTypes) => {

    try {
        
        const response = await axios.post('http://localhost:900/user/login', loginCredentials);

        if(response.data) return response.data
        
    } catch (error) {
        console.error(error)
    }
}