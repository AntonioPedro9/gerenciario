import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput, EmailInput, PhoneInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { IClient, IUpdateClientRequest } from "../../types/Client";

export default function ClientDetails() {
  const clientID = useParams().clientID;
  const [client, setClient] = useState<IClient | null>(null);
  const [editableFields, setEditableFields] = useState<Partial<IUpdateClientRequest>>({});

  const navigate = useNavigate();
  const goBack = () => navigate("/clients/list");

  const fetchClientData = async () => {
    try {
      const response = await api.get(`/clients/${clientID}`, { withCredentials: true });
      setClient(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchClientData();
  }, []);

  const handleInputChange = (fieldName: keyof IUpdateClientRequest, event: React.ChangeEvent<HTMLInputElement>) => {
    setEditableFields({
      ...editableFields,
      [fieldName]: event.target.value,
    });
  };

  const handleUpdateClient = async () => {
    if (client && editableFields) {
      const updatedClientData = {
        id: Number(clientID),
        userID: client.userID,
        ...editableFields,
      };

      try {
        const response = await api.put("/clients/", updatedClientData, { withCredentials: true });

        setClient(response.data);
        setEditableFields({});

        if (response.status === 200) {
          alert("Cliente atualizado com sucesso");
          goBack();
        }
      } catch (error: any) {
        console.error(error.response.data.error);
      }
    }
  };

  const handleDeleteClient = async () => {
    if (client && confirm("Tem certeza de que deseja excluir este cliente?")) {
      try {
        const response = await api.delete(`/clients/${clientID}`, { withCredentials: true });

        if (response.status === 204) {
          goBack();
        }
      } catch (error: any) {
        console.error(error.response.data.error);
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
              <TextInput
                label="CPF"
                id="cpf"
                value={editableFields.cpf !== undefined ? editableFields.cpf : client.cpf}
                onChange={(event) => handleInputChange("cpf", event)}
                required
              />
              <TextInput
                label="Nome"
                id="name"
                value={editableFields.name !== undefined ? editableFields.name : client.name}
                onChange={(event) => handleInputChange("name", event)}
                required
              />
              <EmailInput
                label="Email"
                id="email"
                value={editableFields.email !== undefined ? editableFields.email : client.email}
                onChange={(event) => handleInputChange("email", event)}
              />
              <PhoneInput
                label="Telefone"
                id="phone"
                value={editableFields.phone !== undefined ? editableFields.phone : client.phone}
                onChange={(event) => handleInputChange("phone", event)}
                required
              />

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
