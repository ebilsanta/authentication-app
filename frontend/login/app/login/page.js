"use client"

import Hero from '../components/hero'
import LoginForm from '../components/loginForm'
export default function Login() {
  return <div className="flex sm:flex-row flex-col w-screen h-screen">
    <Hero />
    <LoginForm />
  </div>
}
