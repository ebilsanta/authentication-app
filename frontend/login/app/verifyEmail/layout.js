import Logo from '../components/logo';
import styles from './verifyEmail.module.css'

export default function LoginLayout({ children }) {
  return (
    <div className={"w-screen h-screen flex flex-col bg-[#050A24] " + styles.gradientani}>
      <Logo />
      <div className='w-full my-auto'>
        {children}
      </div>
    </div>
  );
}
