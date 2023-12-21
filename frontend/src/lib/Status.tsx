import { NavigateFunction } from 'react-router-dom';

export const RouteOnNotOK = (responce: Response, navigate: NavigateFunction) => {
  switch (responce.status) {
    case 404:
      navigate('/notFound');
      break;
    // deal with redirect here (maybe)
  }
};
