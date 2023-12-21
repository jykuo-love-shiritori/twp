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
