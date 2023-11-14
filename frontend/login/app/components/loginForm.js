import Link from 'next/link';
import { useState, createContext, useContext } from 'react'
import { useRouter } from 'next/navigation'
import { useLogin } from '../context/context'

export default function LoginForm() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const { hostedLogin, setHostedLoginContext } = useLogin();
  const [errMsg, setErrMsg] = useState("");
  const [loading , setLoading] = useState(false)
  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()
    setLoading(true)
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
      setLoading(false) 
      return
    }
  }

  return <div className="bg-white m-auto rounded-md shadow p-5 w-1/3">
    <h1 className="text-black font-semibold py-5 text-xl">Log in to your account</h1>
    {errMsg !== "" && <p className="text-red-500 text-center">{errMsg}</p>}
    <form onSubmit={(e) => {onSubmit(e)}} className="pb-8 mb-4">
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Email
      </label>
      <input value={email} onChange={(e) => setEmail(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your email" />
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Password
      </label>
      <input value={password} onChange={(e) => setPassword(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="password" placeholder="Enter your password"/>
      <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full flex flex-row items-center text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="submit">
        { loading && 
           <svg aria-hidden="true" className="w-4 h-4 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
           <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
           <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
       </svg>
        }
        <p className='mx-auto'>{!loading ? "Sign In": "Signing in..."}</p>
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
