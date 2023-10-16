import Image from 'next/image'
import styles from './login.module.css'

export default function LoginLayout({ children }) {
  return (
    <div className={"w-screen h-screen flex flex-col bg-[#050A24] " + styles.gradientani}>
      <Image
        className='pt-5 pl-5'
        src="/logo.svg"
        alt="Ascenda Loyalty"
        width={150}
        height={30}
        priority
      />
      <div className='w-full my-auto'>
        {children}
      </div>
    </div>
  );
}
