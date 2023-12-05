import { useQuery } from "react-query";
import { fetchNeighborhoods } from "./queries";

const useNeighborhoodstQuery = (pagination, sortBy) => {
  let sort = JSON.stringify(sortBy);
  const params = { ...pagination, sort };
  return useQuery(["neighborhoods", params], () => fetchNeighborhoods(params));
};

export default useNeighborhoodstQuery;
