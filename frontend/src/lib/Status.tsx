import { NavigateFunction } from 'react-router-dom';

export const RouteOnNotOK = async (response: Response, navigate: NavigateFunction) => {
  let res;
  switch (response.status) {
    case 400:
      res = await response.json();
      alert(res.message);
      break;
    case 401:
      navigate('/forbidden');
      break;
    case 403:
      navigate('/unauthorized');
      break;
    case 404:
      navigate('/notFound');
      break;
    case 500:
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
