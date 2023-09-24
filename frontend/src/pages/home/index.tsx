import React from 'react'
import '@/style/globals.css';
import { Col, Row } from "react-bootstrap"
import News from '@/components/News';

import newsData from '@/pages/home/newsData.json'

const Home = () => {
    return (
        <div>
            <img src='images/home.png' className='home_pic'></img>

            <span id='home_page_title'>Pastel</span>
            <span id='home_page_content'>
                Explore the magic, colorful delights,<br />just a click away.
            </span>
            <div id='home_page_back1' />
            <div id='home_page_back2' />

            <div id='home_news_bg' />

            <div style={{ padding: '1% 15% 1% 15% ' }}>
                <h2 id='home_news_title'>News</h2>

                <Row>
                    {newsData.map((data) => {
                        return (
                            <Col xs={12} md={4}>
                                <News id={data.id} imgUrl={data.imgUrl} title={data.title} />
                            </Col>
                        )
                    })}
                </Row>
            </div>

            <div id='home_sign' className='center'><div id='home_sign_content'>Sign up now</div></div>
        </div>
    )
}

export default Home
