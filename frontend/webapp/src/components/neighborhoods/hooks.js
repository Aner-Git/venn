import { useQuery } from "react-query";
import { fetchNeighborhoods } from "./queries";

const useNeighborhoodstQuery = (pagination) => {
  const params = { ...pagination };
  return useQuery(["neighborhoods", params], () => fetchNeighborhoods(params));
};

export default useNeighborhoodstQuery;
