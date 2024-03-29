import { useContext } from "react";
import { Link } from "react-router-dom";
import { Navbar, Container, Button } from "react-bootstrap";
import { AuthContext } from "../contexts/AuthContext";

export default function TopNavbar() {
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
          <Button variant="outline-light" size="sm" onClick={handleLogout}>
            Sair
          </Button>
        )}
      </Container>
    </Navbar>
  );
}
