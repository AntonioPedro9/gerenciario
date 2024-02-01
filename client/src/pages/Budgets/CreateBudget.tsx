import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form, Row, Col, Badge, Button } from "react-bootstrap";

import { NumberInput, SelectInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";

import { IClient } from "../../types/Client";
import { IService } from "../../types/Service";

export default function CreateBudget() {
  const userID = getUserID() || "";
  const [price, setPrice] = useState<number | "">(0);
  const [selectedService, setSelectedService] = useState<IService | null>(null);
  const [servicesList, setServicesList] = useState<IService[]>([]);

  const [clients, setClients] = useState<IClient[]>([]);
  const [services, setServices] = useState<IService[]>([]);

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const fetchClients = async () => {
    try {
      const response = await api.get(`/clients/list/${userID}`, { withCredentials: true });
      setClients(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  const fetchServices = async () => {
    try {
      const response = await api.get(`/services/list/${userID}`, { withCredentials: true });
      setServices(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchClients();
    fetchServices();
  }, []);

  const handleNumberChange = (setter: React.Dispatch<React.SetStateAction<number | "">>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value === "" ? "" : Number(e.target.value));
  };

  const handleAddService = () => {
    if (selectedService) {
      setServicesList([...servicesList, selectedService]);
      setSelectedService(null);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar orçamento</Card.Title>

        <Form>
          <SelectInput label="Cliente" id="client" options={clients} required />
          <SelectInput label="Serviços" id="service" options={services} onSelect={setSelectedService} required />

          <Button className="mb-3" variant="light" size="sm" onClick={handleAddService}>
            Adicionar serviço ao orçamento +
          </Button>

          {servicesList.length !== 0 ? (
            <div className="mb-3">
              {servicesList.map((service, index) => (
                <Badge key={index} bg="secondary">
                  {service.name}
                </Badge>
              ))}
            </div>
          ) : null}

          <NumberInput label="Preço total" id="price" value={price} onChange={handleNumberChange(setPrice)} required />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
