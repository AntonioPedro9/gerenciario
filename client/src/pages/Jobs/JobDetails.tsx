import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Form, Button } from "react-bootstrap";
import { TextInput, NumberInput } from "../../components/CustomInputs";

import api from "../../service/api";

import { IJob, IUpdateJobRequest } from "../../types/Job";

export default function JobDetails() {
  const jobID = useParams().jobID;
  const [job, setJob] = useState<IJob | null>(null);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [duration, setDuration] = useState<number | "">(0);
  const [price, setPrice] = useState<number | "">(0);

  const navigate = useNavigate();
  const goBack = () => navigate("/jobs/list");

  const fetchJobData = async () => {
    try {
      const response = await api.get(`/jobs/${jobID}`, { withCredentials: true });
      const jobData = response.data;

      setJob(jobData);
      setName(jobData.name);
      setDescription(jobData.description);
      setDuration(jobData.duration);
      setPrice(jobData.price);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchJobData();
  }, []);

  const handleTextChange = (setter: React.Dispatch<React.SetStateAction<string>>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value);
  };

  const handleNumberChange = (setter: React.Dispatch<React.SetStateAction<number | "">>) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setter(e.target.value === "" ? "" : Number(e.target.value));
  };

  const handleUpdateJob = async () => {
    if (job) {
      const updatedJobData: IUpdateJobRequest = {
        id: Number(jobID),
        name,
        description,
        duration: Number(duration),
        price: Number(price),
        userID: job.userID,
      };

      try {
        const response = await api.put("/jobs/", updatedJobData, { withCredentials: true });
        setJob(response.data);
        if (response.status === 200) goBack();
      } catch (error: any) {
        alert(error.response.data.error);
      }
    }
  };

  const handleDeleteJob = async () => {
    if (job && confirm("Tem certeza de que deseja excluir este serviço?")) {
      try {
        const response = await api.delete(`/jobs/${jobID}`, { withCredentials: true });
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
        <span className="material-symbols-outlined" role="button" onClick={handleDeleteJob}>
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {job ? (
          <>
            <Card.Title className="mb-3">Detalhes do serviço</Card.Title>
            <Form>
              <TextInput label="Nome" id="name" value={name} onChange={handleTextChange(setName)} required />
              <TextInput label="Descrição" id="description" value={description} onChange={handleTextChange(setDescription)} />
              <NumberInput label="Duração (horas)" id="duration" value={duration} onChange={handleNumberChange(setDuration)} />
              <NumberInput label="Preço" id="price" value={price} onChange={handleNumberChange(setPrice)} />

              <Button variant="dark" type="button" style={{ width: "100%" }} onClick={handleUpdateJob}>
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
