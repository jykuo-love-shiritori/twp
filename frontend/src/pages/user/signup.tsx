import React from 'react'
import BeforeUser from '@/components/BeforeUser'
import beforeData from '@/pages/user/before.json'

const Signup = () => {
    return (
        <BeforeUser data={beforeData.register} />
    )
}

export default Signup
