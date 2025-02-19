import * as React from "react"

import { NavUser } from "@/components/sidebar/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar"
import Logo from "./logo"
import { NavMain } from "./nav-main"
import { NavWorkspaces } from "./nav-workspaces"
import { getWorkspaces } from "@/hooks/get-workspaces"

export async function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const workspaces = await getWorkspaces()

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <Logo />
      </SidebarHeader>
      <SidebarContent>
        <NavMain />
        <NavWorkspaces data={workspaces ?? []} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
