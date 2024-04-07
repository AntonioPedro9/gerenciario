import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import formatPhone from "../../utils/formatPhone";

import { ICustomer } from "../../types/Customer";

export default function ListCustomers() {
  const [customers, setCustomers] = useState<ICustomer[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const navigate = useNavigate();

  const fetchCustomers = async () => {
    try {
      const response = await api.get(`/customers/all`, { withCredentials: true });
      setCustomers(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchCustomers();
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const filteredCustomers = customers.filter((customer) => {
    const name = customer.name
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto m-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar cliente</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Nome do cliente" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredCustomers.length > 0 ? (
          <>
            <Table style={{ marginBottom: "0" }} striped bordered hover>
              <thead>
                <tr>
                  <th>Nome</th>
                  <th>Telefone</th>
                </tr>
              </thead>
              <tbody>
                {filteredCustomers.map((customer) => (
                  <tr key={customer.id} onClick={() => navigate(`/customers/${customer.id}`)} style={{ cursor: "pointer" }}>
                    <td>{customer.name}</td>
                    <td style={{ whiteSpace: "nowrap" }}>{formatPhone(customer.phone)}</td>
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
