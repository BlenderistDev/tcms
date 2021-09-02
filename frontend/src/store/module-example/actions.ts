import { getUser } from 'src/services/api';
import { ActionTree } from 'vuex';
import { StateInterface } from '../index';
import { ExampleStateInterface } from './state';

const actions: ActionTree<ExampleStateInterface, StateInterface> = {
  fetchUser({ commit }) {
    void getUser().then((user) => {
      commit('setUser', user);
    });
  },
};

export default actions;
