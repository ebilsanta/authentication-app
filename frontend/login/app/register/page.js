import Hero from "../components/hero";
import RegisterForm from "../components/registerForm";

export default function Register() {
  return <div className="flex flex-row w-screen h-screen">
    <Hero />
    <RegisterForm />
  </div>
}
