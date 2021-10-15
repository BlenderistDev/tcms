import {
  User, ContactMap, ChatMap, DialogMap
} from 'src/components/models';
import { MutationTree } from 'vuex';
import { ExampleStateInterface } from './state';

const mutation: MutationTree<ExampleStateInterface> = {
  setUser(state: ExampleStateInterface, user: User) {
    state.user = user;
  },
  setContacts(state: ExampleStateInterface, contacts: ContactMap) {
    state.contacts = contacts;
  },
  setChats(state: ExampleStateInterface, chats: ChatMap) {
    state.chats = chats;
  },
  setDialogs(state: ExampleStateInterface, dialogs: DialogMap) {
    state.dialogs = dialogs;
  },
};

export default mutation;
