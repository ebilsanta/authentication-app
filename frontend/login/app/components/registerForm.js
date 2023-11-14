import Link from 'next/link';
import { useState } from 'react'
import { useRouter } from 'next/navigation'


export default function RegisterForm() {
  const [email, setEmail] = useState("")
  const [firstName, setFirstName] = useState("")
  const [lastName, setLastName] = useState("")
  const [birthdate, setBirthdate] = useState("")
  const [password, setPassword] = useState("")
  const [errMsg, setErrMsg] = useState("")
  const router = useRouter()

  const validatePassword = (password) => {
    const minLength = 12;
    const upperCasePattern = /[A-Z]/;
    const lowerCasePattern = /[a-z]/;
    const numberPattern = /[0-9]/;
    const specialCharPattern = /[^A-Za-z0-9]/;

    return (
        password.length >= minLength &&
        upperCasePattern.test(password) &&
        lowerCasePattern.test(password) &&
        numberPattern.test(password) &&
        specialCharPattern.test(password)
    );
  };

  const isPasswordValid = validatePassword(password);

  async function onSubmit(event) {
    event.preventDefault()

    if (!isPasswordValid) {
      setErrMsg("Password must be at least 12 characters long, contain at least 1 uppercase letter, 1 lowercase letter, 1 number and 1 special character")
      setPassword("")
    }
    const company = "ascenda"
 
    const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/register", {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        company: company,
        email: email,
        firstName: firstName, 
        lastName: lastName,
        birthdate: birthdate,
        password: password
      })
    })
 
    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
        setErrMsg("")
        router.push("/otp")
    } else {
        return
    }
  }

  return <div className="bg-white m-auto rounded-md shadow p-5 w-1/3">
    <h1 className="text-black font-semibold py-5 text-xl">Register for an account</h1>
    {!isPasswordValid && <p className='text-red-300 py-2'>{errMsg}</p>}
    <form onSubmit={(e) => {onSubmit(e)}} className="">
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Email
      </label>
      <input onChange={(e) => setEmail(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-email" type="text" placeholder="Enter your email" />
      <div className="flex flex-wrap -mx-3 mb-3">
        <div className="w-full sm:w-1/2 px-3 mb-6 md:mb-0">
          <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" >
            First Name
          </label>
          <input onChange={(e) => setFirstName(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white" id="grid-first-name" type="text" placeholder="Jane" />
          <p className="hidden text-red-500 text-xs italic">Please fill out this field.</p>
        </div>
        <div className="w-full md:w-1/2 px-3">
          <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" >
            Last Name
          </label>
          <input onChange={(e) => setLastName(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-last-name" type="text" placeholder="Doe" />
        </div>
      </div>
      <div className="flex flex-wrap -mx-3 mb-1">
        <div className="w-full px-3">
          <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" >
            Date of Birth
          </label>
          <input onChange={(e) => setBirthdate(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-birthdate" type="date" placeholder="******************" />
        </div>
      </div>
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Password
      </label>
      <input value={password} onChange={(e) => setPassword(e.target.value)} className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500 mb-6" id="grid-password" type="password" placeholder="Enter a Password" />

      <button className="transition duration-300 bg-[#1e244d] shadow hover:bg-blue-700 w-full text-white font-semibold py-2 px-5 rounded focus:outline-none focus:shadow-outline" type="submit">
        Register Account
      </button>
      <text className="block uppercase tracking-wide text-gray-400 text-xs mb-2 text-center pt-5">
        Already have an account? <Link href="/verifyEmail" className="font-bold text-gray-500">Verify your Email</Link> or <Link href="/login" className="font-bold text-gray-500">Login</Link> instead
      </text>
    </form>
  </div>
}
