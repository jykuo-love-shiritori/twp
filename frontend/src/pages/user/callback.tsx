import { useQuery } from '@tanstack/react-query';
import { Navigate, useSearchParams } from 'react-router-dom';
import { Token } from '@lib/Auth';
import { useContext } from 'react';
import { AuthContext } from '@components/AuthProvider';

const Callback = () => {
  const [searchParams] = useSearchParams();

  const code = searchParams.get('code');
  const state = localStorage.getItem('state');
  const verifier = localStorage.getItem('verifier');

  const { tokenRef } = useContext(AuthContext);

  const getToken = async () => {
    if (searchParams.get('state') !== state) {
      return;
    }

    const tokenUrl = '/api/oauth/token';
    const resp = await fetch(tokenUrl, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: JSON.stringify({
        code: code,
        code_verifier: verifier,
      }),
    });
    return (await resp.json()) as Token;
  };

  const { isPending, error, data, refetch } = useQuery({
    queryKey: ['token'],
    queryFn: getToken,
    enabled: false,
  });

  if (!code || !state || !verifier) {
    console.log('Missing code, state or verifier');
    return <Navigate to='/login' replace={true} />;
  }

  refetch();
  if (isPending) {
    return <>Loading</>;
  }

  if (error) {
    console.log('failed to get token');
    return <Navigate to='/login' replace={true} />;
  }

  if (data) {
    tokenRef.current = data.access_token;
    return <Navigate to='/' replace={true} />;
  }
};

export default Callback;
