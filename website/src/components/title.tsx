export function Title(props: React.PropsWithChildren) {
  return (
    <>
      <p className='font-bold text-2xl lg:text-3xl 2xl:text-4xl'>{props.children}</p>
      <hr className='my-4' />
    </>
  )
}
