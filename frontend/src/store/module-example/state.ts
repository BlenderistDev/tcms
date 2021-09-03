export interface User {
  id: number,
  phone: string,
  name: string
}

export interface Contact {
  Id: number
  Username: string,
  Phone: string,
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
