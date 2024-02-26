import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Card, Table } from "react-bootstrap";
import { SubmitButton } from "../../components/SubmitButton";

import api from "../../service/api";

import { IBudget } from "../../types/Budget";
import { IService } from "../../types/Service";
import { formatDate } from "../../utils/formatDate";
import { formatCurrency } from "../../utils/formatCurrency";

export default function BudgetDetails() {
  const budgetID = useParams().budgetID;
  const [budget, setBudget] = useState<IBudget | null>(null);
  const [budgetServices, setBudgetServices] = useState<IService[]>([]);

  const navigate = useNavigate();
  const goBack = () => navigate("/budgets/list");

  const fetchBudgetData = async () => {
    try {
      const response = await api.get(`/budgets/${budgetID}`, { withCredentials: true });
      const budgetData = response.data;
      setBudget(budgetData);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  const fetchBudgetServices = async () => {
    try {
      const response = await api.get(`/budgets/list/services/${budgetID}`, { withCredentials: true });
      const budgetServicesData = response.data;
      setBudgetServices(budgetServicesData);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchBudgetData();
    fetchBudgetServices();
  }, []);

  return (
    <Card className="mx-auto mt-4" style={{ width: "24rem" }}>
      <Card.Header className="d-flex justify-content-between align-items-center">
        <span className="material-symbols-outlined" role="button" onClick={goBack}>
          arrow_back
        </span>
        <span className="material-symbols-outlined" role="button">
          download
        </span>
        <span className="material-symbols-outlined" role="button">
          delete
        </span>
      </Card.Header>

      <Card.Body>
        {budget ? (
          <>
            <Card.Title className="mb-3">Orçamento de serviço</Card.Title>

            <p>
              <strong>Data:</strong> {formatDate(budget.date)} <br />
              <strong>Cliente:</strong> {budget.clientName} <br />
              <strong>Veículo:</strong> {budget.vehicle} <br />
              <strong>Placa:</strong> {budget.licensePlate} <br />
            </p>

            <Table className="mt-3" striped bordered hover>
              <thead>
                <tr>
                  <th>Serviço</th>
                  <th>Preço</th>
                </tr>
              </thead>
              <tbody>
                {budgetServices.map((service, index) => (
                  <tr key={index}>
                    <td>{service.name}</td>
                    <td>{formatCurrency(service.price)}</td>
                  </tr>
                ))}
              </tbody>
            </Table>

            <p>
              <strong>Preço final: {formatCurrency(budget.price)}</strong>
            </p>

            <SubmitButton text="Aprovar orçamento" />
          </>
        ) : (
          <p>Carregando...</p>
        )}
      </Card.Body>
    </Card>
  );
}
