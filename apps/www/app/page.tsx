'use client';

import { api } from '@/api-client';

export default function Index() {
  const { mutate } = api.users.useCreateUser();

  const handleClick = () => {
    console.log('ttrrr');
    mutate(
      { data: { name: 'test' } },
      {
        onSuccess: (data) => {
          alert(data.status);
        },
        onError: (data) => {
          alert(data.message);
        },
      }
    );
  };

  return (
    <div>
      <button onClick={handleClick}>create user</button>
    </div>
  );
}
