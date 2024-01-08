import axios from 'axios'

export const checkAuthentication = async () => {

    try {
        
        const response = await axios.get('http://localhost:900/user/data', {
            withCredentials: true
        });

        console.log('authenticated', response.data)

        if(response.data) return true
        
    } catch (error) {
        console.error(error)
    }
}