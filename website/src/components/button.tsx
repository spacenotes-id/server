import type { VariantProps } from 'class-variance-authority'
import { cva } from 'class-variance-authority'
import type { LucideIcon } from 'lucide-react'
import { createElement, forwardRef } from 'react'
import { twMerge } from 'tailwind-merge'
import { P, match } from 'ts-pattern'

const button = cva(['text-sm font-medium rounded border'], {
  variants: {
    variants: {
      unstyled: ['border-none'],
      primary: ['border-transparent', 'bg-primary-200 text-black'],
      ghost: ['bg-base-50 text-base-600'],
      danger: ['border-transparent', 'bg-red-100 text-red-900'],
    },
    interactive: {
      yes: '',
      no: '',
    },
  },
  compoundVariants: [
    {
      variants: 'primary',
      interactive: 'yes',
      className: [
        'motion-safe:transition',
        'motion-safe:hover:bg-primary-300',
        'motion-safe:active:bg-primary-400',
      ],
    },
    {
      variants: 'ghost',
      interactive: 'yes',
      className: [
        'motion-safe:transition',
        'motion-safe:hover:bg-base-100',
        'motion-safe:active:bg-base-200',
      ],
    },
    {
      variants: 'danger',
      interactive: 'yes',
      className: [
        'motion-safe:transition',
        'motion-safe:hover:bg-red-200',
        'motion-safe:active:bg-red-300',
      ],
    },
  ],
  defaultVariants: {
    variants: 'primary',
  },
})

export type ButtonProps = VariantProps<typeof button> &
  React.ComponentProps<'button'> & {
    icon?: LucideIcon
  }

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ icon, variants, interactive, ...props }, ref) => {
    const withIcon = match(icon)
      .with(P.nullish, () => null)
      .otherwise((Icon) => <Icon size='1em' />)

    // eslint-disable-next-line react/no-children-prop
    return createElement<ButtonProps, HTMLButtonElement>(
      'button',
      {
        ...props,
        ref,
        className: twMerge(
          button({
            variants,
            interactive,
            className: twMerge(props.className, icon && 'inline-flex items-center space-x-1'),
          }),
        ),
      },
      withIcon,
      props.children,
    )
  },
)

Button.displayName = 'button'
