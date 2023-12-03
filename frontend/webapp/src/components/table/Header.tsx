import { Header as HeaderT, flexRender } from "@tanstack/react-table";
import ArrowUpDown from "../icons/ArrowUpDown";

type Props = { header: HeaderT<any, any> };

const Header = ({ header }: Props) => {
  return (
    <>
      {flexRender(header.column.columnDef.header, header.getContext())}
      {header.column.getCanSort() && (
        <div
          className="mx-2 d-inline-block TableSorter"
          onClick={header.column.getToggleSortingHandler()}
        >
          <ArrowUpDown />
        </div>
      )}
    </>
  );
};

export default Header;
