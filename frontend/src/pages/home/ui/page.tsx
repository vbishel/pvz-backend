import { useNavigate } from "@tanstack/react-router"
import { useEffect } from "react"

export const HomePage = () => {

    const navigate = useNavigate()

    useEffect(() => {
        navigate({ to: "/register" })
    }, [])

    return <></>;
}
