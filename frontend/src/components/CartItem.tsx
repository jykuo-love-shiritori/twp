import React from 'react'
import { Col, Row } from "react-bootstrap"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from "@fortawesome/free-solid-svg-icons";

import goodsData from '@/pages/goods/goodsData.json';

interface Input {
    item_id: number,
    quantity: number,
    updateTotal?: (subtotal: number, id: number) => void;
    removeItem?: (id: number) => void;
    isCart: boolean;
}

interface Props extends Input {
    name: string,
    imgUrl: string,
    subtotal: number
}

const CartItem = ({ item_id, quantity, updateTotal, removeItem, isCart }: Input) => {
    let data: Props = { item_id, quantity, updateTotal, name: "", imgUrl: "", subtotal: 0, isCart: false };
    const matchingGood = goodsData.find((goods) => goods.id === data.item_id);

    if (matchingGood) {
        data.name = matchingGood.name;
        data.imgUrl = matchingGood.imgUrl;
        data.subtotal = matchingGood.price * data.quantity;
    }

    if (updateTotal) {
        updateTotal(data.subtotal, data.item_id);
    }

    return (
        <div className='cart_item' style={{ margin: '2% 0 2% 0' }}>
            <Row>
                <Col xs={3} md={1} className='center'>
                    <img src={data.imgUrl} />
                </Col>

                <Col xs={9} md={5} className='center cart_font'>
                    {data.name}
                </Col>

                <Col xs={4} md={2} className='right  cart_font'>
                    x{data.quantity}
                </Col>

                <Col xs={4} md={2} className='right  cart_font'>
                    ${data.subtotal}
                </Col>

                <Col xs={4} md={2} className='right'>
                    {isCart && removeItem && (
                        <FontAwesomeIcon icon={faTrash} className='trash' onClick={() => removeItem(data.item_id)} />
                    )}
                </Col>
            </Row>
        </div>
    )
}

export default CartItem
