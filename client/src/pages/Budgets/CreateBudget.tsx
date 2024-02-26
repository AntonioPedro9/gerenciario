import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form, Badge, Button } from "react-bootstrap";

import { NumberInput, SelectInput, TextInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";
import { budgetSchema } from "../../utils/validations";

import { IClient } from "../../types/Client";
import { IService } from "../../types/Service";
import { ICreateBudgetRequest } from "../../types/Budget";

export default function CreateBudget() {
  const userID = getUserID() || "";
  const [price, setPrice] = useState<number>(0);
  const [client, setClient] = useState<IClient>();
  const [vehicle, setVehicle] = useState("");
  const [licensePlate, setLicensePlate] = useState("");
  const [selectedService, setSelectedService] = useState<IService>();
  const [servicesList, setServicesList] = useState<IService[]>([]);
  const [clients, setClients] = useState<IClient[]>([]);
  const [services, setServices] = useState<IService[]>([]);

  const navigate = useNavigate();
  const goBack = () => navigate("/budgets/list");

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

  const handleTextChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleClientChange = (selectedClient: IClient) => {
    setClient(selectedClient);
  };

  const handleAddService = () => {
    if (selectedService) {
      setServicesList([...servicesList, selectedService]);
      setPrice(price + selectedService.price);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!client) {
      return;
    }

    const serviceIDs = servicesList.map((service) => service.id);
    const upperLicensePlate = licensePlate.toUpperCase();

    const newBudget: ICreateBudgetRequest = {
      userID,
      clientID: client.id,
      clientName: client.name,
      clientPhone: client.phone,
      serviceIDs,
      vehicle,
      licensePlate: upperLicensePlate,
      price,
    };

    try {
      await budgetSchema.validate(newBudget);
    } catch (error: any) {
      alert(error.message);
      return;
    }

    try {
      const response = await api.post(`/budgets/`, newBudget, { withCredentials: true });

      if (response.status === 201) {
        goBack();
      } else {
        alert("Falha ao cadastrar orçamento");
      }

      setPrice(0);
      setClient(undefined);
      setSelectedService(undefined);
      setServicesList([]);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar orçamento</Card.Title>

        <Form onSubmit={handleSubmit}>
          <SelectInput label="Cliente" id="client" options={clients} onSelect={handleClientChange} required />
          <TextInput label="Veículo" id="vericle" value={vehicle} onChange={handleTextChange(setVehicle)} required />
          <TextInput label="Placa" id="licensePlate" value={licensePlate} onChange={handleTextChange(setLicensePlate)} required />
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

          <NumberInput label="Preço total" id="price" value={price} onChange={(e) => setPrice(Number(e.target.value))} required />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
