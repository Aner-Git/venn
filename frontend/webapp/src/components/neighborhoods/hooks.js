import { useQuery } from "react-query";
import { fetchNeighborhoods } from "./queries";

const useNeighborhoodstQuery = (pagination, sortBy, filters) => {
  let sort = JSON.stringify(sortBy);
  const params = { ...pagination, sort, ...filters };
  return useQuery(
    ["neighborhoods", params],
    () => fetchNeighborhoods(params),
    {}
  );
};

export default useNeighborhoodstQuery;
