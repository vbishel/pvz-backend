import { SessionContextProvider } from "@/entities/session/model/session.context.provider"
import { Toaster } from "@/shared/ui/sonner"


export const RootLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <SessionContextProvider>
            {children}
            <Toaster />
        </SessionContextProvider>
    )
}
