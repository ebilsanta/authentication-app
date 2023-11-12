"use client"

import Hero from "../../components/hero.js"
import PasswordForm from "@/app/components/passwordForm.js"

export default function SetPassword() {
  return <div className="flex sm:flex-row flex-col w-screen h-screen">
    <Hero />
    <PasswordForm />
  </div>
}
