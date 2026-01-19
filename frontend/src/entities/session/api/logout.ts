import { AUTH_API } from "@/shared/lib/axios";

export const logout = () => 
    AUTH_API.post<unknown>("/logout")
