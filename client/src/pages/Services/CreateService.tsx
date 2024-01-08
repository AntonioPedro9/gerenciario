import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form } from "react-bootstrap";

import { TextInput, NumberInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";
import { serviceSchema } from "../../utils/validations";

import { ICreateServiceRequest } from "../../types/Service";

export default function CreateService() {
  const userID = getUserID() || "";
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [duration, setDuration] = useState<number | "">(0);
  const [price, setPrice] = useState<number | "">(0);

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const handleTextChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleNumberChange = (setter: React.Dispatch<React.SetStateAction<number | "">>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value === "" ? "" : Number(e.target.value));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const newService: ICreateServiceRequest = {
      name,
      description,
      duration: duration === "" ? 0 : duration,
      price: price === "" ? 0 : price,
      userID,
    };

    try {
      await serviceSchema.validate(newService);
    } catch (error: any) {
      alert(error.message);
      return;
    }

    try {
      const response = await api.post(`/services/`, newService, { withCredentials: true });

      if (response.status === 201) {
        goBack();
      } else {
        alert("Falha ao cadastrar serviço");
      }

      setName("");
      setDescription("");
      setDuration(0);
      setPrice(0);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar serviço</Card.Title>

        <Form onSubmit={handleSubmit}>
          <TextInput label="Nome" id="name" value={name} onChange={handleTextChange(setName)} required />
          <TextInput label="Descrição" id="description" value={description} onChange={handleTextChange(setDescription)} />
          <NumberInput label="Duração (horas)" id="duration" value={duration} onChange={handleNumberChange(setDuration)} />
          <NumberInput label="Preço" id="price" value={price} onChange={handleNumberChange(setPrice)} />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
