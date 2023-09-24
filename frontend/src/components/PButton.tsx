import React from 'react'
import '@/components/style.css';
import '@/style/globals.css';
import Link from 'next/link';
import { Button } from "react-bootstrap"

interface Props {
    text: string;
    url: string;
}

const PButton = ({ text, url }: Props) => {
    let button = (
        <Button className="news_button button2_bg">
            &nbsp; {text} &nbsp;
        </Button>
    )

    return (
        <div className='center'>
            {url ? <Link href={url}>{button}</Link> : button}
        </div >
    )
}

export default PButton
