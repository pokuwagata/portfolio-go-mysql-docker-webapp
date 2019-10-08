import * as React from 'react';

export const Hoge = () => {
  console.log('fuga');
  // const [hoge, setHoge] = React.useState('hoge');
  React.useEffect(() => {
    console.log('hoge');
  }, []);
  return <p>hoge</p>;
};
