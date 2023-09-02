import { Paper } from '@/components/paper'

import { tw } from '@/libs/common'

import { SidebarSpace } from './sidebar-space'

export function Sidebar() {
  return (
    <Paper
      as='aside'
      className={tw(
        'fixed left-0 inset-y-0',
        'hidden sm:block',
        'mr-4 sm:w-52 md:w-60 xl:w-72',
        'h-full rounded z-50',
      )}
    >
      <div className='h-full max-h-full w-full overflow-y-auto pt-14 pb-8 space-y-6'>
        <SidebarSpace />
      </div>
    </Paper>
  )
}
