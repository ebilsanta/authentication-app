import styles from './login.module.css'

export default function LoginLayout({ children }) {
  return (
    <div className={"w-screen h-screen flex flex-col" + styles.gradientani}>
        {children}
    </div>
  );
}
