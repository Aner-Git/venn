import { ChangeEvent } from "react";
import { Button, Container } from "react-bootstrap";
import FilterSelect from "./FilterSelect";
import FilterBar from "./FilterBar";
import MaxDistance from "./MaxDistance";
import AgeRange, { Field } from "./AgeRange";

type Props = {
  filters: any;
  dispatch: (e: any) => void;
  onFilterQuery: () => void;
  filtersActive: boolean;
};

const Filters = ({
  filters,
  dispatch,
  onFilterQuery,
  filtersActive,
}: Props) => {
  const handleFilterActiveChange = (e: ChangeEvent<HTMLInputElement>) => {
    dispatch({ type: "active", name: e.target.name, active: e.target.checked });
  };

  const handleMaxDistanceChange = (value: number) => {
    dispatch({ type: "maxdistance", value: value });
  };

  const handleAgeRangeChange = (e: Field) => {
    dispatch({ ...e, type: "agerange" });
  };

  return (
    <Container>
      Filters:
      <FilterSelect
        filters={Object.values(filters)}
        onClick={handleFilterActiveChange}
      />
      <FilterBar gap={2}>
        {filters["maxdistance"].active && (
          <MaxDistance
            value={filters["maxdistance"].value}
            onChange={handleMaxDistanceChange}
          />
        )}
        {filters["agerange"].active && (
          <AgeRange
            onChange={handleAgeRangeChange}
            min={filters["agerange"].min}
            max={filters["agerange"].max}
          />
        )}
      </FilterBar>
      {filtersActive && <Button onClick={onFilterQuery}>Filter</Button>}
    </Container>
  );
};

export default Filters;
