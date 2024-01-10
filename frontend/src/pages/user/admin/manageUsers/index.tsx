import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { CheckMutateStatus, RouteOnNotOK } from '@lib/Status';
import { useState } from 'react';
import { useAuth } from '@lib/Auth';
import Pagination from '@components/Pagination';
import UserTableRow from '@components/UserTableRow';
import UserTableHeader from '@components/UserTableHeader';

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

interface IUser {
  address: string;
  credit_card: [ICreditCard];
  email: string;
  enabled: true;
  icon_url: string;
  name: string;
  role: string;
  username: string;
}

const ManageUser = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const [isMore, setIsMore] = useState(true);
  const [searchParams, setSearchParams] = useSearchParams();

  const itemLimit = 10;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }
  const { data: fetchedData, status } = useQuery({
    queryKey: ['adminGetUser', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch('/api/admin/user?' + searchParams.toString(), {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          accept: 'application/json',
        },
      });
      RouteOnNotOK(resp, navigate);
      const response = await resp.json();
      if (response.length === itemLimit + 1) {
        response.pop();
        setIsMore(true);
      } else {
        setIsMore(false);
      }
      return response;
    },
    select: (data) => data as IUser[],
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckMutateStatus status={status} />;
  }

  return (
    <div>
      <Row>
        <Col md={12} xs={12} className='title'>
          Manage Users
        </Col>
      </Row>
      <UserTableHeader />
      {fetchedData.map((data, index) => (
        <UserTableRow data={data} key={index} />
      ))}
      <div className='center'>
        <Pagination limit={itemLimit} isMore={isMore} />
      </div>
    </div>
  );
};
export default ManageUser;
