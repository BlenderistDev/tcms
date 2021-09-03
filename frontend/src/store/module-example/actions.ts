import { getContacts, getUser } from 'src/services/api';
import { ActionTree } from 'vuex';
import { StateInterface } from '../index';
import { ExampleStateInterface } from './state';

const actions: ActionTree<ExampleStateInterface, StateInterface> = {
  fetchUser({ commit }) {
    void getUser().then((user) => {
      commit('setUser', user);
    });
  },
  fetchContacts({ commit }) {
    void getContacts().then((contacts) => {
      commit('setContacts', contacts);
    });
  },
};

export default actions;
