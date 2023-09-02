import { nanoid } from 'nanoid'
import { cache } from 'react'

export type Note = { id: string; name: string }

export const getNoteList = cache(() => {
  return [
    { id: nanoid(), name: 'Note number 1' },
    { id: nanoid(), name: 'Note number 2' },
    { id: nanoid(), name: 'Note number 3' },
    { id: nanoid(), name: 'Note number 4' },
    { id: nanoid(), name: 'Note number 5' },
    { id: nanoid(), name: 'Note number 6' },
  ] as Note[]
})
