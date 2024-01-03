import { Col, Row } from 'react-bootstrap';

import TButton from '@components/TButton';
import UserItem from '@components/UserItem';

export interface BuyerHistoryItemProps {
  user: 'buyer';
  id: number;
  shop_name: string;
  shop_image_url: string;
  thumbnail_url: string;
  product_name: string;
  shipment: number;
  total_price: number;
  status: 'paid' | 'shipped' | 'delivered';
  created_at: string;
}

export interface SellerHistoryItemProps {
  user: 'seller';
  id: number;
  product_name: string;
  thumbnail_url: string;
  user_name: string;
  user_image_url: string;
  shipment: number;
  total_price: number;
  status: 'paid' | 'shipped' | 'delivered';
  created_at: string;
}

interface Props {
  data: BuyerHistoryItemProps | SellerHistoryItemProps;
}

const HistoryItem = ({ data }: Props) => {
  return (
    <div>
      <Row className='history_container dark'>
        <Col sm={12} md={6}>
          <UserItem
            img_path={data.user === 'buyer' ? data.shop_image_url : data.user_image_url}
            name={data.user === 'buyer' ? data.shop_name : data.user_name}
          />
        </Col>
        <Col sm={12} md={6} className='right'>
          Record ID : {data.id}
        </Col>
        <Col sm={12} md={12}>
          <hr
            style={{
              color: 'var(--button_dark)',
              opacity: '1',
              margin: '10px 0px 10px 0px',
              width: '100%',
            }}
          />
        </Col>
        <Col xs={4} md={2} lg={1} className='center'>
          <img src={data.thumbnail_url} style={{ width: '100%', borderRadius: '10px' }} />
        </Col>
        <Col xs={8} md={8} lg={9} className='center_vertical'>
          <Row style={{ width: '100%' }}>
            <Col xs={12} lg={6}>
              {data.product_name}
            </Col>
            <Col xs={12} lg={6}>
              Order Total : ${data.total_price}
            </Col>
          </Row>
        </Col>
        <Col xs={12} md={2} className='right'>
          <TButton
            text='Detail'
            action={
              data.user === 'buyer'
                ? `/user/buyer/order/${data.id}`
                : `/user/seller/order/${data.id}`
            }
          />
        </Col>
      </Row>
    </div>
  );
};

export default HistoryItem;
