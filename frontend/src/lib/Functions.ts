import { NavigateFunction } from 'react-router-dom';

export const RouteOnNotOK = async (responce: Response, navigate: NavigateFunction) => {
  let res;
  switch (responce.status) {
    case 400:
      res = await responce.json();
      alert(res.message);
      break;
    case 404:
      navigate('/notFound');
      break;
    // deal with redirect here (maybe)
  }
};
