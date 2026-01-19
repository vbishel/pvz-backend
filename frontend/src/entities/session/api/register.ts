import { AUTH_API } from "@/shared/lib/axios";

type Req = {
    email: string;
    password: string;
}

export const register = (body: Req) =>
    AUTH_API.post(`/register`, body)
