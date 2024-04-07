import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput, EmailInput, PhoneInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { ICustomer, IUpdateCustomerRequest } from "../../types/Customer";

export default function CustomerDetails() {
  const customerID = useParams().customerID;
  const [customer, setCustomer] = useState<ICustomer | null>(null);
  const [cpf, setCpf] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");

  const navigate = useNavigate();
  const goBack = () => navigate("/customers/all");

  const fetchCustomerData = async () => {
    try {
      const response = await api.get(`/customers/${customerID}`, { withCredentials: true });
      const customerData = response.data;

      setCustomer(customerData);
      setCpf(customerData.cpf);
      setName(customerData.name);
      setEmail(customerData.email);
      setPhone(customerData.phone);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchCustomerData();
  }, []);

  const handleInputChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setter(event.target.value);
  };

  const handleUpdateCustomer = async () => {
    if (customer) {
      const updatedCustomerData: IUpdateCustomerRequest = {
        id: Number(customerID),
        cpf: cpf === "" ? null : cpf,
        name: name === "" ? null : name,
        email: email === "" ? null : email,
        phone: phone === "" ? null : phone,
        userID: customer.userID,
      };

      try {
        const response = await api.patch("/customers/", updatedCustomerData, { withCredentials: true });
        setCustomer(response.data);
        if (response.status === 204) goBack();
      } catch (error: any) {
        alert(error.response.data.error);
      }
    }
  };

  const handleDeleteCustomer = async () => {
    if (customer && confirm("Tem certeza de que deseja excluir este cliente?")) {
      try {
        const response = await api.delete(`/customers/${customerID}`, { withCredentials: true });
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
        <span className="material-symbols-outlined" role="button" onClick={handleDeleteCustomer}>
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {customer ? (
          <>
            <Card.Title className="mb-3">Detalhes do cliente</Card.Title>
            <Form>
              <TextInput label="CPF" id="cpf" value={cpf} onChange={handleInputChange(setCpf)} />
              <TextInput label="Nome" id="name" value={name} onChange={handleInputChange(setName)} required />
              <EmailInput label="Email" id="email" value={email} onChange={handleInputChange(setEmail)} />
              <PhoneInput label="Telefone" id="phone" value={phone} onChange={handleInputChange(setPhone)} required />

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateCustomer}>
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
