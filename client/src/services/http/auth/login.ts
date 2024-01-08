import axios from 'axios'
import { LoginCredTypes } from '../../../types/authCredType'

export const login = async (loginCredentials: LoginCredTypes) => {

    try {
        
        const response = await axios.post('http://localhost:900/user/login', loginCredentials, {
            withCredentials: true
        });

        console.log('login res', response.data)

        if(response) return true
        
    } catch (error) {
        console.error(error)
    }
}