type Props = {
  status: 'pending' | 'error' | 'success';
};

export const CheckStatus = ({ status }: Props) => {
  switch (status) {
    case 'pending':
      return <div>Loading...</div>;
    case 'error':
      return <div>Error...</div>;
  }
};
