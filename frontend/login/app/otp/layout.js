import styles from "./otp.module.css"
export default function OtpLayout({ children }) {
  return (
    <div className={"w-screen h-screen flex flex-col " + styles.bg}>
      {children}
    </div>
  );
}
