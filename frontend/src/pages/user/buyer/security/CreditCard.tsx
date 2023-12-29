import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCreditCard } from '@fortawesome/free-solid-svg-icons';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { CheckFetchStatus } from '@lib/Status';
import { RouteOnNotOK } from '@lib/Status';
import TButton from '@components/TButton';

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

const ContainerStyle = {
  borderRadius: '24px',
  border: '1px solid var(--button_border, #34977F)',
  background: ' var(--button_dark, #135142)',
  padding: '10% 5% 5% 5%',
  color: 'white',
  marginBottom: '15px',
};

const CreditCard = () => {
  const navigate = useNavigate();

  const { data, status, refetch } = useQuery({
    queryKey: ['userGetCreditCard'],
    queryFn: async () => {
      const resp = await fetch('/api/user/security/credit_card');
      RouteOnNotOK(resp, navigate);
      return await resp.json();
    },
    select: (data) => data as [ICreditCard],
    retry: false,
    enabled: true,
    refetchOnWindowFocus: false,
  });
  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const onDelete = async (id: number) => {
    const newCards = id !== 0 ? data.slice(0, id).concat(data.slice(id + 1)) : data.slice(1);
    const resp = await fetch('/api/user/security/credit_card', {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newCards),
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
    } else {
      refetch();
    }
  };

  return (
    <div>
      <div className='title'>Security - Credit Card</div>
      <hr className='hr' />
      <Row>
        <Col sm={12} md={8}></Col>
        <Col sm={12} md={4}>
          <TButton text='Add New Card' action='/user/security/manageCreditCard/newCard' />
        </Col>
      </Row>
      <br />
      <Row>
        {data.map((card, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <div style={ContainerStyle}>
                <div className='title_color' style={{ padding: '0% 5% 5% 10%' }}>
                  <b>{card.name}</b>
                </div>
                <div className='center'>
                  <FontAwesomeIcon icon={faCreditCard} size='3x' />
                </div>
                <div className='center' style={{ padding: '5%' }}>
                  {`---- ---- ---- ${card.card_number.slice(-4)}`}
                </div>
              </div>
              <TButton text='delete' action={() => onDelete(index)} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default CreditCard;
