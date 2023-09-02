import { tw } from '@/libs/common'

import { OrbitIcon, StickyNoteIcon } from 'lucide-react'
import Link from 'next/link'

const items = [
  { name: 'Space', href: '/dashboard/space', icon: OrbitIcon },
  { name: 'Note', href: '/dashboard/note', icon: StickyNoteIcon },
]

export function Navbar() {
  return (
    <nav className='flex items-center ml-auto space-x-1'>
      {items.map((item) => {
        return (
          <Link
            key={item.name}
            href={item.href}
            className={tw(
              'inline-flex items-center',
              'py-1 px-2 rounded motion-safe:transition',
              'motion-safe:hover:bg-primary-100',
            )}
          >
            <item.icon size={14} />
            <span className='text-sm font-medium ml-1'>{item.name}</span>
          </Link>
        )
      })}
    </nav>
  )
}
