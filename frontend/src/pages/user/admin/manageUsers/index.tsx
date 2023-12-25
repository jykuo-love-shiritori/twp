import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { RouteOnNotOK } from '@lib/Functions';
import { CheckFetchStatus } from '@lib/Status';
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
  const [searchParams, setSearchParams] = useSearchParams();
  const {
    data: fetchedData,
    status,
    refetch,
  } = useQuery({
    queryKey: ['adminGetUser'],
    queryFn: async () => {
      console.log('/api/admin/user?' + searchParams.toString());
      const resp = await fetch('/api/admin/user?' + searchParams.toString(), {
        method: 'GET',
        headers: { accept: 'application/json' },
      });
      RouteOnNotOK(resp, navigate);
      const response = [] as IUser[];
      return response;
    },
    select: (data) => data as IUser[],
    enabled: true,
    refetchOnWindowFocus: false,
  });

  const refresh = () => {
    refetch();
    // navigate(0);
  };

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
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
        />
      </div>
    </div>
  );
};
export default ManageUser;
