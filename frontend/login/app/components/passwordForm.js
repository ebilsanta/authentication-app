export default function PasswordForm() {
  return <div className="bg-white m-auto rounded-md p-5">
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
    <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="button">
      Set Password
    </button>
  </div>
}
