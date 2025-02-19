import React from "react"
import { Toaster } from "@/components/ui/sonner"
import QueryProvider from "./query-provider"
import { CookiesProvider } from "next-client-cookies/server"

type Props = {
  children: React.ReactNode
}

export default function RootProvider({ children }: Props) {
  return (
    <>
      <CookiesProvider>
        <QueryProvider>{children}</QueryProvider>
      </CookiesProvider>
      <Toaster richColors />
    </>
  )
}
