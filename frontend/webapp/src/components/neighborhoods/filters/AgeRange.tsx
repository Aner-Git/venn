import Form from "react-bootstrap/Form";
import { ChangeEvent } from "react";
type Field = {
  name: string;
  value: number;
};

type Props = {
  min: number;
  max: number;
  onChange: (e: Field) => void;
};

const AgeRange = ({ min, max, onChange }: Props) => {
  //    const handleChange = (e:ChangeEvent<HTMLInputElement>)=>{ onChange({name:e.target.name, value:e.target.value})
  const handleChange = () => {};

  return (
    <Form.Group className="mb-3">
      <Form.Label>Age Range - Min</Form.Label>
      <Form.Control
        value={min}
        onChange={handleChange}
        name="min"
        type="number"
      />
      <Form.Label>Age Range - Max </Form.Label>
      <Form.Control
        value={max}
        onChange={handleChange}
        name="max"
        type="number"
      />
    </Form.Group>
  );
};

export default AgeRange;
