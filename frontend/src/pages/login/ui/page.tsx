import { z } from "zod";
import { useForm } from "react-hook-form";
import { FormInput } from "@/shared/ui/form-input";
import { Button } from "@/shared/ui/button";
import { useMutation } from "@tanstack/react-query";
import { sessionApi } from "@/entities/session";
import { toast } from "sonner";
import { isAxiosError } from "axios";
import { zodResolver } from "@hookform/resolvers/zod"
import { Link, useNavigate } from "@tanstack/react-router";

export const LoginPage = () => {
    const Schema = z.object({
        email: z.email(),
        password: z.string().min(1),
    })

    const {
        register,
        formState: { errors, isSubmitting },
        handleSubmit
    } = useForm<z.infer<typeof Schema>>({
        defaultValues: {
            email: "",
            password: "",
        },
        resolver: zodResolver(Schema)
    })

    const navigate = useNavigate();

    const { mutateAsync: login } = useMutation({
        mutationFn: sessionApi.login,
    })

    const onSubmit = handleSubmit(async ({ email, password }) => {
        await login({
            email,
            password
        })
            .then(() => {
                toast.success("Account created successfully")
                navigate({ to: "/profile" })
            })
            .catch((err) => {
                if (isAxiosError(err) && typeof err.response?.data.error === "string") {
                    toast.error(err.response.data.error)
                    return
                }
                toast.error("An unknown error occured")
            })
    })

    return (
        <div className={"flex w-full h-[80vh] justify-center items-center flex-col gap-4"}>
            <h6 className={"text-sm"}>Log in to your account</h6>
            <form noValidate onSubmit={onSubmit} className={"flex flex-col gap-4 w-70"}>
                <FormInput
                    label={"Email"}
                    error={errors.email?.message}
                    {...register("email")}
                />
                <FormInput
                    label={"Password"}
                    error={errors.password?.message}
                    {...register("password")}
                />
                <Button disabled={isSubmitting}>
                    Log in
                </Button>
                <div className={"flex gap-1 w-full justify-center"}>
                    <span>Don't have an account?</span>
                    <span className={"underline"}>
                        <Link to={"/auth/register"}>Register</Link>
                    </span>
                </div>
            </form>
        </div>
    )
}
