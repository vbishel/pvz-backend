import axios from "axios"

export const AUTH_API = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL + import.meta.env.VITE_AUTH_SERVICE_URL,
    withCredentials: true,
})
