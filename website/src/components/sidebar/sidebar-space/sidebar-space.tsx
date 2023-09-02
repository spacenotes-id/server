import { SidebarSpaceList } from './sidebar-space-list'
import { SidebarSpaceListHeader } from './sidebar-space-list-header'

export function SidebarSpace() {
  return (
    <div>
      <SidebarSpaceListHeader />

      <hr className='mt-3' />

      <SidebarSpaceList />
    </div>
  )
}
