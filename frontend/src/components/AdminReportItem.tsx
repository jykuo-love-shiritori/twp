import '@style/global.css';

import TButton from '@components/TButton';

import adminReportData from '@pages/user/admin/adminReportData.json';

interface Sellers {
  id: number;
  name: string;
  imgUrl: string;
  amount: number;
}

interface Input {
  year: number;
  month: number;
}

interface Props extends Input {
  sellers: Sellers[];
  totalAmount: number;
  id: number;
}

const AdminReportItem = ({ year, month }: Input) => {
  const reportItemStyle = {
    borderRadius: '24px',
    border: '1px solid var(--button_border, #34977F)',
    background: 'var(--button_dark, #135142)',
    padding: '10%',
    margin: '0 0 10px 0',
  };

  const Months = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
  ];

  let data: Props | undefined;
  adminReportData.findIndex((item) => {
    if (item.year === year && item.month === month) {
      data = item;
    }
  });

  if (data) {
    return (
      <div>
        <div style={reportItemStyle}>
          <div className='center' style={{ fontSize: '48px' }}>
            <b>{data.month}</b>
          </div>
          <div className='center' style={{ margin: '0' }}>
            <h5>
              <b>{data.year}</b>
            </h5>
          </div>
          <div className='center title_color' style={{ paddingTop: '40px' }}>
            {Months[data.month - 1]}
          </div>
        </div>
        <TButton text='more' url={`/admin/reports/${data.id}`} />
      </div>
    );
  }
};

export default AdminReportItem;
