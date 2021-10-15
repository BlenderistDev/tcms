import {
  ChatMap, ContactMap, DialogMap, User
} from 'src/components/models';

export interface ExampleStateInterface {
  user: User|null;
  contacts: ContactMap|null;
  chats: ChatMap|null;
  dialogs: DialogMap|null;
}

function state(): ExampleStateInterface {
  return {
    user: null,
    contacts: null,
    chats: null,
    dialogs: null,
  };
}

export default state;
