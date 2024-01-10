import { NavigateFunction } from 'react-router-dom';

type NavigateType = null | NavigateFunction;
export const RouteOnNotOK = async (resp: Response, navigate: NavigateType = null) => {
  if (!navigate) {
    if (!resp.ok) {
      alert((await resp.json()).message);
    }
    return;
  }
  switch (resp.status) {
    case 400:
      navigate('/notFound');
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

type MutateStatusProps = {
  status: 'error' | 'idle' | 'pending';
};
export const CheckMutateStatus = ({ status }: MutateStatusProps) => {
  switch (status) {
    case 'pending':
      return <div>Loading...</div>;
    case 'error':
      return <div>Error...</div>;
  }
};
