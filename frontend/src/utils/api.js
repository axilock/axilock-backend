import axios from "axios";

// Create an axios instance with default configurations
const apiClient = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL || "https://api.axilock.ai/v1/"  ,
  headers: {
    "Content-Type": "application/json",
  },
});

// Interceptors for request/response (optional)
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("auth-token");
    if (token) {
      config.headers.Authorization = `bearer ${token}`; // Add token to headers
    }
    // config.headers["Content-Type"] = "Application/json"
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

apiClient.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    // console.error("API Error:", error.response || error.message);
    
    if(error.code === "ERR_NETWORK"){ 
      return Promise.reject({"message":"Network Error! Service Down, Please try again later."});
    }
    if(error.response.status === 401){
      localStorage.removeItem("auth-token");
    }
    return Promise.reject(error);
  }
);

export default apiClient;
