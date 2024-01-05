import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";

import { IService } from "../../types/Service";
import { formatCurrency } from "../../utils/formatCurrency";

export default function ListServices() {
  const userID = getUserID() || "";
  const [services, setServices] = useState<IService[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const fetchServices = async () => {
    try {
      const response = await api.get(`/services/list/${userID}`, { withCredentials: true });
      setServices(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchServices();
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const filteredServices = services.filter((service) => {
    const name = service.name
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto my-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar serviço</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Nome do serviço" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredServices.length > 0 ? (
          <>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>Nome</th>
                  <th>Preço</th>
                </tr>
              </thead>
              <tbody>
                {filteredServices.map((service) => (
                  <tr key={service.id}>
                    <td>
                      <Link to={`/services/${service.id}`}>{service.name}</Link>
                    </td>
                    <td>{formatCurrency(service.price)}</td>
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
