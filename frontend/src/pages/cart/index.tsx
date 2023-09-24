import React, { useState } from 'react'
import BoughtPage from '@/components/BoughtPage';

interface Props {
    item_id: number;
    quantity: number;
    subtotal: number;
}

const Cart = () => {
    let initialData: Props[] = [];
    for (let i = 0; i < 3; i++) {
        const newData: Props = {
            item_id: 2 + i,
            quantity: 3 + 3 * i,
            subtotal: 0
        };
        initialData.push(newData);
    }

    const [cartContainer, setCartContainer] = useState<Props[]>(initialData);

    const [total, setTotal] = useState<number[]>([]);
    const [total_count, setTotalCount] = useState(0);

    const updateTotal = (subtotal: number, id: number) => {
        let isExist = false;
        for (let i = 0; i < total.length; i += 2) {
            if (total[i] === id) {
                isExist = true;
            }
        }

        if (!isExist) {
            setTotal((prevTotal) => [...prevTotal, id, subtotal]);
            setTotalCount((prevTotal) => prevTotal + subtotal);
        }

        for (let i = 0; i < cartContainer.length; i++) {
            if (cartContainer[i].item_id === id) {
                cartContainer[i].subtotal = subtotal;
                break;
            }
        }
    }

    const removeItem = (id: number) => {
        setCartContainer((prevCartContainer) => {
            const item = prevCartContainer.filter(
                (item) => item.item_id === id
            )
            setTotalCount((prevTotal) => prevTotal - item[0].subtotal);

            const updateCartContainer = prevCartContainer.filter(
                (item) => item.item_id !== id
            );
            return updateCartContainer;
        })
    }

    return (
        <BoughtPage title='Cart' cartContainer={cartContainer} total={total_count} updateTotal={updateTotal} removeItem={removeItem} isCart={true} isButtonNeeded={true} />
    )
}

export default Cart;
