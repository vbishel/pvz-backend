import { AUTH_API } from "@/shared/lib/axios";
import type { Session } from "../model/session.types";

type Res = Session

export const getMe = () =>
    AUTH_API.get<Res>(`/me`)
        .then(data => data.data)
