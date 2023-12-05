import Table from "react-bootstrap/Table";
import {
  flexRender,
  useReactTable,
  createColumnHelper,
  getCoreRowModel,
  getSortedRowModel,
  SortingState,
  OnChangeFn,
} from "@tanstack/react-table";
import Header from "../table/Header";
import PublicTransportAvailability from "./PublicTransportAvailability";
import USDollars from "./USDollars";

type Props = {
  data: Neighborhood[];
  sorting: SortingState;
  onSorting: OnChangeFn<SortingState>;
};

type Neighborhood = {
  neighborhood: string;
  city: string;
  state: string;
  average_age: number;
  distance_from_city_center: number;
  public_transport_availability: string | null;
  average_income: number;
  latitude: number;
  longitude: number;
};
const columnHelper = createColumnHelper<Neighborhood>();

function getColumns() {
  return [
    columnHelper.accessor("neighborhood", {
      header: "Name",
      cell: (props: any) => <>{props.getValue()}</>,
    }),

    columnHelper.accessor("city", {
      header: "City",
      cell: (props: any) => <>{props.getValue()}</>,
    }),
    columnHelper.accessor("state", {
      header: "State",
      cell: (props: any) => <>{props.getValue()}</>,
    }),
    columnHelper.accessor("average_age", {
      header: `Avg age`,
      cell: (props: any) => <>{props.getValue()}</>,
    }),
    columnHelper.accessor("average_income", {
      header: "Avg income",
      cell: (props: any) => <USDollars amount={props.getValue()} />,
    }),
    columnHelper.accessor("distance_from_city_center", {
      header: "Distance from center",
      cell: (props: any) => <>{props.getValue()}</>,
    }),
    columnHelper.accessor("public_transport_availability", {
      header: "Public transport ava",
      cell: (props: any) => (
        <PublicTransportAvailability value={props.getValue()} />
      ),
    }),
    columnHelper.accessor("longitude", {
      header: "Longitude",
      enableSorting: false,
      cell: (props: any) => <>{props.getValue()}</>,
    }),
    columnHelper.accessor("latitude", {
      header: "Latitude",
      enableSorting: false,
      cell: (props: any) => <>{props.getValue()}</>,
    }),
  ];
}

const NeighborhoodTable = ({ data, sorting, onSorting }: Props) => {
  const columns = getColumns();
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    onSortingChange: onSorting,
    state: { sorting },
  });

  return (
    <Table size="sm">
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th key={header.id}>
                {header.isPlaceholder ? null : <Header header={header} />}
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody>
        {table.getRowModel().rows.map((row) => (
          <tr key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default NeighborhoodTable;
