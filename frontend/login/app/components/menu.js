import Link from 'next/link';

export default function Menu() {
    return <div className="bg-white rounded-md p-5 m-auto w-1/3">
      <div>
        <button className="text-black font-semibold py-5 text-xl">Let's get started</button>
        <div>
        <Link href="/login"> 
          <button className="my-5 transition duration-300 bg-[#1e244d] shadow hover:bg-blue-700 w-full text-white font-semibold py-2 px-5 rounded focus:outline-none focus:shadow-outline" type="submit">
            Login
          </button>
          </Link>

          <hr></hr>
         <Link href="/register"> 
          <button className="my-5 transition duration-300 bg-[#1e244d] shadow hover:bg-blue-700 w-full text-white font-semibold py-2 px-5 rounded focus:outline-none focus:shadow-outline" type="submit">
            Register
          </button>
          </Link>
        </div>
      </div>
    </div>
  }
  