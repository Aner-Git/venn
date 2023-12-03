import { ChangeEvent } from "react";
import Form from "react-bootstrap/Form";
type Props = {
  value: number;
  onChange: (v: number) => void;
};

const MaxDistance = ({ value, onChange }: Props) => {
  //validate input
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(e.target.value);
    if (value < 1) {
      return;
    }
    onChange(value);
  };
  return (
    <Form.Group className="mb-3">
      <Form.Label>Max Distance</Form.Label>
      <Form.Control value={value} type="number" onChange={handleChange} />
    </Form.Group>
  );
};

export default MaxDistance;
