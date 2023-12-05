import { useState, FormEventHandler } from 'react';
import { useSearchParams } from 'react-router-dom';

interface Token {
  access_token: string;
}

const Callback = () => {
  const [searchParams] = useSearchParams();

  const [code, setCode] = useState(searchParams.get('code') ?? ''); // TODO proper error handling
  const [verifier, setVerifier] = useState('');

  const tokenUrl = '/api/oauth/token';

  const getToken: FormEventHandler = async (e) => {
    e.preventDefault();

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
    console.log(token.access_token);
  };

  return (
    <>
      <form onSubmit={getToken}>
        <label htmlFor='code' style={{ color: 'white' }}>
          Code
        </label>
        <br></br>
        <input type='text' id='code' value={code} onChange={(e) => setCode(e.target.value)} />
        <br></br>
        <label htmlFor='verifier' style={{ color: 'white' }}>
          Verifier
        </label>
        <br></br>
        <input
          type='text'
          id='verifier'
          value={verifier}
          onChange={(e) => setVerifier(e.target.value)}
        />
        <br></br>
        <button type='submit'>Submit</button>
      </form>
    </>
  );
};

export default Callback;
