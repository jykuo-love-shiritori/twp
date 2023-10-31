import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faCheckCircle,
  faMinusCircle,
  faTimesCircle,
  IconDefinition,
} from '@fortawesome/free-solid-svg-icons';

interface Props {
  icon: IconDefinition;
  text: string;
  date: string | null;
}

const RecordStatus = ({ icon, text, date }: Props) => {
  return (
    <div>
      <div className='center'>
        <div className='status_icon center'>
          <FontAwesomeIcon icon={icon} size='3x' />
        </div>
      </div>
      <h5 style={{ margin: '7% 0 7% 0' }} className='center'>
        {text}
      </h5>
      <div className='center'>
        <FontAwesomeIcon
          icon={date ? (date === 'Confirming' ? faMinusCircle : faCheckCircle) : faTimesCircle}
          size='2x'
          style={{ color: date ? (date === 'Confirming' ? '#FCD265' : 'var(--title)') : '#ED7E6D' }}
        />
      </div>
      <div className='center' style={{ margin: '2% 0 0 0 ', fontSize: '14px' }}>
        <span
          style={{
            color: date ? (date === 'Confirming' ? '#FCD265' : 'var(--title)') : '#ED7E6D',
          }}
        >
          {date}
        </span>
      </div>
    </div>
  );
};

export default RecordStatus;
