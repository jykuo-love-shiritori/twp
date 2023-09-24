import BeforeUser from '@/components/BeforeUser'
import React from 'react'
import beforeData from '@/pages/user/before.json'

const Login = () => {
    return (
        <BeforeUser data={beforeData.login} />
    )
}

export default Login
