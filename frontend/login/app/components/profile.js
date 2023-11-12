import Link from 'next/link';
import { useState, useEffect } from 'react'
import { useLogin } from '../context/context';

export default function Profile() {
  const [profile, setProfile] = useState("")
  const { hostedLogin } = useLogin();

  console.log(hostedLogin);

  useEffect(() => {
  const fetchProfile = async() => {
    if (hostedLogin == true) {
      var userinfo = "hosted/user";
    } else {
      var userinfo = "bankSSO/userInfo";
    }

    console.log(userinfo);
    const res = await fetch("https://client.itsag2t1.com/api/" + userinfo, {
        method: "GET",
        credentials: "include",
        headers : {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }});
        
      const data = await res.json();
      console.log(data);
      var date_str = data.birthdate;
      var birthday = new Date(date_str);
      const formattedDate = birthday.toISOString().slice(0, 10);
      data.birthdate = formattedDate;
      setProfile(data);
    };
    fetchProfile();
  }, []);


return <div className="m-auto">
      <div>
        <Link href="/home" className="hover:text-white text-gray-500 font-semibold m-10 py-5 text-xl hover:duration-500">Home</Link>
      </div>
 <div className="flex flex-col lg:w-1/3 w-4/5 mt-20 mx-auto rounded-md">
    <h1 className="block uppercase tracking-wide text-white font-bold py-5 text-xl">My Profile</h1>
      <div>
        <label className="block uppercase tracking-wide text-white text-xs font-bold mb-2">
          Name
        </label> 
        <text className='p-2 text-white'>{profile.name}</text>
        <label className="block uppercase tracking-wide text-white text-xs font-bold mt-4 mb-2">
          birthdate
        </label>
        <text className='p-2 text-white'>{profile.birthdate}</text>
        <label className="block uppercase tracking-wide text-white text-xs font-bold mt-4 mb-2">
          Gender
        </label>
        <text className='p-2 text-white'>{profile.gender}</text>
        <label className="block uppercase tracking-wide text-white text-xs font-bold mt-4 mb-2">
          Mobile number
        </label>
        <text className='p-2 text-white'>{profile.phone_number}</text>
        <label className="block uppercase tracking-wide text-white text-xs font-bold mt-4 mb-2">
          Email
        </label>
        <text className='p-2 text-white'>{profile.email}</text>
      </div>
      <Link href='/mfa' className='bg-[#f5f5f5] shadow hover:bg-[#1e244d] hover:text-[#f5f5f5] mt-3 w-full text-[#1e244d] font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300 text-center'>Change Password</Link>
    </div>
    </div>
  }