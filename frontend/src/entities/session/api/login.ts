import { AUTH_API } from "@/shared/lib/axios";

type Req = {
    email: string;
    password: string;
}

export const login = (body: Req) =>
    AUTH_API.post(`/login`, body)
