import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle, faTimesCircle, IconDefinition } from '@fortawesome/free-solid-svg-icons';

export interface StatusProps {
  icon: IconDefinition;
  text: string;
  status: boolean;
}

const RecordStatus = ({ icon, text, status }: StatusProps) => {
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
          icon={status ? faCheckCircle : faTimesCircle}
          size='2x'
          style={{ color: status ? 'var(--title)' : '#ED7E6D' }}
        />
      </div>
    </div>
  );
};

export default RecordStatus;
