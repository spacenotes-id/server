import { cva } from 'class-variance-authority'
import type { VariantProps } from 'class-variance-authority'
import type { ReactHTML } from 'react'
import { createElement, forwardRef } from 'react'
import { twMerge } from 'tailwind-merge'

const paper = cva(['border'], {
  variants: {
    variants: {
      base: ['bg-white'],
      ghost: ['bg-base-50'],
    },
  },
  defaultVariants: {
    variants: 'base',
  },
})

type PaperProps = VariantProps<typeof paper>

export type TProps = {
  as?: keyof ReactHTML
  className?: string
} & React.ComponentProps<keyof ReactHTML> &
  PaperProps

export const Paper = forwardRef<HTMLElement, TProps>((params, ref) => {
  const { as: componentTypes, variants, className, ...props } = params

  return createElement(componentTypes ?? 'div', {
    ...props,
    ref,
    className: twMerge(paper({ variants, className })),
  })
})

Paper.displayName = 'Paper'
