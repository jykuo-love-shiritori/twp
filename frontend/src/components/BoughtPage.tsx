import React from 'react'
import CartItem from './CartItem';
import PButton from './PButton';

interface Props {
    item_id: number;
    quantity: number;
}

interface Input {
    title: string;
    cartContainer: Props[];
    total: number;
    updateTotal?: (subtotal: number, id: number) => void;
    removeItem?: (id: number) => void;
    isCart: boolean;
    isButtonNeeded: boolean;
}

const BoughtPage = ({ title, cartContainer, total, updateTotal, removeItem, isCart, isButtonNeeded }: Input) => {

    return (
        <div className='bg flex-wrapper'>
            <span className='titleWhite'>{title}</span>
            {cartContainer.map((data) => {
                return (
                    <CartItem item_id={data.item_id} quantity={data.quantity} removeItem={removeItem} updateTotal={updateTotal} isCart={isCart} />
                )
            })}
            <hr className='white_bg' />
            <div className='center title2'>Total : ${total}</div> <br />
            {isButtonNeeded ? <PButton text='Submit' url='' /> : ""}

        </div>
    )
}

export default BoughtPage
