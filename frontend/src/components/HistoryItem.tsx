import React from 'react'
import PButton from './PButton'
import historyData from '@/pages/cart/boughtData.json'
import goodsData from '@/pages/goods/goodsData.json';

const HistoryItem = ({ id }: { id: number }) => {

    const record = historyData.find((item) => item.recordID === id);
    const firstItem = goodsData.find((item) => record?.items[0].item_id === item.id);

    let Total = 0;
    if (record != undefined) {
        Total = record.items.reduce((accumulator, item) => {
            const foundItem = goodsData.find((goods) => goods.id === item.item_id);

            if (foundItem) {
                const subtotal = item.quantity * foundItem.price;
                accumulator += subtotal;
            }
            return accumulator;
        }, 0);
    }

    const url = `/user/history/${id}`

    return (
        <div className='history_container'>
            <img src={firstItem?.imgUrl} />
            <div>{record?.items.length} items</div>
            <div>${Total}</div>
            <div className='right'>{record?.time}</div>
            <PButton text='more' url={url} />
        </div>
    )
}

export default HistoryItem
