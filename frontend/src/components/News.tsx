import React from 'react'
import '@/components/style.css';
import '@/style/globals.css';

import { Button } from "react-bootstrap"
import Link from 'next/link';

interface Props {
    id: number;
    imgUrl: string;
    title: string;
}

const News = ({ id, imgUrl, title }: Props) => {
    return (
        <div>
            <img src={imgUrl} className='new_pic_c' />

            <div style={{ padding: '1% 15% 1% 15%' }}>
                <span className='news_title_c '>{title}</span>
            </div>

            <div className="center">
                <Link href={`/news/${id}`}>
                    <Button className="news_button">
                        &nbsp; more &nbsp;
                    </Button>
                </Link>
            </div>
            <br />
        </div>
    )
}

export default News