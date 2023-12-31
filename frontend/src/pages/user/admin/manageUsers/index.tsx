import { Col, Row } from 'react-bootstrap';
import { useMutation } from '@tanstack/react-query';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { CheckMutateStatus, RouteOnNotOK } from '@lib/Status';
import Pagination from '@components/Pagination';
import UserTableRow from '@components/UserTableRow';
import UserTableHeader from '@components/UserTableHeader';
import { useEffect, useState } from 'react';

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
  const [isMore, setIsMore] = useState(true);
  const [searchParams, setSearchParams] = useSearchParams();

  const itemLimit = 10;

  const {
    data: fetchedData,
    status,
    mutate,
  } = useMutation({
    mutationKey: ['adminGetUser'],
    mutationFn: async (isPage: boolean) => {
      if (!isPage) {
        const newSearchParams = new URLSearchParams({
          offset: '0',
          limit: (itemLimit + 1).toString(), // request one more to check if there is more
        });
        setSearchParams(newSearchParams, { replace: true });
      }
      const resp = await fetch('/api/admin/user?' + searchParams.toString(), {
        method: 'GET',
        headers: { accept: 'application/json' },
      });
      RouteOnNotOK(resp, navigate);
      const response = await resp.json();
      if (response.length === itemLimit + 1) {
        response.pop();
        setIsMore(true);
      } else {
        setIsMore(false);
      }
      return response as IUser[];
    },
  });

  const refresh = () => {
    mutate(true);
  };

  // refresh on first render
  useEffect(() => {
    mutate(false);
  }, [mutate]);

  if (status !== 'success') {
    return <CheckMutateStatus status={status} />;
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <div style={{ flexGrow: '9' }}>
        {/* title */}
        <Row>
          <Col md={12} xs={12} className='title'>
            Manage Users
          </Col>
        </Row>
        <UserTableHeader />
        {fetchedData.map((data, index) => (
          <UserTableRow data={data} refresh={refresh} key={index} />
        ))}
      </div>
      <div
        className='center'
        style={{ display: 'flex', flexDirection: 'row', paddingBottom: '10px' }}
      >
        <Pagination
          searchParams={searchParams}
          setSearchParams={setSearchParams}
          refresh={refresh}
          limit={itemLimit}
          isMore={isMore}
        />
      </div>
    </div>
  );
};
export default ManageUser;
