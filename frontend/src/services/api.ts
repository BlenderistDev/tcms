import { AxiosResponse } from 'axios';
import _ from 'lodash';
import { api } from 'src/boot/axios';
import {
  User, ContactMap, Contact, ChatMap, DialogMap
} from 'src/components/models';

export const login = (phone: string):Promise<AxiosResponse> => api.post('/login', { phone });
export const sign = (code: string):Promise<AxiosResponse> => api.post('/sign', { code });
export const getUser = () => api.get<User>('/me').then((response) => response.data);
export const getContacts = () => api.get<ContactMap>('/contacts')
  .then((response) => response.data)
  .then((data) => _.mapValues(data,
    (contact) => _.mapKeys(contact, (_value, key) => _.camelCase(key))));

export const getChats = () => api.get<ChatMap>('/chats')
  .then((response) => response.data)
  .then((data) => _.mapValues(data,
    (chat) => _.mapKeys(chat, (_value, key) => _.camelCase(key))));

export const getDialogs = () => api.get<DialogMap>('/dialogs')
  .then((response) => response.data)
  .then((data) => _.mapValues(data,
    (dialog) => _.mapKeys(dialog, (_value, key) => _.camelCase(key))));

export type SendMessage = {
  recipient: Contact | null,
  message: string,
}

export const sendMessage = (message: SendMessage) => {
  if (message.recipient === null) {
    throw new Error('Undefined recipient');
  }
  return api.post<void>('/message', {
    accessHash: message.recipient.accessHash,
    id: message.recipient.id,
    message: message.message,
  });
};
