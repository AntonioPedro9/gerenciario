import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import formatCurrency from "../../utils/formatCurrency";
import formatDate from "../../utils/formatDate";

import { IListBudgets } from "../../types/Budget";

export default function ListBudgets() {
  const [budgets, setBudgets] = useState<IListBudgets[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const navigate = useNavigate();

  const fetchBudgets = async () => {
    try {
      const response = await api.get(`/budgets/list/`, { withCredentials: true });
      setBudgets(response.data);
    } catch (error: any) {
      console.error(error.response.data.error);
    }
  };

  useEffect(() => {
    fetchBudgets();
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const filteredBudgets = budgets.filter((budget) => {
    const name = budget.customerName
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto m-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar orçamento</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Nome do cliente" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredBudgets.length > 0 ? (
          <>
            <Table style={{ marginBottom: "0" }} striped bordered hover>
              <thead>
                <tr>
                  <th>Data</th>
                  <th>Cliente</th>
                  <th>Preço</th>
                </tr>
              </thead>
              <tbody>
                {filteredBudgets.map((budget) => (
                  <tr key={budget.id} onClick={() => navigate(`/budgets/${budget.id}`)} style={{ cursor: "pointer" }}>
                    <td>{formatDate(budget.budgetDate)}</td>
                    <td>{budget.customerName}</td>
                    <td>{formatCurrency(budget.price)}</td>
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
