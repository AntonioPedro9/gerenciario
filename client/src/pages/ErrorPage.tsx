import { Alert } from "react-bootstrap";

export default function ErrorPage() {
  return (
    <Alert variant="danger" className="mx-auto mt-4" style={{ width: "32rem" }}>
      <strong>Error 404</strong>
      <br />
      Essa rota não existe!
    </Alert>
  );
}
