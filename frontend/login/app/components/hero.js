import Logo from "./logo";
import styles from "./hero.module.css"

export default function Hero() {
  return <div className={"w-1/2 h-full flex flex-col " + styles.bg}>
    <Logo />
    <h1 className={"px-10 pb-20  mt-auto text-5xl font-medium " + styles.textgradient}>
      Accelerate your growth with world-class rewards
    </h1>
  </div>;
}
