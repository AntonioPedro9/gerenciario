import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, InputGroup, Button } from "react-bootstrap";

import api from "../../service/api";

import { IClient } from "../../types/Client";

export default function ClientDetails() {
  const clientID = useParams().clientID;
  const [client, setClient] = useState<IClient | null>(null);
  const [editableFields, setEditableFields] = useState<Partial<IClient>>({});

  const fetchClientData = async () => {
    try {
      const response = await api.get(`/clients/${clientID}`, { withCredentials: true });
      setClient(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchClientData();
  }, []);

  const navigate = useNavigate();
  const goBack = () => navigate("/clients/list");

  const handleInputChange = (fieldName: keyof IClient, value: string) => {
    setEditableFields({
      ...editableFields,
      [fieldName]: value,
    });
  };

  const handleUpdateClient = () => {
    if (client && editableFields) {
      const updatedClientData = {
        id: Number(clientID),
        userID: client.userID,
        ...editableFields,
      };

      api
        .put("/clients/", updatedClientData, { withCredentials: true })
        .then((response) => {
          setClient(response.data);
          setEditableFields({});
          alert("Cliente atualizado com sucesso");
        })
        .catch((error) => console.error(error));
    }
  };

  const handleDeleteClient = () => {
    if (client) {
      if (confirm("Tem certeza de que deseja excluir este cliente?")) {
        api
          .delete(`/clients/${client.id}`, { withCredentials: true })
          .then(() => goBack())
          .catch((error) => console.error(error));
      }
    }
  };

  return (
    <Card className="mx-auto my-4" style={{ width: "32rem" }}>
      <Card.Header className="d-flex justify-content-between align-items-center">
        <span className="material-symbols-outlined" role="button" onClick={goBack}>
          arrow_back
        </span>
        <span className="material-symbols-outlined" role="button" onClick={handleDeleteClient}>
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {client ? (
          <>
            <Card.Title className="mb-3">Detalhes do cliente</Card.Title>

            <Form>
              <Form.Group className="mb-3" controlId="cpf">
                <Form.Label>
                  CPF <span className="text-red">*</span>
                </Form.Label>
                <Form.Control
                  type="text"
                  name="CPF"
                  autoComplete="off"
                  value={editableFields.cpf || client.cpf}
                  onChange={(e) => handleInputChange("cpf", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="name">
                <Form.Label>
                  Nome <span className="text-red">*</span>
                </Form.Label>
                <Form.Control
                  type="text"
                  name="Name"
                  autoComplete="off"
                  value={editableFields.name || client.name}
                  onChange={(e) => handleInputChange("name", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="email">
                <Form.Label>Email</Form.Label>
                <Form.Control
                  type="email"
                  name="Email"
                  autoComplete="off"
                  value={editableFields.email || client.email}
                  onChange={(e) => handleInputChange("email", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="phone">
                <Form.Label>
                  Telefone <span className="text-red">*</span>
                </Form.Label>
                <InputGroup>
                  <InputGroup.Text>+55</InputGroup.Text>
                  <Form.Control
                    type="text"
                    name="Phone"
                    autoComplete="off"
                    value={editableFields.phone || client.phone}
                    onChange={(e) => handleInputChange("phone", e.target.value)}
                  />
                </InputGroup>
              </Form.Group>

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateClient}>
                Salvar alterações
              </Button>
            </Form>
          </>
        ) : (
          <div>Carregando...</div>
        )}
      </Card.Body>
    </Card>
  );
}
