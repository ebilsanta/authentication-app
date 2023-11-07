import Link from 'next/link';

export default function Menu() {
    return <div className="m-auto">
      <div>
        <button class="peer font-semibold py-5 text-4xl text-white">Start your journey with us</button>
        <div class="hidden peer-hover:flex hover:flex w-[300px] flex-col drop-shadow-lg">
          <Link href="/register" className="hover:text-white text-gray-500 font-semibold py-5 text-xl hover:duration-500">Register</Link>
          <Link href="/login" className="hover:text-white text-gray-500 font-semibold py-5 text-xl hover:duration-500">Login</Link>
        </div>
      </div>
    </div>
  }
  