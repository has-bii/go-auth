export const fetchHelper = (input: string | URL | globalThis.Request, init?: RequestInit) =>
  fetch(`${process.env.NEXT_PUBLIC_API_URL}${input}`, init)
