import React from 'react'
import { Col, Row } from "react-bootstrap"

interface Props {
    text: string;
    isMore: boolean;
}

const InfoItem = ({ text, isMore }: Props) => {
    return (
        <Row style={{ margin: '2% 0% 2% 0% ' }}>
            <Col xs={12} md={4} className='info_text'>
                {text}
            </Col>
            <Col xs={12} md={8}>
                {!isMore ? (
                    <input
                        type="text"
                        placeholder={text}
                        className='inputBox'
                    />
                ) : (
                    <textarea
                        placeholder={text}
                        className='inputBox'
                    />
                )}
            </Col>
        </Row>
    )
}

export default InfoItem
