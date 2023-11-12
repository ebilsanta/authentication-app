import { useRouter } from 'next/navigation'

export default function RequestOTPForm() {

  const router = useRouter()

  async function onSubmit(event) {
    event.preventDefault()
 
    const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/otp", {
      method: "POST",
      credentials: "include",
      headers : {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
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
    <h1 className="text-black font-semibold p-5 text-center text-xl">Change Password</h1>
    <form className="px-8 pb-8 mb-4">
      <button onClick={(e) => onSubmit(e)} className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300">
        Verify with Email
      </button>
      <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300">
        Verify with Phone Number
      </button>
    </form>
  </div>;
}
