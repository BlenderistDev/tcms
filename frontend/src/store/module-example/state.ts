export interface User {
  id: number,
  phone: string,
  name: string
}

export interface Contact {
  id: number
  username: string,
  phone: string,
  firstName: string,
  lastName: string,
  mutualContact: boolean,
  bot: boolean,
  deleted: boolean,
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
