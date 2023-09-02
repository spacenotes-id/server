import { Button } from '@/components/button'
import { Paper } from '@/components/paper'

import { spaceList } from '@/db/space'

import { ArchiveIcon, HeartIcon, OrbitIcon, PencilIcon, Trash2Icon } from 'lucide-react'

export default function SpacePage() {
  return (
    <>
      <p className='font-bold text-2xl lg:text-3xl 2xl:text-4xl'>Manage your space</p>
      <hr className='my-4' />

      <div className='grid gap-3 grid-cols-[repeat(auto-fit,minmax(min(100%,8rem),1fr))]'>
        <Paper className='p-4 rounded'>
          <div className='inline-flex items-center text-xl lg:text-2xl 2xl:text-3xl font-semibold'>
            <OrbitIcon className='mr-2 text-primary-600' size='1em' />
            <span>64</span>
          </div>
          <p className='mt-2'>Space created</p>
        </Paper>
        <Paper className='p-4 rounded'>
          <div className='inline-flex items-center text-xl lg:text-2xl 2xl:text-3xl font-semibold'>
            <HeartIcon className='mr-2 text-rose-600' size='1em' />
            <span>128</span>
          </div>
          <p className='mt-2'>Favorite Space</p>
        </Paper>
        <Paper className='p-4 rounded'>
          <div className='inline-flex items-center text-xl lg:text-2xl 2xl:text-3xl font-semibold'>
            <ArchiveIcon className='mr-2 text-violet-600' size='1em' />
            <span>14</span>
          </div>
          <p className='mt-2'>Space archived</p>
        </Paper>
      </div>

      <div className='mt-8'>
        <p className='font-bold text-lg lg:text-xl 2xl:text-2xl mb-4'>Your spaces</p>
        <div className='grid gap-3 grid-cols-[repeat(auto-fit,minmax(min(100%,14rem),1fr))]'>
          {spaceList.map((item) => {
            return (
              <Paper key={item.id} className='flex flex-col px-4 py-2 h-28 rounded'>
                <p className='font-semibold mb-2'>{item.name}</p>

                <hr />

                <div className='flex items-center space-x-2 mt-auto ml-auto'>
                  <Button icon={PencilIcon} variants='ghost' className='py-1 px-2 text-xs'>
                    Modify
                  </Button>
                  <Button icon={Trash2Icon} variants='danger' className='py-1 px-3 text-xs'>
                    Delete
                  </Button>
                </div>
              </Paper>
            )
          })}
        </div>
      </div>
    </>
  )
}
