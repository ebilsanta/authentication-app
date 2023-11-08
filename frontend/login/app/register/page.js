"use client"

import Hero from "../components/hero";
import RegisterForm from "../components/registerForm";

export default function Register() {
  return <div className="flex sm:flex-row flex-col w-screen h-screen">
    <Hero />
    <RegisterForm />
  </div>
}
