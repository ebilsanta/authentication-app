import { useState } from 'react'
import { useRouter } from 'next/navigation'

export default function VerifyEmailForm() {
  const [email, setEmail] = useState("")

  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()

    const company = "ascenda"
 
    const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/otp", {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          company: company,
          email: email,
      })
    })
 
    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
      router.push("/otp")
    } else {
      return
    }
  }

  return <div className="bg-white flex flex-col lg:w-1/3 w-4/5 -mt-20 mx-auto rounded-md shadow">
    <h1 className="text-black font-semibold p-5 text-center text-xl">Verify your Account via Email</h1>
    <form onSubmit={(e) => {onSubmit(e)}} className="px-8 pb-8 mb-4">
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Email
      </label>
      <input onChange={(e) => setEmail(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your email" />
      <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="submit">
        Verify Email
      </button>
    </form>
  </div>;
}
