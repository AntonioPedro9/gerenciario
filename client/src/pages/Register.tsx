import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { Card, Form } from "react-bootstrap";

import { TextInput, EmailInput, PasswordInput } from "../components/CustomInputs";
import { SubmitButton } from "../components/SubmitButton";

import api from "../service/api";

import { ICreateUserRequest } from "../types/User";

export default function Register() {
  const navigate = useNavigate();
  const [confirmPass, setConfirmPass] = useState("");
  const [form, setForm] = useState<ICreateUserRequest>({
    name: "",
    email: "",
    password: "",
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm({
      ...form,
      [name]: value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (form.password !== confirmPass) {
      alert("As senhas são diferentes.");
      return;
    }

    try {
      const response = await api.post("/users/", form);
      if (response.status === 201) navigate("/login");
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Criar conta</Card.Title>
        <Form onSubmit={handleSubmit}>
          <TextInput label="Nome" id="name" value={form.name} onChange={handleInputChange} />
          <EmailInput label="Email" id="email" value={form.email} onChange={handleInputChange} />
          <PasswordInput label="Senha" id="password" value={form.password} onChange={handleInputChange} />
          <PasswordInput label="Confirmar senha" id="confirmPass" value={confirmPass} onChange={(e) => setConfirmPass(e.target.value)} />
          <SubmitButton text="Cadastrar" />
        </Form>
        Já tem uma conta? <Link to="/login">Faça login</Link>
      </Card.Body>
    </Card>
  );
}
