import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form } from "react-bootstrap";

import { TextInput, EmailInput, PhoneInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";
import { clientSchema } from "../../utils/validations";

import { ICreateClientRequest } from "../../types/Client";

export default function CreateClient() {
  const userID = getUserID() || "";
  const [form, setForm] = useState<ICreateClientRequest>({
    cpf: "",
    name: "",
    email: "",
    phone: "",
    userID: userID,
  });

  const navigate = useNavigate();
  const goBack = () => navigate("/clients/list");

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
      await clientSchema.validate(form);
    } catch (error: any) {
      alert(error.message);
      return;
    }

    try {
      const response = await api.post(`/clients/`, form, { withCredentials: true });

      if (response.status === 201) {
        alert("Cliente criado com sucesso");
        goBack();
      }

      setForm({
        cpf: "",
        name: "",
        email: "",
        phone: "",
        userID: userID,
      });
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "32rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar cliente</Card.Title>

        <Form onSubmit={handleSubmit}>
          <TextInput label="CPF" id="cpf" value={form.cpf} onChange={handleInputChange} required />
          <TextInput label="Nome" id="name" value={form.name} onChange={handleInputChange} required />
          <EmailInput label="Email" id="email" value={form.email} onChange={handleInputChange} />
          <PhoneInput label="Telefone" id="phone" value={form.phone} onChange={handleInputChange} required />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
