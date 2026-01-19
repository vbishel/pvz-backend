import { sessionApi } from "@/entities/session";
import { useProtectedSessionContext } from "@/entities/session/model/session.context"
import { Button } from "@/shared/ui/button";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { toast } from "sonner";

export const ProfilePage = () => {
    const { session } = useProtectedSessionContext();
    const [isLogoutPending, setIsLogoutPending] = useState<boolean>(false);

    const { mutateAsync: logout } = useMutation({
        mutationFn: sessionApi.logout,
    })

    const navigate = useNavigate();
    const queryClient = useQueryClient()

    const handleLogout = async () => {
        setIsLogoutPending(true)
        logout()
            .then(async () => {
                await queryClient.invalidateQueries({
                    queryKey: sessionApi.queries.all()
                })
                toast.success("logout successful")
                navigate({ to: "/auth/login" })
            })
            .catch(() => {
                toast.error("logout failed")
            })
            .finally(() => {
                setIsLogoutPending(false)
            })
    }

    return (
        <div className={"flex flex-col gap-4 h-[80vh] w-full justify-center items-center"}>
            <div>
                Welcome, {session.email}!
            </div>
            <Button disabled={isLogoutPending} onClick={handleLogout}>
                Logout
            </Button>
        </div>
    )
}
