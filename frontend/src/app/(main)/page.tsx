import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
} from "@/components/ui/breadcrumb"
import { Separator } from "@/components/ui/separator"
import { SidebarTrigger } from "@/components/ui/sidebar"
import { getWorkspaces } from "@/hooks/get-workspaces"
import Workspace from "@/components/workspace"
import { Button } from "@/components/ui/button"
import { Plus } from "lucide-react"

export default async function Home() {
  const workspaces = await getWorkspaces()

  return (
    <>
      <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div className="flex items-center gap-2 px-4">
          <SidebarTrigger className="-ml-1" />
          <Separator orientation="vertical" className="mr-2 h-4" />
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem>
                <BreadcrumbPage>Workspaces</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </div>
      </header>
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div className="flex items-center justify-between lg:justify-start gap-2">
          <h2 className="text-2xl font-semibold">My Workspaces</h2>
          <Button size="sm">
            <Plus />
            Tambah
          </Button>
        </div>
        {workspaces?.map((workspace, i) => (
          <Workspace key={i} data={workspace} />
        ))}
      </div>
    </>
  )
}
