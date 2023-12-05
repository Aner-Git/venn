import { api } from "../../ApiClient";

export const fetchNeighborhoods = async (params: any) => {
  return await api.get(`/neighborhoods`, params);
};
