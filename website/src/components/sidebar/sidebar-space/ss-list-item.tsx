'use client'

import { Button } from '@/components/button'
import { Paper } from '@/components/paper'

import type { SpaceWithNotes } from '@/db/space'
import { tw } from '@/libs/common'

import { SidebarSpaceListItemNote } from './ss-list-item.note'

import { Disclosure } from '@headlessui/react'
import { ChevronDownIcon, FolderIcon } from 'lucide-react'

export function SidebarSpaceListItem(props: SpaceWithNotes) {
  return (
    <Disclosure as={Paper} variants='ghost' className='rounded'>
      {({ open }) => {
        return (
          <>
            <Disclosure.Button
              as={Button}
              icon={FolderIcon}
              variants='ghost'
              className='w-full py-2 px-3 border-transparent'
            >
              <span className='truncate pr-1'>{props.name}</span>
              <ChevronDownIcon
                className={tw('ml-auto shrink-0 motion-safe:transition', open && 'rotate-180')}
                size='1em'
              />
            </Disclosure.Button>

            <Disclosure.Panel as='div' className='flex flex-col space-y-2 px-2 py-1.5'>
              {props.notes.map((note) => {
                return <SidebarSpaceListItemNote key={note.id} {...note} />
              })}
            </Disclosure.Panel>
          </>
        )
      }}
    </Disclosure>
  )
}
