import React from 'react'
import { Col, Row } from "react-bootstrap"
import { faUser, faFile } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import History from '@/components/History'
import Link from 'next/link';

const OrderHistory = () => {
    return (
        <div className='flex-wrapper userBG'>
            <Row>
                <Col xs={12} md={12}>
                    <img src='/images/user.jpg' />
                </Col>
                <Col xs={12} md={2}>
                    <Row className='userIcon'>
                        <Col xs={12} className='center'>
                            <img src='/images/head.jpg' className='userImg' />
                        </Col>
                        <Col xs={12} className='center'>
                            <h4>John Jonathan</h4>
                        </Col>
                    </Row>

                    <div className='userButton'>
                        <Link href={"/user/user_info"} className='none'>
                            <FontAwesomeIcon icon={faUser} className='white_word' /> &nbsp; <span className='white_word'>Personal info</span>
                        </Link>
                    </div>
                    <div className='userButton'>
                        <Link href={"/user/order_history"} className='none'>
                            <FontAwesomeIcon icon={faFile} className='white_word' /> &nbsp; <span className='white_word'>Order history</span>
                        </Link>
                    </div>
                </Col>
                <Col xs={12} md={10} style={{ padding: '6% 10% 6% 10%' }}>
                    <History />
                </Col>
            </Row>
        </div>
    )
}

export default OrderHistory
