import { Workspace as TWorkspace } from "@/types/model"
import React from "react"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "./ui/button"
import { SquareArrowOutUpRight } from "lucide-react"
import Link from "next/link"

type Props = {
  data: TWorkspace
}

export default function Workspace({ data }: Props) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{data.name}</CardTitle>
      </CardHeader>
      <CardFooter>
        <Button asChild size="sm" className="ml-auto">
          <Link href={`/workspace/${data.id}`}>
            <SquareArrowOutUpRight />
            Open
          </Link>
        </Button>
      </CardFooter>
    </Card>
  )
}
