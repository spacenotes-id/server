import type { VariantProps } from 'class-variance-authority'
import { cva } from 'class-variance-authority'
import type { LucideIcon } from 'lucide-react'
import { createElement } from 'react'
import { twMerge } from 'tailwind-merge'
import { P, match } from 'ts-pattern'

const button = cva(['text-sm font-medium rounded border'], {
  variants: {
    variants: {
      primary: [
        'border-transparent',
        'bg-primary-200 text-black',
        'motion-safe:transition',
        'motion-safe:hover:bg-primary-300',
        'motion-safe:active:bg-primary-400',
      ],
      ghost: [
        'bg-base-50 text-base-600',
        'motion-safe:transition',
        'motion-safe:hover:bg-base-100',
        'motion-safe:active:bg-base-200',
      ],
      danger: [
        'bg-red-100 text-red-900',
        'motion-safe:transition',
        'motion-safe:hover:bg-red-200',
        'motion-safe:active:bg-red-300',
      ],
    },
  },
  defaultVariants: {
    variants: 'primary',
  },
})

export type ButtonProps = VariantProps<typeof button> &
  React.ComponentProps<'button'> & {
    icon?: LucideIcon
  }

export function Button({ variants, icon, children, ...props }: ButtonProps) {
  const withIcon = match(icon)
    .with(P.nullish, () => null)
    .otherwise((Icon) => <Icon size='1em' className='mr-1' />)

  // eslint-disable-next-line react/no-children-prop
  return createElement<ButtonProps, HTMLButtonElement>(
    'button',
    {
      ...props,
      className: twMerge(
        button({
          variants,
          className: twMerge(props.className, icon && 'inline-flex items-center'),
        }),
      ),
    },
    withIcon,
    children,
  )
}
