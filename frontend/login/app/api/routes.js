import 'dotenv/config'

const baseURL = process.env.BASE_URL

async function register(company = "ascenda", email, first_name, last_name, birthdate, password) {
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
    const res_json = await res.json()
    console.log(res_json)
    return res_json
}

async function verifyEmail(otp) {
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
    const res_json = await res.json()
    console.log(res_json)
    return res_json
}

async function login(otp) {
    const res = await fetch(baseURL + "hosted/login", {
        method: "POST",
        headers : {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            otp: otp
        })
    })
    const res_json = await res.json()
    console.log(res_json)
    return res_json
}
