import { Header } from '@/components/header'
import { Sidebar } from '@/components/sidebar'

import { Fragment } from 'react'

export default function DashboardLayout(props: React.PropsWithChildren) {
  return (
    <Fragment>
      <Header />

      <Sidebar />

      <div className='sm:pl-52 md:pl-60 xl:pl-72 pt-14'>
        <div className='mx-auto w-11/12 py-4'>{props.children}</div>
      </div>
    </Fragment>
  )
}
