import { MutationTree } from 'vuex';
import { ExampleStateInterface, User } from './state';

const mutation: MutationTree<ExampleStateInterface> = {
  setUser(state: ExampleStateInterface, user: User) {
    state.user = user;
  },
};

export default mutation;
