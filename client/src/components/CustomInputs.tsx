import { Form, InputGroup } from "react-bootstrap";

interface ICustomInput {
  label: string;
  id: string;
  value: string | undefined;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
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
