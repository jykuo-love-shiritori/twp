interface Props {
  searchParams: URLSearchParams;
  setSearchParams: (params: URLSearchParams, options?: { replace?: boolean }) => void;
  refresh: () => void;
  limit?: number;
  maxPage?: number;
  isMore?: boolean;
}

const Pagination = ({
  searchParams,
  setSearchParams,
  refresh,
  limit = 10,
  maxPage = 100,
  isMore = true,
}: Props) => {
  if (!searchParams.has('offset')) {
    searchParams.set('offset', '0');
  }
  if (!searchParams.has('limit') || Number(searchParams.get('limit')) !== limit) {
    searchParams.set('limit', limit.toString());
  }

  const getPage = () => {
    return Number(searchParams.get('offset')) / limit + 1;
  };

  const onPrevious = () => {
    const page = getPage();
    if (page > 1) {
      searchParams.set('offset', ((page - 2) * limit).toString());
      setSearchParams(searchParams, { replace: true });
      refresh();
    }
  };
  const onNext = () => {
    const page = getPage();
    if (page < maxPage && isMore) {
      searchParams.set('offset', (page * limit).toString());
      setSearchParams(searchParams, { replace: true });
      refresh();
    }
  };

  const handlePageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputPage = parseInt(e.target.value);
    if (inputPage > 0 && inputPage < maxPage) {
      if (isMore || (!isMore && inputPage < getPage())) {
        searchParams.set('offset', ((inputPage - 1) * limit).toString());
        setSearchParams(searchParams, { replace: true });
        refresh();
      }
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
      <div className='center'>{maxPage}</div>
      <div className='center' onClick={onNext}>
        {'>'}
      </div>
    </div>
  );
};

export default Pagination;
