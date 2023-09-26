import '../../style/global.css';

import { Col, Row } from "react-bootstrap"
import SearchBar from '../../components/SearchBar';
import GoodItem from '../../components/GoodItem';
import goodsData from '../../pages/goods/goodsData.json'

const Goods = () => {
    return (
        <div className='pureBG' style={{ padding: '10%' }}>
            <Row>
                <Col xs={12} md={3}>
                    <span className='titleWhite'>Goods</span>
                </Col>
                <Col md={5} />
                <Col xs={12} md={4} className='center'>
                    <SearchBar />
                </Col>
            </Row>
            <div style={{ padding: '2% 7% 2% 7%' }}>
                <Row>
                    {goodsData.map((data) => {
                        return (
                            <Col xs={12} md={4}>
                                <GoodItem id={data.id} name={data.name} imgUrl={data.imgUrl} />
                            </Col>
                        )
                    })}
                </Row>
            </div>
        </div>
    )
}

export default Goods;
