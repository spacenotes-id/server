import { useTheme as useNextTheme } from 'next-themes'
import { useCallback, useEffect, useMemo, useState } from 'react'
import { P, match } from 'ts-pattern'

export const useTheme = () => {
  const { theme, setTheme, systemTheme } = useNextTheme()
  const [mounted, setMounted] = useState<boolean>(false)

  const changeTheme = useCallback(() => {
    const newTheme = match(theme)
      .with('dark', () => 'light')
      .otherwise(() => 'dark')
    setTheme(newTheme)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [theme])

  const actualTheme = useMemo(() => {
    return match({ theme, systemTheme })
      .with(
        P.shape({ theme: 'dark' }).or({ systemTheme: 'dark', theme: 'system' }),
        () => 'dark' as const,
      )
      .with(
        P.shape({ theme: 'light' }).or({ systemTheme: 'light', theme: 'system' }),
        () => 'light' as const,
      )
      .otherwise(() => 'light' as const)
  }, [theme, systemTheme])

  useEffect(() => {
    setMounted(true)
  }, [])

  return {
    theme,
    mounted,
    changeTheme,
    systemTheme,
    actualTheme,
  }
}
