import { useQuery } from '@tanstack/react-query';
import { useContext } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { AuthContext } from '../../components/AuthProvider';

interface Token {
  access_token: string;
}

const Callback = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  // TODO proper error handling
  const code = searchParams.get('code') ?? '';

  const state = localStorage.getItem('state') ?? '';
  const verifier = localStorage.getItem('verifier') ?? '';

  const { setToken } = useContext(AuthContext);

  const tokenUrl = '/api/oauth/token';

  const getToken = async () => {
    if (searchParams.get('state') !== state) {
      return;
    }

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

    const token = (await resp.json()) as Token;
    setToken(token.access_token);

    navigate('/', { replace: true });

    return token;
  };

  const { isPending, error } = useQuery({
    queryKey: ['token'],
    queryFn: getToken,
  });

  if (isPending) {
    return <>Loading</>;
  }

  if (error) {
    return <>error</>;
  }
};

export default Callback;
