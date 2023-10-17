import Logo from "../components/logo";
import styles from "./register.module.css"
export default function RegisterLayout({ children }) {
  return (
    <div className={"w-screen h-screen flex flex-col " + styles.bg}>
      {children}
    </div>
  );
}
