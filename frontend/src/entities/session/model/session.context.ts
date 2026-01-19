import React, { use } from "react";
import type { Session } from "./session.types";

type ContextProps = {
    isAuthenticated: boolean;
    session?: Session;
}

export const SessionContext = React.createContext<ContextProps | null>(null);

export const useSessionContext = () => {
    const context = use(SessionContext);

    if (!context) {
        throw new Error("SessionContext must be used within a SessionContextProvider")
    }

    return context;
}

export const useProtectedSessionContext = () => {
    const context = useSessionContext();

    if (!context.session) {
        throw new Error("Use protected session context only when user is guaranteed to be logged in")
    }

    return context as Required<ContextProps>;
}
