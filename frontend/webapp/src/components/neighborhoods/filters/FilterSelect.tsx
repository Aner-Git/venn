import Form from "react-bootstrap/Form";
import { ChangeEvent } from "react";

type Filter = {
  name: string;
  id: string;
};

type Props = {
  filters: Filter[];
  onClick: (e: ChangeEvent<HTMLInputElement>) => void;
};

const FilterSelect = ({ filters, onClick }: Props) => {
  return (
    <Form>
      {filters.map((f) => (
        <Form.Check
          onClick={onClick as any}
          inline
          label={`${f.name}`}
          name={`${f.id}`}
          type="checkbox"
          id={`${f.id}`}
          key={`${f.id}`}
        />
      ))}
    </Form>
  );
};

export default FilterSelect;
