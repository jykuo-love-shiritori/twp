import { useState } from 'react';

interface Props {
  currentPageInit: number;
  totalPage: number;
}

const Pagination = ({ currentPageInit = 1, totalPage }: Props) => {
  const [currentPage, setCurrentPage] = useState(currentPageInit);

  const onPrevious = () => {
    if (currentPage - 1 > 0) {
      setCurrentPage(currentPage - 1);
    }
  };
  const onNext = () => {
    if (currentPage + 1 <= totalPage) {
      setCurrentPage(currentPage + 1);
    }
  };
  const handlePageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputPage = parseInt(e.target.value);
    if (inputPage > 0 && inputPage <= totalPage) {
      setCurrentPage(inputPage);
    }
  };
  return (
    <div className='pagination center_vertical center'>
      <div className='center' onClick={onPrevious}>
        {'<'}
      </div>
      <div className='center'>{'Page:'}</div>
      <input className='center' type='text' value={currentPage} onChange={handlePageChange} />
      <div className='center'>of</div>
      <div className='center'>{totalPage}</div>
      <div className='center' onClick={onNext}>
        {'>'}
      </div>
    </div>
  );
};

export default Pagination;
