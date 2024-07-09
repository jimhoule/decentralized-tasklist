import axios from 'axios';
import type { InternalAxiosRequestConfig } from 'axios';

const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 10000
});

axiosInstance.interceptors.request.use(
    (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
        // Set Bearer Authorization Header here
        // config.headers.Authorization = `Bearer ${accessToken}`

        // Set Locale Header
        //config.headers.locale = 'en';

        return config;
    }
);

export type { AxiosResponse, AxiosInstance } from 'axios';
export { axiosInstance };