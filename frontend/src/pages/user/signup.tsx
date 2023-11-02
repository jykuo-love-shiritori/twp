import BeforeUser from '@components/BeforeUser';

import beforeData from '@pages/user/buyer/before.json';

const Signup = () => {
  return <BeforeUser data={beforeData.register} />;
};

export default Signup;
