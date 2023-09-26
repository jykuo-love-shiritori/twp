import { Col, Row } from "react-bootstrap"
import Footer from "./Footer";

interface Props {
    imgUrl: string;
    title: string;
    content: string;
    subContent: string;
    way: string;
    path: string;
    url: string;
}

interface Data {
    data: Props | undefined;
}

const BeforeUser = ({ data }: Data) => {
    if (data) {
        return (
            <div>
                <div className='center flex-wrapper before_user'>
                    <Row style={{ backgroundColor: 'white' }}>
                        <Col xs={12} md={6}>
                            <img src={data?.imgUrl} style={{ height: '100%', width: '100%' }} />
                        </Col>
                        <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
                            <div className='before_title'>{data?.title}</div>
                            <div className='before_content'>
                                {data?.content}
                            </div>

                            <div className='center' style={{ margin: '5%' }}>
                                <div className='forTest' />

                            </div>

                            <div className='before_subcontent center' >
                                {data?.subContent}
                            </div>
                            <br />
                            <Row>
                                <Col xs={4}>
                                    <hr />
                                </Col>
                                <Col xs={4} className='center'>
                                    Or With
                                </Col>
                                <Col xs={4}>
                                    <hr />
                                </Col>
                            </Row>

                            <div className='center'>
                                {data?.way} &nbsp; <span><u><a href={data?.url}>{data?.path}</a></u></span>
                            </div>

                        </Col>
                    </Row>
                </div>
                <Footer />
            </div>

        )

    }
}

export default BeforeUser
