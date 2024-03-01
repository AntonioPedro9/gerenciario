import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import formatCurrency from "../../utils/formatCurrency";

import { IJob } from "../../types/Job";

export default function ListJobs() {
  const [jobs, setJobs] = useState<IJob[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const navigate = useNavigate();

  const fetchJobs = async () => {
    try {
      const response = await api.get(`/jobs/list/`, { withCredentials: true });
      setJobs(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchJobs();
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const filteredJobs = jobs.filter((job) => {
    const name = job.name
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto m-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar serviço</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Nome do serviço" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredJobs.length > 0 ? (
          <>
            <Table style={{ marginBottom: "0" }} striped bordered hover>
              <thead>
                <tr>
                  <th>Nome</th>
                  <th>Preço</th>
                </tr>
              </thead>
              <tbody>
                {filteredJobs.map((job) => (
                  <tr key={job.id} onClick={() => navigate(`/jobs/${job.id}`)} style={{ cursor: "pointer" }}>
                    <td>{job.name}</td>
                    <td>{formatCurrency(job.price)}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </>
        ) : null}
      </Card.Body>
    </Card>
  );
}
