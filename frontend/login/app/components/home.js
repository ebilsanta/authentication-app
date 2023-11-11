import Link from 'next/link';

export default function Home() {
    return <div className="m-auto">
      <div>
        <Link href="/profile" className="hover:text-white text-gray-500 font-semibold m-10 py-5 text-xl hover:duration-500">Profile</Link>
      </div>
    </div>
  }