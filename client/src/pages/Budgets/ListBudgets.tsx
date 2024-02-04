import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Card, InputGroup, Form, Table } from "react-bootstrap";

import api from "../../service/api";
import getUserID from "../../utils/getUserID";

import { IBudget } from "../../types/Budget";
import { formatCurrency } from "../../utils/formatCurrency";
import { formatDate } from "../../utils/formatDate";

export default function ListBudgets() {
  const userID = getUserID() || "";
  const [budgets, setBudgets] = useState<IBudget[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const fetchBudgets = async () => {
    try {
      const response = await api.get(`/budgets/list/${userID}`, { withCredentials: true });
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
    const name = budget.clientName
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "");

    return name.includes(searchTerm.toLowerCase());
  });

  return (
    <Card className="mx-auto my-4" style={{ width: "24rem" }}>
      <Card.Body>
        <Card.Title className="mb-3">Buscar orçamento</Card.Title>

        <InputGroup className="mb-3">
          <Form.Control placeholder="Cliente" onChange={handleInputChange} />
          <InputGroup.Text>
            <span className="material-symbols-outlined">search</span>
          </InputGroup.Text>
        </InputGroup>

        {filteredBudgets.length > 0 ? (
          <>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>Data</th>
                  <th>Cliente</th>
                  <th>Preço</th>
                </tr>
              </thead>
              <tbody>
                {filteredBudgets.map((budget) => (
                  <tr key={budget.id}>
                    <td>{formatDate(budget.date)}</td>
                    <td>
                      <Link to={`/budgets/${budget.id}`}>{budget.clientName}</Link>
                    </td>
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
