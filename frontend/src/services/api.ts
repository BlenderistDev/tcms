import { AxiosResponse } from 'axios';
import { api } from 'src/boot/axios';
import { User } from 'src/store/module-example/state';

export const login = (phone: string):Promise<AxiosResponse> => api.post('/login', { phone });
export const sign = (code: string):Promise<AxiosResponse> => api.post('/sign', { code });

export interface user {
    id: number
}

export const getUser = () => api.get<User>('/me').then((response) => response.data);
