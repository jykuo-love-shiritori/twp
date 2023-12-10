import { useState } from 'react';

import TButton from '@components/TButton';
import FormItem from '@components/FormItem';

const SellerReport = () => {
  const now = new Date();
  const [year, setYear] = useState(now.getFullYear());
  const [month, setMonth] = useState(now.getMonth());
  return (
    <div style={{ padding: '5% 8% 10% 8% ' }}>
      <h3>At what time would you like to view the report?</h3>
      <hr className='hr' />
      <FormItem label='Year'>
        <input
          type='text'
          defaultValue={year}
          onChange={(e) => {
            setYear(parseInt(e.target.value));
          }}
        />
      </FormItem>
      <FormItem label='Month'>
        <input
          type='text'
          defaultValue={month}
          onChange={(e) => {
            setMonth(parseInt(e.target.value));
          }}
        />
      </FormItem>
      <TButton text='Confirm' action={`/user/seller/reports/${year}/${month}`} />
    </div>
  );
};

export default SellerReport;
