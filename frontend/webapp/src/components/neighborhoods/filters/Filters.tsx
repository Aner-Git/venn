import { ChangeEvent } from "react";
import { Container } from "react-bootstrap";
import FilterSelect from "./FilterSelect";
import FilterBar from "./FilterBar";
import MaxDistance from "./MaxDistance";
import AgeRange from "./AgeRange";

type Props = { filters: any; dispatch: (e: any) => void };

const Filters = ({ filters, dispatch }: Props) => {
  const handleFilterActiveChange = (e: ChangeEvent<HTMLInputElement>) => {
    dispatch({ type: "active", name: e.target.name, active: e.target.checked });
  };

  const handleMaxDistanceChange = (value: number) => {
    dispatch({ type: "maxdistance", value: value });
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
            min={filters["agerange"].min}
            max={filters["agerange"].max}
          />
        )}
      </FilterBar>
    </Container>
  );
};

export default Filters;
