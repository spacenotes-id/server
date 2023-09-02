'use client'

import { ThemeProvider } from 'next-themes'

export default function Providers(props: React.PropsWithChildren) {
  return (
    <ThemeProvider attribute='class' storageKey='app_theme' enableSystem>
      {props.children}
    </ThemeProvider>
  )
}
