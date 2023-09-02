import { tw } from '@/libs/common'

import type { ReactHTML } from 'react'
import { createElement } from 'react'

type TAsElement = keyof ReactHTML

export type TProps = {
  as?: TAsElement
  className?: string
} & React.ComponentProps<TAsElement>

export function Paper(props: TProps) {
  return createElement(props.as ?? 'div', {
    ...props,
    className: tw('bg-white dark:bg-base-800', 'border dark:border-base-700', props.className),
  })
}
