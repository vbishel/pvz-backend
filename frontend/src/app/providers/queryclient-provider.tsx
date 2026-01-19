import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { isAxiosError } from 'axios';

const NO_RETRY_STATUSES = [400, 401, 403, 404, 409, 422, 500];

export function getContext() {
    const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                retry: (failureCount, error) => {
                    if (
                        isAxiosError(error) &&
                        error.status &&
                        NO_RETRY_STATUSES.includes(error.status)
                    ) {
                        return false;
                    }

                    return failureCount < 4;
                },
            },
            mutations: {
                retry: false,
            },
        },
    });
    return {
        queryClient,
    }
}

export function Provider({
    children,
    queryClient,
}: {
    children: React.ReactNode
    queryClient: QueryClient
}) {
    return (
        <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    )
}
