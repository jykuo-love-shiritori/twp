import Form from 'react-bootstrap/Form';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import FormItem from '@components/FormItem';

interface Props {
  year: number;
  month: number;
}

const SellerReport = () => {
  const navigate = useNavigate();
  const baseYear = 2020;
  const currentYear = new Date().getFullYear();
  const years = Array.from(
    { length: currentYear - baseYear + 1 },
    (_, index) => currentYear - index,
  );
  const months = Array.from({ length: 12 }, (_, index) => index + 1);

  const { register, handleSubmit } = useForm<Props>({
    defaultValues: {
      year: years[0],
      month: months[0],
    },
  });

  const onSubmit: SubmitHandler<Props> = (data) => {
    navigate(`/user/seller/reports/${data.year}/${data.month}`);
  };

  return (
    <div style={{ padding: '5% 8% 10% 8% ' }}>
      <h3>Select report period</h3>
      <hr className='hr' />

      <form onSubmit={handleSubmit(onSubmit)}>
        <FormItem label='Year'>
          <Form.Select
            {...register('year', {
              setValueAs: (value) => parseInt(value),
            })}
            aria-label='Year'
          >
            {years.map((year) => (
              <option key={year} value={year}>
                {year}
              </option>
            ))}
          </Form.Select>
        </FormItem>

        <FormItem label='Month'>
          <Form.Select
            {...register('month', {
              setValueAs: (value) => parseInt(value),
            })}
            aria-label='Month'
          >
            {months.map((m) => (
              <option key={m} value={m}>
                {m}
              </option>
            ))}
          </Form.Select>
        </FormItem>

        <div className='form_item_wrapper'>
          <input type='submit' value='Confirm' />
        </div>
      </form>
    </div>
  );
};

export default SellerReport;
