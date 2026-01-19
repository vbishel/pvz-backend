import { GuestGuard } from '@/app/guards/guest-guard'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_guest/')({
    component: GuestGuard,
})
