import axios, {
  type AxiosRequestTransformer,
  type AxiosResponseTransformer,
} from "axios";
import { toCamelCase, toSnakeCase } from "./caseTransform";

const normalizeTransform = <T,>(transform: T | T[] | undefined): T[] => {
  if (!transform) {
    return [];
  }
  return Array.isArray(transform) ? transform : [transform];
};

const shouldTransformRequest = (data: unknown): boolean => {
  if (data === null || data === undefined) {
    return false;
  }
  if (typeof FormData !== "undefined" && data instanceof FormData) {
    return false;
  }
  if (typeof URLSearchParams !== "undefined" && data instanceof URLSearchParams) {
    return false;
  }
  if (typeof Blob !== "undefined" && data instanceof Blob) {
    return false;
  }
  if (typeof ArrayBuffer !== "undefined" && data instanceof ArrayBuffer) {
    return false;
  }
  return typeof data === "object";
};

const requestTransforms: AxiosRequestTransformer[] = [
  (data) => (shouldTransformRequest(data) ? toSnakeCase(data) : data),
  ...normalizeTransform<AxiosRequestTransformer>(axios.defaults.transformRequest),
];

const responseTransforms: AxiosResponseTransformer[] = [
  ...normalizeTransform<AxiosResponseTransformer>(
    axios.defaults.transformResponse
  ),
  (data) => toCamelCase(data),
];

const api = axios.create({
  baseURL: "http://localhost:8081",
  headers: {
    "Content-Type": "application/json",
  },
  transformRequest: requestTransforms,
  transformResponse: responseTransforms,
});

// Add a request interceptor to include the token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token"); // Assuming you store the token in localStorage
    if (token) {
      console.log("token", token);
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export default api;
