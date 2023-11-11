'use client'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
export default function Logo() {
  const router = useRouter()
  return (
    <Image
      className='pt-10 pl-10 mb-10 cursor-pointer'
      src="/logo.svg"
      alt="Ascenda Loyalty"
      width={150}
      height={30}
      priority
      onClick={() => {
        router.push('/')
      }}
    />)
}
