import { useReducer, useState } from "react";
import type { SortingState } from "@tanstack/react-table";
import NeighborhoodTable from "./NeighborhoodTable";
import Stack from "react-bootstrap/Stack";
import Spinner from "react-bootstrap/Spinner";
import useOnPage from "../pagination/hooks";
import { Paginator } from "../pagination/Paginator";
import Filters from "./filters/Filters";
import filtersReducer from "./filters/filtersReducer";
import { initialFilters } from "./filters/initialFilters";
import Error from "../errors/Error";
import useNeighborhoodstQuery from "./hooks";
type Props = {};

const NeighborhoodPanel = ({}: Props) => {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [page, handlePrev, handleNext] = useOnPage(1);
  const [filters, dispatch] = useReducer(filtersReducer, initialFilters);

  const { isLoading, isError, isSuccess, data, error } = useNeighborhoodstQuery(
    {
      page,
      pageSize: Paginator.defaultPageSize,
    }
  );

  let content = null;
  if (isLoading) {
    content = <Spinner animation="border" />;
  }

  if (isError) {
    content = <Error msg={`downloading failed: ${(error as Error).message}`} />;
  }

  if (isSuccess) {
    let hoodData = data.body.data;
    let pg = data.body.meta;
    const onPage = Paginator.onPage(pg.page, pg.page_size, pg.total);
    content = (
      <Paginator
        handleNext={handleNext}
        handlePrev={handlePrev}
        onfirst={onPage.first}
        onlast={onPage.last}
      >
        <NeighborhoodTable
          data={hoodData}
          onSorting={setSorting}
          sorting={sorting}
        />
      </Paginator>
    );
  }

  return (
    <Stack gap={4}>
      <Filters filters={filters} dispatch={dispatch} />
      {content}
    </Stack>
  );
};

export default NeighborhoodPanel;
