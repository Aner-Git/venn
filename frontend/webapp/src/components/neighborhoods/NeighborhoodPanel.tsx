import { useState } from "react";
import type { SortingState } from "@tanstack/react-table";
import NeighborhoodTable from "./NeighborhoodTable";
import useOnPage from "../pagination/hooks";
import { Paginator } from "../pagination/Paginator";

type Props = {};

const hoodData = [
  {
    neigborhood: "Hallows",
    city: "Los Dos Caminos",
    average_age: 58,
    distance_from_city_center: 43.5,
    average_income: 44438,
    public_transport_availability: "high",
    latitude: 10.4980067,
    longitude: -66.8335096,
  },
  {
    neigborhood: "Stang",
    state: "B7",
    city: "Angoulême",
    average_age: 20,
    distance_from_city_center: 66.1,
    average_income: 92866,
    public_transport_availability: "low",
    latitude: 45.6258385,
    longitude: 0.0629891,
  },
  {
    neigborhood: "Rockefeller",
    city: "Nanganga",
    average_age: 64,
    distance_from_city_center: 67.6,
    average_income: 32437,
    public_transport_availability: null,
    latitude: -10.3931514,
    longitude: 39.1361568,
  },
];

const NeighborhoodPanel = ({}: Props) => {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [page, handlePrev, handleNext] = useOnPage(1);

  console.log(sorting);
  //come from server
  const pg = { page: 1, page_size: 25, total: 3 };
  const onPage = Paginator.onPage(pg.page, pg.page_size, pg.total);
  return (
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
};

export default NeighborhoodPanel;
