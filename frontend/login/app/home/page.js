'use client'
import { useEffect } from 'react'
import { useRouter } from 'next/navigation';
import Home from '../components/home'


export default function HomePage(props) { 
  const router = useRouter()
  //useEffect(() => {
  //  const fetchData = async () => {
  //   const res = await fetch("https://client.itsag2t1.com/api/" + "hosted/user"
  //    , {
  //      method: "GET",
  //      credentials: "include",
  //      headers : {
  //          'Accept': 'application/json',
  //          'Content-Type': 'application/json'
  //      }
   //   })
    //  const data = await res.json()
  //    if (res.status != 200) {
  //      router.push("/login")
  //    }
  //  }
  //  fetchData().catch((err) => {
  //    router.push("/login")
  //  })
//}, [])


  return <Home />
}