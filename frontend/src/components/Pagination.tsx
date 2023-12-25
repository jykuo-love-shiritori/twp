interface Props {
  searchParams: URLSearchParams;
  setSearchParams: (params: URLSearchParams, options?: { replace?: boolean }) => void;
  refresh: () => void;
}

const MAX_PAGE = 100;

const Pagination = ({ searchParams, setSearchParams, refresh }: Props) => {
  if (!searchParams.has('offset')) {
    searchParams.set('offset', '0');
  }
  if (!searchParams.has('limit')) {
    searchParams.set('limit', '10');
  }

  const getPage = () => {
    return Number(searchParams.get('offset')) / Number(searchParams.get('limit')) + 1;
  };

  const onPrevious = () => {
    const page = getPage();
    if (page > 1) {
      searchParams.set('offset', ((page - 2) * Number(searchParams.get('limit'))).toString());
      setSearchParams(searchParams, { replace: true });
      refresh();
    }
  };
  const onNext = () => {
    const page = getPage();
    if (page < MAX_PAGE) {
      searchParams.set('offset', (page * Number(searchParams.get('limit'))).toString());
      setSearchParams(searchParams, { replace: true });
      refresh();
    }
  };

  const handlePageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputPage = parseInt(e.target.value);
    if (inputPage > 0 && inputPage < MAX_PAGE) {
      searchParams.set('offset', ((inputPage - 1) * Number(searchParams.get('limit'))).toString());
      setSearchParams(searchParams, { replace: true });
      refresh();
    }
  };
  return (
    <div className='pagination center_vertical center'>
      <div className='center' onClick={onPrevious}>
        {'<'}
      </div>
      <div className='center'>{'Page:'}</div>
      <input className='center' type='text' value={getPage()} onChange={handlePageChange} />
      <div className='center'>of</div>
      <div className='center'>{MAX_PAGE}</div>
      <div className='center' onClick={onNext}>
        {'>'}
      </div>
    </div>
  );
};

export default Pagination;
