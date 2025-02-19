"use client"

import register from "@/actions/auth/register"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { cn } from "@/lib/utils"
import { Loader } from "lucide-react"
import Link from "next/link"
import React from "react"

export default function Page() {
  const [state, formAction, isLoading] = React.useActionState(register, null)

  return (
    <div className="flex flex-col gap-6">
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl">Register</CardTitle>
          <CardDescription>Enter your information below to sign up to your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form action={formAction}>
            <div className="flex flex-col gap-6">
              {state && (
                <div
                  className={cn(
                    "p-4 border rounded-lg",
                    state.success
                      ? "bg-green-100 border-green-400"
                      : "bg-destructive/15 border-destructive",
                  )}
                >
                  <p
                    className={cn(
                      "text-center text-sm",
                      state.success ? "text-green-500" : "text-destructive",
                    )}
                  >
                    {state.message}
                  </p>
                </div>
              )}
              <div className="grid gap-2">
                <Label htmlFor="name">Full Name</Label>
                <Input
                  id="name"
                  name="name"
                  type="name"
                  placeholder="Ali Ramdani"
                  required
                  defaultValue={state?.rawInputs?.name}
                />
                {state?.errors?.name && (
                  <span className="text-sm text-destructive">{state.errors.name[0]}</span>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                  defaultValue={state?.rawInputs?.email}
                />
                {state?.errors?.email && (
                  <span className="text-sm text-destructive">{state.errors.email[0]}</span>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  name="password"
                  type="password"
                  required
                  defaultValue={state?.rawInputs?.password}
                />
                {state?.errors?.password && (
                  <span className="text-sm text-destructive">{state.errors.password[0]}</span>
                )}
              </div>
              <Button type="submit" className="w-full" disabled={isLoading}>
                {isLoading && <Loader className="animate-spin" />}
                Register
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Already have an account?{" "}
              <Link href="/login" className="underline underline-offset-4">
                Sign in
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
