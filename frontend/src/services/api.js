import axios from 'axios';

const api = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL || 'http://api.localdnd.com:7788',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true
});

export default api;