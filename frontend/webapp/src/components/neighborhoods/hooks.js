import { useQuery } from "react-query";
import { fetchNeighborhoods } from "./queries";

const useNeighborhoodstQuery = (pagination, sortBy, enabled) => {
  let sort = JSON.stringify(sortBy);
  const params = { ...pagination, sort };
  return useQuery(["neighborhoods", params], () => fetchNeighborhoods(params), {
    enabled: enabled,
  });
};

export default useNeighborhoodstQuery;
