'use client'

import { useTheme } from '@/hooks/use-theme'

import { Loader2Icon, MoonStarIcon, SunDimIcon } from 'lucide-react'
import { match } from 'ts-pattern'

export function ThemeSwitch() {
  const { changeTheme, actualTheme, mounted } = useTheme()

  return match({ changeTheme, actualTheme, mounted })
    .with({ mounted: false }, () => (
      <button className='h-8 w-8 inline-flex items-center justify-center'>
        <Loader2Icon size={14} className='animate-spin' />
      </button>
    ))
    .otherwise(({ actualTheme, changeTheme }) => (
      <button
        onClick={changeTheme}
        className='inline-flex items-center justify-center w-8 h-8 ml-auto'
      >
        {match(actualTheme)
          .with('dark', () => <MoonStarIcon size={14} />)
          .otherwise(() => (
            <SunDimIcon size={14} />
          ))}
      </button>
    ))
}
