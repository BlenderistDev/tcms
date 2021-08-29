import { AxiosResponse } from 'axios';
import { api } from 'src/boot/axios';

export const login = (phone: string):Promise<AxiosResponse> => api.post('/login', { phone });
export const sign = (code: string):Promise<AxiosResponse> => api.post('/sign', { code });
