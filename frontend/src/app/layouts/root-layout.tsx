import { SessionContextProvider } from "@/entities/session/model/session.context.provider"


export const RootLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <SessionContextProvider>
            {children}
        </SessionContextProvider>
    )
}
