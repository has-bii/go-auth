"use server"

import { fetchHelper } from "@/utils/fetch"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import { z } from "zod"

type State = {
  success: boolean
  message: string
  errors?: {
    email?: string[]
    password?: string[]
  }
}

const formSchema = z.object({
  email: z.string().email("Invalid email!"),
  password: z.string().min(6, "Min. 6 characters!").max(255, "Max 255 characters!"),
})

export default async function login(_: State | null, formData: FormData): Promise<State> {
  try {
    const { data, error } = formSchema.safeParse({
      email: formData.get("email"),
      password: formData.get("password"),
    })

    if (error) {
      return {
        success: false,
        message: "Invalid data!",
        errors: error.flatten().fieldErrors,
      }
    }

    const res = await fetchHelper("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })

    const resData = await res.json()

    if (!res.ok) {
      return { success: false, message: resData.message }
    }

    const cookieStore = await cookies()
    const expiresAt = new Date(Date.now() + 7 * 24 * 60 * 60 * 1000)
    cookieStore.set("access_token", resData.access_token, {
      httpOnly: true,
      secure: true,
      sameSite: "lax",
      expires: expiresAt,
    })
  } catch (error) {
    if (error instanceof Error) return { success: false, message: error.message }

    return { success: false, message: "Internal server error!" }
  }

  redirect("/")
}
