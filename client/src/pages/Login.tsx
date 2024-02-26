import React, { useState, useContext } from "react";
import { Link } from "react-router-dom";
import { Card, Form } from "react-bootstrap";

import { EmailInput, PasswordInput } from "../components/CustomInputs";
import { SubmitButton } from "../components/SubmitButton";

import { AuthContext } from "../contexts/AuthContext";

import { ILoginUserRequest } from "../types/User";

export default function Login() {
  const { login } = useContext(AuthContext);
  const [form, setForm] = useState<ILoginUserRequest>({
    email: "",
    password: "",
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
      await login(form.email, form.password);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Login</Card.Title>
        <Form className="mb-3" onSubmit={handleSubmit}>
          <EmailInput label="Email" id="email" value={form.email} onChange={handleChange} />
          <PasswordInput label="Senha" id="password" value={form.password} onChange={handleChange} />
          <SubmitButton text="Entrar" />
        </Form>
        Não tem uma conta? <Link to="/register">Cadastre-se</Link>
      </Card.Body>
    </Card>
  );
}
