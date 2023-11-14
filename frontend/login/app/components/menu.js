import Link from 'next/link';

export default function Menu() {
    return <div className="bg-white rounded-md p-5 m-auto">
      <div>
        <button className="peer text-[#1e244d] font-semibold py-5 text-4xl">Let's get started</button>
        <div class="hidden peer-hover:flex hover:flex w-[300px] flex-col drop-shadow-lg">
          <Link href="/register" className="hover:text-white text-gray-500 font-semibold py-5 text-xl hover:duration-500">Register</Link>
          <Link href="/login" className="hover:text-white text-gray-500 font-semibold py-5 text-xl hover:duration-500">Login</Link>
        </div>
      </div>
    </div>
  }
  