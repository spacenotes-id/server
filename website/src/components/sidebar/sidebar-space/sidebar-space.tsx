import { SidebarSpaceList } from './ss-list'
import { SidebarSpaceListHeader } from './ss-header'

export function SidebarSpace() {
  return (
    <div>
      <SidebarSpaceListHeader />

      <hr className='mt-3' />

      <SidebarSpaceList />
    </div>
  )
}
