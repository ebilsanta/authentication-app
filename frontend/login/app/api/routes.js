'use client'

import 'dotenv/config'

// const baseURL = process.env.REACT_APP_BASE_URL
const baseURL = "https://client.itsag2t1.com/api/"

async function register(router, email, first_name, last_name, birthdate, password, company = "ascenda") {
    const res = await fetch(baseURL + "hosted/register", {
        method: "POST",
        headers : {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            company: company,
            email: email,
            first_name: first_name, 
            last_name: last_name,
            birthdate: birthdate,
            password: password
        })
    })
    const data = await res.json()
    if (res.status != 201) {
        console.log(data)
        return null
    }
    console.log(data)
    // return { props: { data } }
    router.push("/otp")
}

async function verifyEmail(router, otp) {
    const res = await fetch(baseURL + "hosted/verify-email", {
        method: "POST",
        headers : {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            otp: otp
        })
    })
    const data = await res.json()
    if (res.status != 200) {
        console.log(data)
        return
    }
    console.log(data)
    // return { props: { data } }
    router.push("/login")
}

async function login(router, email, password, company = "ascenda") {
    const res = await fetch(baseURL + "hosted/login", {
        method: "POST",
        headers : {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            company: company,
            email: email,
            password: password
        })
    })

    const data = await res.json()
    console.log(data)
    if (res.status === 200) {
      router.push("/home")
    } else {
        console.log("Login Failed")
      return
    }
}

export { register, verifyEmail, login }
