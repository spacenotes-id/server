'use client'

import { Button } from '@/components/button'

import { PlusIcon } from 'lucide-react'
import { usePathname } from 'next/navigation'
import { P, match } from 'ts-pattern'

export function SidebarSpaceListHeader() {
  const pathname = usePathname()

  const addButton = match(pathname)
    .with(P.string.endsWith('/space'), () => null)
    .otherwise(() => (
      <Button variants='ghost' className='p-0.5 ml-auto'>
        <PlusIcon size='1em' />
        <span className='sr-only'>Create space</span>
      </Button>
    ))

  return (
    <div className='flex items-center px-4'>
      <p className='font-medium text-sm'>Your Spaces</p>

      {addButton}
    </div>
  )
}
