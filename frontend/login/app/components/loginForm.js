import Link from 'next/link';

export default function LoginForm() {
  return <div className="bg-white flex flex-col lg:w-1/3 w-4/5 -mt-20 mx-auto rounded-md shadow">
    <h1 className="text-black font-semibold p-5 text-center text-xl">Log in to your account</h1>
    <form className="px-8 pb-8 mb-4">
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Email
      </label>
      <input className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your email" />
      <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
        Password
      </label>
      <input className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-password" type="text" placeholder="Enter your password" />
      <button className="bg-[#1e244d] shadow hover:bg-blue-700 mt-3 w-full text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300" type="button">
        Sign In
      </button>
      <text className="block uppercase tracking-wide text-gray-400 text-xs mb-2 text-center pt-5">
        Don&apos;t have an account? <Link href="/register" className="font-bold text-gray-500">Register</Link> instead
      </text>
    </form>
  </div>;
}
