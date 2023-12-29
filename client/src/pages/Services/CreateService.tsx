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
  const [form, setForm] = useState<ICreateServiceRequest>({
    name: "",
    description: "",
    duration: 0,
    price: 0,
    userID: userID,
  });

  const navigate = useNavigate();
  const goBack = () => navigate("/services/list");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm({
      ...form,
      [name]: value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await serviceSchema.validate(form);
    } catch (error: any) {
      alert(error.message);
      return;
    }

    try {
      const response = await api.post(`/services/`, form, { withCredentials: true });

      if (response.status === 201) {
        alert("Serviço criado com sucesso");
        goBack();
      }

      setForm({
        name: "",
        description: "",
        duration: 0,
        price: 0,
        userID: userID,
      });
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "32rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar serviço</Card.Title>

        <Form onSubmit={handleSubmit}>
          <TextInput label="Nome" id="name" value={form.name} onChange={handleInputChange} required />
          <TextInput label="Description" id="description" value={form.description} onChange={handleInputChange} />
          <NumberInput label="Duração (horas)" id="duration" value={form.duration} onChange={handleInputChange} />
          <NumberInput label="Preço" id="price" value={form.price} onChange={handleInputChange} />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
