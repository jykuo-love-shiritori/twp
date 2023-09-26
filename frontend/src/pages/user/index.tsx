import { Col, Row } from "react-bootstrap"
import { faUser, faFile } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { Outlet } from "react-router-dom"

const User = () => {
    return (
        <div className='flex-wrapper userBG'>
            <Row>
                <Col xs={12} md={12}>
                    <img src='/images/user.jpg' style={{ width: '100%' }} />
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
                        <a href={"/user/user_info"} className='none'>
                            <FontAwesomeIcon icon={faUser} className='white_word' /> &nbsp; <span className='white_word'>Personal info</span>
                        </a>
                    </div>
                    <div className='userButton'>
                        <a href={"/user/order_history"} className='none'>
                            <FontAwesomeIcon icon={faFile} className='white_word' /> &nbsp; <span className='white_word'>Order history</span>
                        </a>
                    </div>

                </Col>
                <Col xs={12} md={10} style={{ padding: '6% 10% 6% 10%' }} >
                    <Outlet />
                </Col>
            </Row>
        </div>
    )
}

export default User
