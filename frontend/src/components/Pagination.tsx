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
    <div className='pagination center_vertical'>
      <div onClick={onPrevious}>{'<'}</div>
      <div>{'Page:'}</div>
      <input type='text' value={currentPage} onChange={handlePageChange} />
      <div>of</div>
      <div>{totalPage}</div>
      <div onClick={onNext}>{'>'}</div>
    </div>
  );
};

export default Pagination;
