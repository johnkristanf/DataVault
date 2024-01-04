
export type SignupCredTypes = {
    firstname: string,
    lastname: string,
    email: string 
    password: string
}

export const emailPattern = /^[\w-]+(\.[\w-]+)*@([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$/i


export type LoginCredTypes = {
    email: string 
    password: string
}