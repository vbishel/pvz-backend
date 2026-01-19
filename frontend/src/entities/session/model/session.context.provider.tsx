import { useQuery } from "@tanstack/react-query"
import { sessionApi } from ".."
import { SessionContext } from "./session.context"


export const SessionContextProvider = ({ children }: { children: React.ReactNode }) => {

    const { data: session, isLoading } = useQuery(
        sessionApi.queries.me()
    )

    if (isLoading) return <></>

    return (
        <SessionContext.Provider value={{ isAuthorized: !!session?.email, session: session }}>
            { children }
        </SessionContext.Provider>
    )
}
