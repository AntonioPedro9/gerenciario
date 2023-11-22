import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";
import formatPhone from "../../utils/formatPhone";

import { IClient } from "../../types/Client";

export default function ListClients() {
  const userID = getUserID() || "";
  const [clients, setClients] = useState<IClient[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    api
      .get(`/clients/list/${userID}`, { withCredentials: true })
      .then((response) => setClients(response.data))
      .catch((error) => console.error(error));
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const filteredClients = clients.filter((client) => {
    const name = client.name
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto my-4" style={{ width: "32rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar cliente</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Nome do cliente" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredClients.length > 0 ? (
          <>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>Nome</th>
                  <th>Telefone</th>
                </tr>
              </thead>
              <tbody>
                {filteredClients.map((client) => (
                  <tr key={client.id}>
                    <td>
                      <Link to={`/clients/${client.id}`}>{client.name}</Link>
                    </td>
                    <td>{formatPhone(client.phone)}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </>
        ) : null}
      </Card.Body>
    </Card>
  );
}
