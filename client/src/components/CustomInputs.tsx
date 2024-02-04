import { Form, InputGroup } from "react-bootstrap";
import { IClient } from "../types/Client";
import { IService } from "../types/Service";

interface ICustomInput {
  label: string;
  id: string;
  value?: string | number | undefined;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
}

interface ISelectInput {
  label: string;
  id: string;
  required?: boolean;
  options: (IClient | IService)[];
  onSelect?: (value: any) => void;
}

export function TextInput({ label, id, value, onChange, required }: ICustomInput) {
  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>
      <Form.Control type="text" name={id} autoComplete="off" onChange={onChange} value={value} />
    </Form.Group>
  );
}

export function NumberInput({ label, id, value, onChange, required }: ICustomInput) {
  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>
      <Form.Control type="number" name={id} autoComplete="off" onChange={onChange} value={value} />
    </Form.Group>
  );
}

export function EmailInput({ label, id, value, onChange, required }: ICustomInput) {
  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>
      <Form.Control type="email" name={id} autoComplete="off" onChange={onChange} value={value} />
    </Form.Group>
  );
}

export function PasswordInput({ label, id, value, onChange, required }: ICustomInput) {
  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>
      <Form.Control type="password" name={id} autoComplete="off" onChange={onChange} value={value} />
    </Form.Group>
  );
}

export function PhoneInput({ label, id, value, onChange, required }: ICustomInput) {
  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>
      <InputGroup>
        <InputGroup.Text>+55</InputGroup.Text>
        <Form.Control type="tel" name={id} autoComplete="off" onChange={onChange} value={value} />
      </InputGroup>
    </Form.Group>
  );
}

export function SelectInput({ label, id, required, options, onSelect }: ISelectInput) {
  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = event.target.value;
    const selectedOption = options.find((option) => option.id.toString() === selectedId);

    if (onSelect && selectedOption) {
      onSelect(selectedOption);
    }
  };

  return (
    <Form.Group className="mb-3" controlId={id}>
      <Form.Label>
        {label}
        {required ? <span className="text-red"> *</span> : null}
      </Form.Label>

      <Form.Select aria-label="Default select example" onChange={handleChange}>
        <option>Selecionar</option>

        {options.map((option, index) => (
          <option key={index} value={option.id}>
            {option.name}
          </option>
        ))}
      </Form.Select>
    </Form.Group>
  );
}
