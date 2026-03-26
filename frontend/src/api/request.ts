import axios from "axios";

const request = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

request.interceptors.response.use(
  (response) => {
    const res = response.data;
    if (typeof res?.code !== "number") {
      return Promise.reject(new Error("Unexpected response format"));
    }
    if (res.code !== 0) {
      return Promise.reject(new Error(res.message || "Request failed"));
    }
    return res;
  },
  (error) => {
    const fallback = "Network error";
    const message =
      error?.response?.data?.message ||
      error?.message ||
      fallback;

    return Promise.reject(new Error(message));
  },
);

export default request;
