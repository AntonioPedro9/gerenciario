import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";

import api from "../../service/api";

import { IService } from "../../types/Service";

export default function ServiceDetails() {
  const serviceID = useParams().serviceID;
  const [service, setService] = useState<IService | null>(null);
  const [editableFields, setEditableFields] = useState<Partial<IService>>({});

  console.log(serviceID);

  const fetchServiceData = async () => {
    try {
      const response = await api.get(`/services/${serviceID}`, { withCredentials: true });
      setService(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchServiceData();
  }, []);

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const handleInputChange = (fieldName: keyof IService, value: string) => {
    setEditableFields({
      ...editableFields,
      [fieldName]: value,
    });
  };

  const handleUpdateService = () => {
    if (service && editableFields) {
      const updatedServiceData = {
        id: Number(serviceID),
        userID: service.userID,
        ...editableFields,
      };

      api
        .put("/services/", updatedServiceData, { withCredentials: true })
        .then((response) => {
          setService(response.data);
          setEditableFields({});
          alert("Serviço atualizado com sucesso");
        })
        .catch((error) => console.error(error));
    }
  };

  const handleDeleteService = () => {
    if (service) {
      if (confirm("Tem certeza de que deseja excluir este serviço?")) {
        api
          .delete(`/services/${service.id}`, { withCredentials: true })
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
        <span className="material-symbols-outlined" role="button" onClick={handleDeleteService}>
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {service ? (
          <>
            <Card.Title className="mb-3">Detalhes do serviço</Card.Title>

            <Form>
              <Form.Group className="mb-3" controlId="cpf">
                <Form.Label>
                  Nome <span className="text-red">*</span>
                </Form.Label>
                <Form.Control
                  type="text"
                  name="Name"
                  autoComplete="off"
                  value={editableFields.name || service.name}
                  onChange={(e) => handleInputChange("name", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="name">
                <Form.Label>Descrição</Form.Label>
                <Form.Control
                  type="text"
                  name="Description"
                  autoComplete="off"
                  value={editableFields.description || service.description}
                  onChange={(e) => handleInputChange("description", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="phone">
                <Form.Label>Duração (horas)</Form.Label>
                <Form.Control
                  type="number"
                  name="Duration"
                  autoComplete="off"
                  value={editableFields.duration || service.duration}
                  onChange={(e) => handleInputChange("duration", e.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="phone">
                <Form.Label>Preço</Form.Label>
                <Form.Control
                  type="number"
                  name="Price"
                  autoComplete="off"
                  value={editableFields.price || service.price}
                  onChange={(e) => handleInputChange("price", e.target.value)}
                />
              </Form.Group>

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateService}>
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
