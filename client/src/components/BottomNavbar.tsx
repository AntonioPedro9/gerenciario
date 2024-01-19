import { useContext } from "react";
import { Link } from "react-router-dom";
import { Navbar, Container, Nav, NavDropdown } from "react-bootstrap";

import { AuthContext } from "../contexts/AuthContext";

export default function BottomNavbar() {
  const { isLoggedIn } = useContext(AuthContext);

  return (
    isLoggedIn && (
      <Navbar bg="dark" variant="dark" style={{ position: "fixed", bottom: 0, width: "100%" }}>
        <Container>
          <Nav className="mx-auto justify-content-between w-100">
            <NavDropdown drop="up" title="Clientes" menuVariant="dark">
              <NavDropdown.Item as={Link} to="/clients/create">
                Cadastrar
              </NavDropdown.Item>
              <NavDropdown.Item as={Link} to="/clients/list">
                Buscar
              </NavDropdown.Item>
            </NavDropdown>

            <NavDropdown drop="up" title="Serviços" menuVariant="dark">
              <NavDropdown.Item as={Link} to="/services/create">
                Cadastrar
              </NavDropdown.Item>
              <NavDropdown.Item as={Link} to="/services/list">
                Buscar
              </NavDropdown.Item>
            </NavDropdown>

            <NavDropdown drop="up" title="Orçamentos" menuVariant="dark">
              <NavDropdown.Item as={Link} to="/budgets/create">
                Cadastrar
              </NavDropdown.Item>
              <NavDropdown.Item as={Link} to="/budgets/list">
                Buscar
              </NavDropdown.Item>
            </NavDropdown>
          </Nav>
        </Container>
      </Navbar>
    )
  );
}
