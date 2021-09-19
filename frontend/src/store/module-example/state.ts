import { ContactMap, User } from 'src/components/models';

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
