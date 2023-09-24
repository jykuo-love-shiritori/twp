import React from 'react'
import { Col, Row } from "react-bootstrap"
import HistoryItem from './HistoryItem'

import historyData from '@/pages/cart/boughtData.json'

const History = () => {
    return (
        <div>
            <div className='info_title'>Order history</div>
            <hr style={{ opacity: '1' }} /><br />

            <Row >
                {historyData.map((item) => {
                    return (
                        <Col xs={12} md={4}>
                            <HistoryItem id={item.recordID} />
                        </Col>
                    )
                })}
            </Row>
        </div>
    )
}

export default History
