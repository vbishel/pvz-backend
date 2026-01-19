import { useSessionContext } from "@/entities/session/model/session.context"
import { Outlet, useNavigate } from "@tanstack/react-router"
import { useEffect } from "react";


export const GuestGuard = () => {
    const { isAuthenticated } = useSessionContext();
    const navigate = useNavigate()

    useEffect(() => {
        if (isAuthenticated) navigate({ to: "/" })
    }, [isAuthenticated])

    if (isAuthenticated) return <></>

    return (
        <Outlet />
    )
}
