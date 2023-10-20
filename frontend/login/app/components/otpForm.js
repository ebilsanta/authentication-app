"use client"
import { useState } from "react"
import OtpInput from 'react-otp-input';

export default function OtpForm() {
  const [otp, setOtp] = useState('');

  return <div className="bg-white m-auto rounded-md shadow p-5 text-black">
    <h1 className="text-black font-semibold text-2xl py-5"> OTP Verification </h1>
    <OtpInput
      value={otp}
      onChange={setOtp}
      numInputs={6}
      renderSeparator={<span></span>}
      renderInput={(props) => <input {...props} />}
      inputStyle="border border-black rounded-md m-2 text-3xl w-full block"
    />
    <button className="transition duration-300 bg-[#1e244d] shadow hover:bg-blue-700 w-full text-white font-semibold py-2 px-5 rounded focus:outline-none focus:shadow-outline" type="button">
      Register Account
    </button>
  </div>
}
