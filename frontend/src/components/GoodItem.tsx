import React from 'react'
import '@/components/style.css';
import '@/style/globals.css';
import PButton from './PButton';

interface Props {
    id: number;
    name: string;
    imgUrl: string;
}

const GoodItem = ({ id, name, imgUrl }: Props) => {
    let url = `/goods/${id}`

    return (
        <div className='goods_item'>
            <img src={imgUrl} style={{ borderRadius: '0 0 30px 0' }} />
            <div className='goods_title_c' style={{ padding: '2% 7% 2% 7% ' }}>
                {name}
            </div>
            <div className='center'>
                <PButton text='more' url={url} />
            </div>
        </div>
    )
}

export default GoodItem;
