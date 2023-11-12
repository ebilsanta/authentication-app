"use client"
import { useState } from 'react'
import { useLogin } from '../context/context';
import OtpInput from 'react-otp-input';
import { useRouter } from 'next/navigation'

export default function OtpForm() {
  const [otp, setOtp] = useState('');
  const { hostedLogin } = useLogin();

  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()

    const route = hostedLogin ? "hosted/verify-otp" : "hosted/verify-email"
 
    const res = await fetch("https://client.itsag2t1.com/api/" + route, {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          otp: otp
      })
    })
 
    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
      if (data["message"] === "OTP Verified") {
        router.push("/changePassword")
      } else {
        router.push("/login")
      }
      
    } else {
      return
    }
  }

  return <div className="bg-white m-auto rounded-md shadow p-5 lg:w-1/3 md:w-2/5 sm:w-1/2 -mt-20 mx-auto text-black">
    <h1 className="text-black font-semibold text-2xl py-5"> OTP Verification </h1>
    <OtpInput
      value={otp}
      onChange={setOtp}
      numInputs={6}
      renderSeparator={<span></span>}
      renderInput={(props) => <input {...props} />}
      inputStyle="border border-black rounded-md m-2 text-3xl w-full block"
    />
    <button onClick={(e) => onSubmit(e)} className="transition duration-300 bg-[#1e244d] mt-3 shadow hover:bg-blue-700 w-full text-white font-semibold py-2 px-5 rounded focus:outline-none focus:shadow-outline" type="button">
      Register Account
    </button>
  </div>
}
