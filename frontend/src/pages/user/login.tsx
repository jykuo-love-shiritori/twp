import BeforeUser from '@components/BeforeUser';

import beforeData from '@pages/user/buyer/before.json';

const Login = () => {
  return <BeforeUser data={beforeData.login} />;
};

export default Login;
