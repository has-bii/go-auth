"use client"

import { Kanban } from "lucide-react"

import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar"
import Link from "next/link"
import { Workspace } from "@/types/model"
import { usePathname } from "next/navigation"

type Props = {
  data: Workspace[]
}

export function NavWorkspaces({ data }: Props) {
  const pathname = usePathname()
  const { isMobile, setOpenMobile } = useSidebar()

  return (
    <SidebarGroup>
      <SidebarGroupLabel>Workspaces</SidebarGroupLabel>
      <SidebarMenu>
        {data?.map((workspace, i) => (
          <SidebarMenuItem key={i}>
            <SidebarMenuButton asChild isActive={`/workspace/${workspace.id}` === pathname}>
              <Link
                href={`/workspace/${workspace.id}`}
                onClick={() => {
                  if (isMobile) setOpenMobile(false)
                }}
              >
                <Kanban />
                <span>{workspace.name}</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        ))}
      </SidebarMenu>
    </SidebarGroup>
  )
}
