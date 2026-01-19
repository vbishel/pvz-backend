import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanStackDevtools } from '@tanstack/react-devtools'

import { Devtools } from "@/shared/lib/react-query"

import type { QueryClient } from '@tanstack/react-query'
import { RootLayout } from '@/app/layouts/root-layout'

interface MyRouterContext {
    queryClient: QueryClient
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
    component: () => (
        <RootLayout>
            <Outlet />
            <TanStackDevtools
                config={{
                    position: 'bottom-right',
                }}
                plugins={[
                    {
                        name: 'Tanstack Router',
                        render: <TanStackRouterDevtoolsPanel />,
                    },
                    Devtools,
                ]}
            />
        </RootLayout>
    ),
})
