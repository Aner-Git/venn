import axios from "axios";

const timeout = 5000; //3 seconds
//
class ApiClient {
  constructor() {
    this.axiosClient = axios.create({
      baseURL: window.__REACT_APP_BASE_API_URL__,
      timeout: timeout,
    });
  }

  async request(options) {
    let response = null;
    try {
      response = await this.axiosClient.request({
        url: options.url,
        headers: options.headers,
        method: options.method,
        params: options.params,
        data: options.data,
        responseType: options.responseType,
      });
    } catch (err) {
      throw Error(err.request.responseURL, {
        cause: err,
      });
    }

    return {
      status: response.status,
      body: response.data,
    };
  }

  //params: must be a plain object or a URLSearchParams object
  async get(url, params, options) {
    return this.request({ method: "get", url, params: params, ...options });
  }

  async post(url, data, options) {
    return this.request({ method: "post", url, data, ...options });
  }

  async delete(url, options) {
    return this.request({ method: "delete", url, ...options });
  }

  async put(url, data, options) {
    return this.request({ method: "put", url, data, ...options });
  }

  async patch(url, data, options) {
    return this.request({ method: "patch", url, data, ...options });
  }
}

export default ApiClient;

const api = new ApiClient();
export { api };
