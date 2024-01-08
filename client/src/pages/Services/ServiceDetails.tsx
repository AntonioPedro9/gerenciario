import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput, NumberInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { IService, IUpdateServiceRequest } from "../../types/Service";

export default function ServiceDetails() {
  const userID = useParams().userID;
  const [service, setService] = useState<IService | null>(null);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [duration, setDuration] = useState<number | "">(0);
  const [price, setPrice] = useState<number | "">(0);

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const fetchServiceData = async () => {
    try {
      const response = await api.get(`/services/${userID}`, { withCredentials: true });
      const serviceData = response.data;

      setService(serviceData);
      setName(serviceData.name);
      setDescription(serviceData.description);
      setDuration(serviceData.duration);
      setPrice(serviceData.price);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchServiceData();
  }, []);

  const handleTextChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleNumberChange = (setter: React.Dispatch<React.SetStateAction<number | "">>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value === "" ? "" : Number(e.target.value));
  };

  const handleUpdateService = async () => {
    if (service) {
      const updatedServiceData: IUpdateServiceRequest = {
        id: Number(userID),
        name,
        description,
        duration: Number(duration),
        price: Number(price),
        userID: service.userID,
      };

      try {
        const response = await api.put("/services/", updatedServiceData, { withCredentials: true });
        setService(response.data);
        if (response.status === 200) goBack();
      } catch (error: any) {
        alert(error.response.data.error);
      }
    }
  };

  const handleDeleteService = async () => {
    if (service && confirm("Tem certeza de que deseja excluir este serviço?")) {
      try {
        const response = await api.delete(`/services/${userID}`, { withCredentials: true });
        if (response.status === 204) goBack();
      } catch (error: any) {
        alert(error.response.data.error);
      }
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
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
              <TextInput label="Nome" id="name" value={name} onChange={handleTextChange(setName)} required />
              <TextInput label="Descrição" id="description" value={description} onChange={handleTextChange(setDescription)} />
              <NumberInput label="Duração (horas)" id="duration" value={duration} onChange={handleNumberChange(setDuration)} />
              <NumberInput label="Preço" id="price" value={price} onChange={handleNumberChange(setPrice)} />

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateService}>
                Salvar alterações
              </Button>
            </Form>
          </>
        ) : (
          <p>Carregando...</p>
        )}
      </Card.Body>
    </Card>
  );
}
