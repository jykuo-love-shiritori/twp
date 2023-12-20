import { useState } from 'react';
import Form from 'react-bootstrap/Form';

import TButton from '@components/TButton';
import FormItem from '@components/FormItem';

const SellerReport = () => {
  const currentYear = new Date().getFullYear();
  const years = Array.from({ length: currentYear - 2019 }, (_, index) => 2020 + index);
  const months = Array.from({ length: 12 }, (_, index) => index + 1);

  const [year, setYear] = useState(years[0]);
  const [month, setMonth] = useState(months[0]);

  return (
    <div style={{ padding: '5% 8% 10% 8% ' }}>
      <h3>Select report period</h3>
      <hr className='hr' />

      <FormItem label='Year'>
        <Form.Select
          aria-label='Year'
          value={year}
          onChange={(e) => setYear(parseInt(e.target.value))}
        >
          <option value=''>Select year</option>
          {years.map((y) => (
            <option key={y} value={y}>
              {y}
            </option>
          ))}
        </Form.Select>
      </FormItem>

      <FormItem label='Month'>
        <Form.Select
          aria-label='Month'
          value={month}
          onChange={(e) => setMonth(parseInt(e.target.value))}
        >
          <option>Select month</option>
          {months.map((m) => (
            <option key={m} value={m}>
              {m}
            </option>
          ))}
        </Form.Select>
      </FormItem>
      <TButton text='Confirm' action={`/user/seller/reports/${year}/${month}`} />
    </div>
  );
};

export default SellerReport;
