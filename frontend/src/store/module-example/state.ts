export interface User {
  id: number,
  phone: string,
  name: string
}

export interface ExampleStateInterface {
  user: User|null;
}

function state(): ExampleStateInterface {
  return {
    user: null,
  };
}

export default state;
