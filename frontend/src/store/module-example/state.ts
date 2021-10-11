import { ChatMap, ContactMap, User } from 'src/components/models';

export interface ExampleStateInterface {
  user: User|null;
  contacts: ContactMap|null;
  chats: ChatMap|null;
}

function state(): ExampleStateInterface {
  return {
    user: null,
    contacts: null,
    chats: null,
  };
}

export default state;
