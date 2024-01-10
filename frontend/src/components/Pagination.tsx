import { useSearchParams } from 'react-router-dom';

interface Props {
  limit?: number;
  isMore?: boolean;
}

const Pagination = ({ limit = 10, isMore = true }: Props) => {
  const [searchParams, setSearchParams] = useSearchParams();

  if (!searchParams.has('offset')) {
    searchParams.set('offset', '0');
  }
  if (Number(searchParams.get('limit')) !== limit + 1) {
    searchParams.set('limit', (limit + 1).toString());
  }
  const getPage = () => {
    return Number(searchParams.get('offset')) / limit + 1;
  };

  const onPrevious = () => {
    const page = getPage();
    if (page > 1) {
      searchParams.set('offset', ((page - 2) * limit).toString());
      setSearchParams(searchParams, { replace: true });
    }
  };
  const onNext = () => {
    const page = getPage();
    if (isMore) {
      searchParams.set('offset', (page * limit).toString());
      setSearchParams(searchParams, { replace: true });
    }
  };

  return (
    <div className='pagination center_vertical center' style={{ padding: '5px' }}>
      <div className='center' onClick={onPrevious}>
        {'<'}
      </div>
      <div className='center'>{'Page: '}</div>
      <div
        className='center'
        style={{
          width: '40px',
          textAlign: 'center',
          borderRadius: '10px',
          fontWeight: 'bold',
          background: 'white',
          color: 'var(--button_light)',
        }}
      >
        {getPage()}
      </div>
      <div className='center' onClick={onNext}>
        {'>'}
      </div>
    </div>
  );
};

export default Pagination;
