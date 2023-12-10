import { useQuery } from '@tanstack/react-query';
import { Navigate, useSearchParams } from 'react-router-dom';
import { Token } from '@lib/Auth';
import { useContext } from 'react';
import { AuthContext } from '@components/AuthProvider';

const Callback = () => {
  const [searchParams] = useSearchParams();

  // TODO proper error handling
  const code = searchParams.get('code') ?? '';

  const state = localStorage.getItem('state') ?? '';
  const verifier = localStorage.getItem('verifier') ?? '';

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

  const { isPending, error, data } = useQuery({
    queryKey: ['token'],
    queryFn: getToken,
  });

  if (isPending) {
    return <>Loading</>;
  }

  if (error) {
    return <>error</>;
  }

  if (data) {
    tokenRef.current = data.access_token;
    return <Navigate to='/' replace={true} />;
  }
};

export default Callback;
