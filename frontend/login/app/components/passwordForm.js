import Link from 'next/link';
import { useState, createContext, useContext } from 'react'
import { useRouter } from 'next/navigation'

export default function PasswordForm() {

  const [password, setPassword] = useState("")

  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()
 
    const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/change-password", {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          password: password
      })
    })
 
    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
      router.push("/profile")
    } else {
      return
    }
  }

  return <div className="bg-white flex flex-col lg:w-1/3 w-4/5 -mt-20 mx-auto rounded-md shadow p-5">
    <h1 className="text-black font-semibold text-left text-2xl">Email Verified!</h1>
    <p className="text-black text-gray-700 py-3">Your email has been  verified. Please set your password below</p>

    <h2 className="text-black font-semibold text-left text-xl py-3">Set Password</h2>
    <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
      Password
    </label>
    <input className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your password" />
    <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2 pt-3">
      Re-enter your password
    </label>
    <input className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your password" />
    <button onClick={(e) => onSubmit(e)} className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="button">
      Set Password
    </button>
  </div>
}
