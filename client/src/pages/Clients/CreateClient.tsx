import React, { useState } from "react";
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

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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
      await api.post(`/clients/`, form, { withCredentials: true });

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
          <TextInput label="CPF" id="cpf" value={form.cpf} onChange={handleChange} required />
          <TextInput label="Nome" id="name" value={form.name} onChange={handleChange} required />
          <EmailInput label="Email" id="email" value={form.email} onChange={handleChange} />
          <PhoneInput label="Telefone" id="phone" value={form.phone} onChange={handleChange} required />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
