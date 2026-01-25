import axios, { type AxiosInstance } from "axios";

const attachAuthInterceptor = (instance: AxiosInstance): void => {
  instance.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem("token"); // Assuming you store the token in localStorage
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );
};

const createApi = (): AxiosInstance =>
  axios.create({
    baseURL: "http://localhost:8081",
    headers: {
      "Content-Type": "application/json",
    },
  });

const api = createApi();
attachAuthInterceptor(api);

export const apiRaw = createApi();
export default api;
