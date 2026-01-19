import { useSessionContext } from "@/entities/session/model/session.context"
import { Outlet, useNavigate } from "@tanstack/react-router"
import { useEffect } from "react";


export const AuthGuard = () => {
    const { isAuthenticated } = useSessionContext();
    const navigate = useNavigate()

    useEffect(() => {
        if (!isAuthenticated) navigate({ to: "/auth/register" })
    }, [isAuthenticated])

    if (!isAuthenticated) return <></>

    return (
        <Outlet />
    )
}
