import { Button } from "react-bootstrap";

interface ICustomInput {
  text: string;
}

export function SubmitButton({ text }: ICustomInput) {
  return (
    <Button variant="dark" type="submit" style={{ width: "100%" }}>
      {text}
    </Button>
  );
}
