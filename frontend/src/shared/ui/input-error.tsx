import { CircleAlertIcon } from "lucide-react"

type Props = {
    error?: string;
    testId?: string;
};

/**
 * Текстовая ошибка для инпута, подсвеченная красным цветом
 *
 * @param error текст ошибки
 */
export const InputError = ({ error, testId }: Props) => {
    if (!error) return <></>;

    return (
        <div className={"flex gap-2 items-center text-destructive -mt-2"}>
            <CircleAlertIcon size={14} />
            <p
                className={"text-destructive font-medium text-sm"}
                data-test={testId}
            >
                {error}
            </p>
        </div>
    );
};
