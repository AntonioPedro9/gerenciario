import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form } from "react-bootstrap";

import { TextInput, EmailInput, PhoneInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";

import { ICreateCustomerRequest } from "../../types/Customer";
import { customerSchema } from "../../utils/validations";

export default function CreateCustomer() {
  const userID = getUserID() || "";
  const [cpf, setCpf] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");

  const navigate = useNavigate();
  const goBack = () => navigate("/customers/list");

  const handleInputChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const newCustomer: ICreateCustomerRequest = {
      cpf,
      name,
      email,
      phone,
      userID,
    };

    try {
      await customerSchema.validate(newCustomer);
    } catch (error: any) {
      alert(error.message);
      return;
    }

    try {
      const response = await api.post(`/customers/`, newCustomer, { withCredentials: true });

      if (response.status === 201) goBack();

      setCpf("");
      setName("");
      setEmail("");
      setPhone("");
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar cliente</Card.Title>

        <Form onSubmit={handleSubmit}>
          <TextInput label="CPF" id="cpf" value={cpf} onChange={handleInputChange(setCpf)} />
          <TextInput label="Nome" id="name" value={name} onChange={handleInputChange(setName)} required />
          <EmailInput label="Email" id="email" value={email} onChange={handleInputChange(setEmail)} />
          <PhoneInput label="Telefone" id="phone" value={phone} onChange={handleInputChange(setPhone)} required />
          <SubmitButton text="Cadastrar" />
        </Form>
      </Card.Body>
    </Card>
  );
}
