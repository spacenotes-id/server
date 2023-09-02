import { Paper } from '@/components/paper'

import { Navbar } from './navbar'

export function Header() {
  return (
    <div className='fixed top-0 inset-x-0'>
      <Paper as='header' className='h-14 sm:pl-52 md:pl-60 xl:pl-72'>
        <div className='flex items-center h-full w-11/12 mx-auto'>
          <Navbar />
        </div>
      </Paper>
    </div>
  )
}
