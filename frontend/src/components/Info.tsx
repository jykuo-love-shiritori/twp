import React from 'react'
import PButton from './PButton';
import InfoItem from './InfoItem';

const Info = () => {
    return (
        <div>
            <div className='info_title'>Personal info</div>
            <hr style={{ opacity: '1' }} />
            <InfoItem text="Name" isMore={false} />
            <InfoItem text="Email Address" isMore={false} />
            <InfoItem text="Address" isMore={true} />
            <PButton text='Save' url='' />
        </div>
    )
}

export default Info;
