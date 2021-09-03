import { AxiosResponse } from 'axios';
import { api } from 'src/boot/axios';
import { ContactMap, User } from 'src/store/module-example/state';

export const login = (phone: string):Promise<AxiosResponse> => api.post('/login', { phone });
export const sign = (code: string):Promise<AxiosResponse> => api.post('/sign', { code });
export const getUser = () => api.get<User>('/me').then((response) => response.data);
export const getContacts = () => api.get<ContactMap>('/contacts').then((response) => response.data);
