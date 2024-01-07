import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { IService, IUpdateServiceRequest } from "../../types/Service";

export default function ServiceDetails() {
  const serviceID = useParams().serviceID;
  const [service, setService] = useState<IService | null>(null);
  const [editableFields, setEditableFields] = useState<Partial<IUpdateServiceRequest>>({});

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const fetchServiceData = async () => {
    try {
      const response = await api.get(`/services/${serviceID}`, { withCredentials: true });
      setService(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchServiceData();
  }, []);

  const handleInputChange = (fieldName: keyof IUpdateServiceRequest, event: React.ChangeEvent<HTMLInputElement>) => {
    let value: string | number = event.target.value;

    if (fieldName === "duration" || fieldName === "price") {
      value = Number(value);
    }

    setEditableFields({
      ...editableFields,
      [fieldName]: value,
    });
  };

  const handleUpdateService = async () => {
    if (service && editableFields) {
      const updatedServiceData = {
        id: Number(serviceID),
        userID: service.userID,
        ...editableFields,
      };

      try {
        const response = await api.put("/services/", updatedServiceData, { withCredentials: true });

        setService(response.data);
        setEditableFields({});

        if (response.status === 200) {
          goBack();
        } else {
          alert("Falha ao atualizar serviço");
        }
      } catch (error: any) {
        console.error(error.response.data.error);
      }
    }
  };

  const handleDeleteService = async () => {
    if (service && confirm("Tem certeza de que deseja excluir este serviço?")) {
      try {
        const response = await api.delete(`/services/${serviceID}`, { withCredentials: true });

        if (response.status === 204) {
          goBack();
        }
      } catch (error: any) {
        console.error(error.response.data.error);
      }
    }
  };

  return (
    <Card className="mx-auto my-4" style={{ width: "24rem" }}>
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
              <TextInput
                label="Nome"
                id="name"
                value={editableFields.name !== undefined ? editableFields.name : service.name}
                onChange={(event) => handleInputChange("name", event)}
                required
              />
              <TextInput
                label="Descrição"
                id="description"
                value={editableFields.description !== undefined ? editableFields.description : service.description}
                onChange={(event) => handleInputChange("description", event)}
              />
              <TextInput
                label="Duração (horas)"
                id="duration"
                value={editableFields.duration !== undefined ? editableFields.duration : service.duration}
                onChange={(event) => handleInputChange("duration", event)}
              />
              <TextInput
                label="Preço"
                id="price"
                value={editableFields.price !== undefined ? editableFields.price : service.price}
                onChange={(event) => handleInputChange("price", event)}
              />

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
