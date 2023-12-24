import { NavigateFunction } from 'react-router-dom';

export const RouteOnNotOK = (response: Response, navigate: NavigateFunction) => {
  switch (response.status) {
    case 404:
      navigate('/notFound');
      break;
    // deal with redirect here (maybe)
  }
};
