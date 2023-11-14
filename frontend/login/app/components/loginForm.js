import Link from 'next/link';
import { useState, createContext, useContext } from 'react'
import { useRouter } from 'next/navigation'
import { useLogin } from '../context/context'

export default function LoginForm() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const { hostedLogin, setHostedLoginContext } = useLogin();
  const [errMsg, setErrMsg] = useState("");

  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()

    const company = "ascenda"
 
    const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/login", {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          company: company,
          email: email,
          password: password
      })
    })
 
    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
      setErrMsg("")
      setHostedLoginContext(true)
      console.log(hostedLogin)
      router.push("/home")
    } else {
      setErrMsg("Invalid email or password!")
      setEmail("")
      setPassword("")
      return
    }
  }

  return <div className="bg-white flex flex-col lg:w-1/3 w-4/5 -mt-20 mx-auto rounded-md shadow">
    <h1 className="text-black font-semibold p-5 text-center text-xl">Log in to your account</h1>
    {errMsg !== "" && <p className="text-red-500 text-center">{errMsg}</p>}
    <form onSubmit={(e) => {onSubmit(e)}} className="px-8 pb-8 mb-4">
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Email
      </label>
      <input value={email} onChange={(e) => setEmail(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your email" />
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Password
      </label>
      <input value={password} onChange={(e) => setPassword(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="password" placeholder="Enter your password"/>
      <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="submit">
        Sign In
      </button>

      <a href="https://client.itsag2t1.com/api/bankSSO/login" type="button" className="bg-[#f5f5f5] shadow hover:bg-[#1e244d] hover:text-[#f5f5f5] mt-3 w-full text-[#1e244d] font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300 text-center">Login with BankSSO</a>
      <text className="block uppercase tracking-wide text-gray-400 text-xs mb-2 text-center pt-5">
        Don&apos;t have an account? <Link href="/register" className="font-bold text-gray-500">Register</Link> instead
        <br></br>
        Yet to verify your email? <Link href="/verifyEmail" className="font-bold text-gray-500">Verify it Now!</Link>
      </text>
    </form>
  </div>;
}
