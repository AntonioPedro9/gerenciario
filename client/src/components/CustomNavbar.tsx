import { useContext } from "react";
import { Link } from "react-router-dom";
import { Navbar, Container, Nav, NavDropdown, Button } from "react-bootstrap";

import { AuthContext } from "../contexts/AuthContext";

function CustomNavbar() {
  const { isLoggedIn } = useContext(AuthContext);
  const { logout } = useContext(AuthContext);

  const handleLogout = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await logout();
    } catch (error: any) {
      alert("Erro ao sair da conta");
      console.error(error.message);
    }
  };

  return (
    <Navbar bg="dark" variant="dark">
      <Container>
        <Navbar.Brand>
          <Link to="/" className="navbar-brand" style={{ display: "flex", alignItems: "center" }}>
            <span className="material-symbols-outlined">edit_note</span>
            <span style={{ marginLeft: "4px" }}>GERENCIARIO</span>
          </Link>
        </Navbar.Brand>

        {isLoggedIn && (
          <>
            <Nav className="me-auto">
              <NavDropdown title="Clientes" menuVariant="dark">
                <NavDropdown.Item as={Link} to="/clients/create">
                  Cadastrar
                </NavDropdown.Item>
                <NavDropdown.Item as={Link} to="/clients/list">
                  Buscar
                </NavDropdown.Item>
              </NavDropdown>

              <NavDropdown title="Serviços" menuVariant="dark">
                <NavDropdown.Item as={Link} to="/services/create">
                  Cadastrar
                </NavDropdown.Item>
                <NavDropdown.Item as={Link} to="/services/list">
                  Buscar
                </NavDropdown.Item>
              </NavDropdown>

              <NavDropdown title="Orçamentos" menuVariant="dark">
                <NavDropdown.Item as={Link} to="/budgets/create">
                  Cadastrar
                </NavDropdown.Item>
                <NavDropdown.Item as={Link} to="/budgets/list">
                  Buscar
                </NavDropdown.Item>
              </NavDropdown>
            </Nav>

            <Button variant="outline-light" size="sm" onClick={handleLogout}>
              Sair
            </Button>
          </>
        )}
      </Container>
    </Navbar>
  );
}

export default CustomNavbar;
