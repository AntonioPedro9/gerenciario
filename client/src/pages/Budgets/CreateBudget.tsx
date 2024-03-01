import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Form, Badge, Button } from "react-bootstrap";

import { NumberInput, SelectInput, TextInput } from "../../components/CustomInputs";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";

import { ICustomer } from "../../types/Customer";
import { IJob } from "../../types/Job";
import { ICreateBudgetRequest } from "../../types/Budget";
import { budgetSchema } from "../../utils/validations";

export default function CreateBudget() {
  const userID = getUserID() || "";
  const [price, setPrice] = useState<number>(0);
  const [customer, setCustomer] = useState<ICustomer>();
  const [vehicle, setVehicle] = useState("");
  const [licensePlate, setLicensePlate] = useState("");
  const [selectedJob, setSelectedJob] = useState<IJob>();
  const [jobsList, setJobsList] = useState<IJob[]>([]);
  const [customers, setCustomers] = useState<ICustomer[]>([]);
  const [jobs, setJobs] = useState<IJob[]>([]);

  const navigate = useNavigate();
  const goBack = () => navigate("/budgets/list");

  const fetchCustomers = async () => {
    try {
      const response = await api.get(`/customers/list/`, { withCredentials: true });
      setCustomers(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  const fetchJobs = async () => {
    try {
      const response = await api.get(`/jobs/list/`, { withCredentials: true });
      setJobs(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchCustomers();
    fetchJobs();
  }, []);

  const handleTextChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleCustomerChange = (selectedCustomer: ICustomer) => {
    setCustomer(selectedCustomer);
  };

  const handleAddJob = () => {
    if (selectedJob) {
      setJobsList([...jobsList, selectedJob]);
      setPrice(price + selectedJob.price);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!customer) {
      return;
    }

    const jobIDs = jobsList.map((job) => job.id);
    const upperLicensePlate = licensePlate.toUpperCase();

    const newBudget: ICreateBudgetRequest = {
      userID,
      customerID: customer.id,
      jobIDs,
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
      setCustomer(undefined);
      setSelectedJob(undefined);
      setJobsList([]);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Cadastrar orçamento</Card.Title>

        <Form onSubmit={handleSubmit}>
          <SelectInput label="Cliente" id="customer" options={customers} onSelect={handleCustomerChange} required />{" "}
          <TextInput label="Veículo" id="vericle" value={vehicle} onChange={handleTextChange(setVehicle)} required />
          <TextInput label="Placa" id="licensePlate" value={licensePlate} onChange={handleTextChange(setLicensePlate)} required />
          <SelectInput label="Serviços" id="job" options={jobs} onSelect={setSelectedJob} required />{" "}
          <Button className="mb-3" variant="light" size="sm" onClick={handleAddJob}>
            Adicionar serviço ao orçamento +
          </Button>
          {jobsList.length !== 0 ? (
            <div className="mb-3">
              {jobsList.map((job, index) => (
                <Badge key={index} bg="secondary">
                  {job.name}
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
