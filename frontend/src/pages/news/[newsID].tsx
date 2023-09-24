import React from 'react'
import '@/style/globals.css';
import { useRouter } from 'next/router';
import { Col, Row } from "react-bootstrap"
import newsData from '@/pages/home/newsData.json'
import NotFound from '@/components/NotFound';

interface Props {
    id: number | null;
    imgUrl: string;
    title: string;
    date: string;
    smallTitle: string;
    content: string;
}

const EachNews = () => {
    const router = useRouter();
    const { newsID } = router.query;
    const id = Array.isArray(newsID) ? newsID[0] : newsID;
    const data: Props = { id: null, imgUrl: "", title: "", date: "", smallTitle: "", content: "" };
    const foundNews = newsData.find((news) => news.id.toString() === id);

    if (foundNews) {
        Object.assign(data, foundNews);
    }
    const isNewsExist = !!foundNews;

    if (isNewsExist) {
        return (
            <div className='pureBG' style={{ padding: '10% 10% 0% 10%' }}>
                <div id='news_bg' className='flex-wrapper'>
                    <Row>
                        <Col xs={12} md={4} >
                            <img src={data.imgUrl} />
                        </Col>
                        <Col xs={12} md={8}>
                            <h4 className='title2'>{data.title}</h4> <br />
                            <span className='right'>{data.date}</span>
                            <hr /><br />
                            <p>{data.smallTitle}</p>
                            <p>{data.content}</p>
                        </Col>
                    </Row>
                </div>
            </div>
        )
    }
    else {
        return (
            <NotFound />
        )
    }
}

export default EachNews;
