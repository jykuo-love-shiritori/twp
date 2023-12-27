import { NavigateFunction } from 'react-router-dom';

export const RouteOnNotOK = (response: Response, navigate: NavigateFunction) => {
  switch (response.status) {
    case 404:
      navigate('/notFound');
      break;
    // deal with redirect here (maybe)
  }
};

type FetchStatusProps = {
  status: 'pending' | 'error' | 'success';
};
export const CheckFetchStatus = ({ status }: FetchStatusProps) => {
  switch (status) {
    case 'pending':
      return <div>Loading...</div>;
    case 'error':
      return <div>Error...</div>;
  }
};
