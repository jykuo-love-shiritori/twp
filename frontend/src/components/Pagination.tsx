import { useState } from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import { Next } from 'react-bootstrap/esm/PageItem';

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
  const handlePageChange = (e: any) => {
    const inputPage = parseInt(e.target.value);
    if (inputPage > 0 && inputPage <= totalPage) {
      setCurrentPage(inputPage);
    }
  };

  let newPage = currentPage;

  return (
    <div style={{ width: '300px' }}>
      <Container className='pagination center_vertical'>
        <div onClick={onPrevious}>{'<'}</div>
        <div>{'Page:'}</div>
        <input type='text' value={newPage} onChange={handlePageChange} />
        <div>of</div>
        <div>{totalPage}</div>
        <div onClick={onNext}>{'>'}</div>
      </Container>
    </div>
  );
};

export default Pagination;
