import { ChangeEvent } from "react";
import Form from "react-bootstrap/Form";

type Props = {
  value: number;
  onChange: (v: number) => void;
};

const MaxDistance = ({ value, onChange }: Props) => {
  //validate input
  const handleChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const value = parseInt(e.target.value);
    onChange(value);
  };
  return (
    <Form.Group className="mb-3">
      <Form.Label>Max Distance</Form.Label>
      <Form.Select value={value} onChange={handleChange}>
        <option value="1">1</option>
        <option value="2">2</option>
        <option value="5">5</option>
        <option value="10">10</option>
        <option value="15">15</option>
        <option value="20">20</option>
        <option value="25">25</option>
        <option value="30">30</option>
        <option value="35">35</option>
        <option value="40">40</option>
      </Form.Select>
    </Form.Group>
  );
};
export default MaxDistance;
