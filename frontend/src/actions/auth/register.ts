"use server"

import { fetchHelper } from "@/utils/fetch"
import { z } from "zod"

type State = {
  success: boolean
  message: string
  rawInputs?: {
    name?: string
    email?: string
    password?: string
  }
  errors?: {
    name?: string[]
    email?: string[]
    password?: string[]
  }
}

const formSchema = z.object({
  name: z.string().min(6, "Min. 6 characters!").max(255, "Max 255 characters!"),
  email: z.string().email("Invalid email!"),
  password: z.string().min(6, "Min. 6 characters!").max(255, "Max 255 characters!"),
})

export default async function register(_: State | null, formData: FormData): Promise<State> {
  try {
    const { data, error } = formSchema.safeParse({
      name: formData.get("name"),
      email: formData.get("email"),
      password: formData.get("password"),
    })

    if (error) {
      return {
        success: false,
        message: "Invalid data!",
        errors: error.flatten().fieldErrors,
        rawInputs: {
          name: formData.get("name")?.toString(),
          email: formData.get("email")?.toString(),
          password: formData.get("password")?.toString(),
        },
      }
    }

    const res = await fetchHelper("/auth/register", {
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

    return { success: true, message: "Registered successfully." }
  } catch (error) {
    if (error instanceof Error) return { success: false, message: error.message }

    return { success: false, message: "Internal server error!" }
  }
}
