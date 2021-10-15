import {
  getChats, getContacts, getDialogs, getUser
} from 'src/services/api';
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
  fetchChats({ commit }) {
    void getChats().then((chats) => {
      commit('setChats', chats);
    });
  },
  fetchDialogs({ commit }) {
    void getDialogs().then((dialogs) => {
      commit('setDialogs', dialogs);
    });
  },
};

export default actions;
