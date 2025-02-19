import { Workspace } from "@/types/model"
import { fetchHelper } from "@/utils/fetch"
import { cookies } from "next/headers"

type Response = {
  message: string
  data: Workspace[] | null
}

export const getWorkspaces = async () => {
  const cookieStore = await cookies()
  const access_token = cookieStore.get("access_token")?.value as string

  const res = (await fetchHelper("/workspace", {
    method: "GET",
    headers: {
      Authorization: `Bearer ${access_token}`,
    },
  }).then((res) => res.json())) as Response

  return res.data
}
