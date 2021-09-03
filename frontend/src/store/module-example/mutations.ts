import { MutationTree } from 'vuex';
import { ContactMap, ExampleStateInterface, User } from './state';

const mutation: MutationTree<ExampleStateInterface> = {
  setUser(state: ExampleStateInterface, user: User) {
    state.user = user;
  },
  setContacts(state: ExampleStateInterface, contacts: ContactMap) {
    state.contacts = contacts;
  },
};

export default mutation;
