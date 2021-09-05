export interface User {
  id: number,
  phone: string,
  name: string
}

export interface Contact {
  id: number
  username: string,
  phone: string,
  bot: boolean,
  firstName: string,
  lastName: string,
}

export type ContactMap = {
  [key: number]: Contact
}

export interface ExampleStateInterface {
  user: User|null;
  contacts: ContactMap|null
}

function state(): ExampleStateInterface {
  return {
    user: null,
    contacts: null,
  };
}

export default state;
