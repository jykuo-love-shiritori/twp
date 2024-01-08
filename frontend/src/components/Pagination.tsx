// import { useState } from 'react';
import { useForm } from 'react-hook-form';

interface Props {
  searchParams: URLSearchParams;
  setSearchParams: (params: URLSearchParams, options?: { replace?: boolean }) => void;
  refetch: () => void;
  limit?: number;
  isMore?: boolean;
}

const Pagination = ({
  searchParams,
  setSearchParams,
  refetch,
  limit = 10,
  isMore = true,
}: Props) => {
  // const [MAX_PAGE, setMaxPage] = useState(100);
  const MAX_PAGE = 100;

  if (!searchParams.has('offset')) {
    searchParams.set('offset', '0');
  }
  if (!searchParams.has('limit') || Number(searchParams.get('limit')) !== limit + 1) {
    // request one more to check if there is more
    searchParams.set('limit', (limit + 1).toString());
  }
  const getPage = () => {
    return Number(searchParams.get('offset')) / limit + 1;
  };

  // set the real max page when it ever get there (this is a cheap bug fix :P)
  // if (MAX_PAGE !== getPage() && !isMore) {
  //   console.log('set max page');
  //   setMaxPage(getPage());
  // }

  const { register, handleSubmit, setValue } = useForm<{ newPage: number }>({
    defaultValues: { newPage: getPage() },
  });

  const onPrevious = () => {
    const page = getPage();
    if (page > 1) {
      searchParams.set('offset', ((page - 2) * limit).toString());
      setSearchParams(searchParams, { replace: true });
      setValue('newPage', page - 1);
      refetch();
    }
  };
  const onNext = () => {
    const page = getPage();
    if (page < MAX_PAGE && isMore) {
      searchParams.set('offset', (page * limit).toString());
      setSearchParams(searchParams, { replace: true });
      setValue('newPage', page + 1);
      refetch();
    }
  };

  const onSubmit = (data: { newPage: number }) => {
    const inputPage = data.newPage;
    if (inputPage > 0 && inputPage < MAX_PAGE && (isMore || (!isMore && inputPage < getPage()))) {
      searchParams.set('offset', ((inputPage - 1) * limit).toString());
      setSearchParams(searchParams, { replace: true });
      refetch();
    } else {
      setValue('newPage', getPage());
    }
  };
  return (
    <div className='pagination center_vertical center' style={{ padding: '5px' }}>
      <div className='center' onClick={onPrevious}>
        {'<'}
      </div>
      <div className='center'>{'Page: '}</div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          className='center'
          {...register('newPage')}
          style={{
            maxWidth: '50px',
            border: 'none',
            textAlign: 'center',
            borderRadius: '10px',
            fontWeight: 'bold',
            color: 'var(--button_light)',
          }}
        />
      </form>
      <div className='center' onClick={onNext}>
        {'>'}
      </div>
    </div>
  );
};

export default Pagination;
