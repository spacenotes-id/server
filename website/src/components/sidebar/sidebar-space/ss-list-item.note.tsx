import { Button } from '@/components/button'

import type { Note } from '@/db/note'

import { StickyNoteIcon } from 'lucide-react'
import { useRouter } from 'next/navigation'

export function SidebarSpaceListItemNote(props: Note) {
  const router = useRouter()

  const onClick = () => router.push('/dashboard/note/' + props.id)

  return (
    <Button
      onClick={onClick}
      icon={StickyNoteIcon}
      variants='unstyled'
      className='py-1 px-2 text-base-600 hover:bg-base-100'
    >
      <span>{props.name}</span>
    </Button>
  )
}
