import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import { MutationCache, QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AuthProvider from '@components/AuthProvider.tsx';

const queryClient = new QueryClient({
  mutationCache: new MutationCache({
    onError: (error: Error) => {
      alert(error);
    },
  }),
});

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <App />
      </AuthProvider>
    </QueryClientProvider>
  </React.StrictMode>,
);
