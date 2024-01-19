import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput, EmailInput, PhoneInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { IClient, IUpdateClientRequest } from "../../types/Client";

export default function ClientDetails() {
  const userID = useParams().userID;
  const [client, setClient] = useState<IClient | null>(null);
  const [cpf, setCpf] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");

  const navigate = useNavigate();
  const goBack = () => navigate("/clients/list");

  const fetchClientData = async () => {
    try {
      const response = await api.get(`/clients/${userID}`, { withCredentials: true });
      const clientData = response.data;

      setClient(clientData);
      setCpf(clientData.cpf);
      setName(clientData.name);
      setEmail(clientData.email);
      setPhone(clientData.phone);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchClientData();
  }, []);

  const handleInputChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setter(event.target.value);
  };

  const handleUpdateClient = async () => {
    if (client) {
      const updatedClientData: IUpdateClientRequest = {
        id: Number(userID),
        cpf,
        name,
        email,
        phone,
        userID: client.userID,
      };

      try {
        const response = await api.put("/clients/", updatedClientData, { withCredentials: true });
        setClient(response.data);
        if (response.status === 200) goBack();
      } catch (error: any) {
        alert(error.response.data.error);
      }
    }
  };

  const handleDeleteClient = async () => {
    if (client && confirm("Tem certeza de que deseja excluir este cliente?")) {
      try {
        const response = await api.delete(`/clients/${userID}`, { withCredentials: true });
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
        <span className="material-symbols-outlined" role="button" onClick={handleDeleteClient}>
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {client ? (
          <>
            <Card.Title className="mb-3">Detalhes do cliente</Card.Title>
            <Form>
              <TextInput label="CPF" id="cpf" value={cpf} onChange={handleInputChange(setCpf)} />
              <TextInput label="Nome" id="name" value={name} onChange={handleInputChange(setName)} required />
              <EmailInput label="Email" id="email" value={email} onChange={handleInputChange(setEmail)} />
              <PhoneInput label="Telefone" id="phone" value={phone} onChange={handleInputChange(setPhone)} required />

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateClient}>
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
