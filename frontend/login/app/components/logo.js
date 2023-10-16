import Image from 'next/image'
export default function Logo() {
  return (
    <Image
      className='pt-10 pl-10'
      src="/logo.svg"
      alt="Ascenda Loyalty"
      width={150}
      height={30}
      priority
    />)
}
