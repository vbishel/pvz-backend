import { Input } from "./input";
import { cn } from "@/shared/lib/tailwind";
import { forwardRef } from "react";
import { Label } from "./label";
import { InputError } from "./input-error";

type Props = {
    id?: string;
    label?: string;
    error?: string;
    testId?: string;
    tooltip?: string;
} & React.ComponentProps<typeof Input>;

/**
 * Инпут для формы
 *
 * @param id    html id инпута
 * @param label подпись инпута
 * @param delay задержка перед появлением (анимация)
 * @param error ошибка инпута
 */
export const FormInput = forwardRef<HTMLInputElement, Props>(
    (
        { id, label, disabled, error, className, testId, tooltip, ...props },
        ref,
    ) => {
        return (
            <div className={"flex flex-col gap-4 w-full"}>
                <div className={"flex flex-col gap-1"}>
                    {label && (
                        <Label
                            htmlFor={id ?? props.name}
                            className={cn(
                                "transition",
                                disabled && "opacity-25",
                            )}
                        >
                            { label }
                        </Label>
                    )}
                    <Input
                        autoComplete={"off"}
                        id={id ?? props.name}
                        ref={ref}
                        name={id}
                        data-test={testId}
                        disabled={disabled}
                        className={cn(
                            "transition w-full",
                            error && "border-destructive bg-card",
                            className,
                        )}
                        {...props}
                    />
                </div>
                <InputError testId={`${testId}-error`} error={error} />
            </div>
        );
    },
);
