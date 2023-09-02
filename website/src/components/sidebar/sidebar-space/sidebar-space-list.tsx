import { spaceList } from '@/db/space'

import { SidebarSpaceListItem } from './sidebar-space-list-item'

export function SidebarSpaceList() {
  return (
    <div className='flex flex-col space-y-2 px-4 py-3'>
      {spaceList.map((item) => (
        <SidebarSpaceListItem key={item.id} {...item} />
      ))}
    </div>
  )
}
