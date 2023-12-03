import Navbar from "react-bootstrap/Navbar";
import Container from "react-bootstrap/Container";

function Header() {
  return (
    <Navbar bg="dark" sticky="top">
      <Container>
        <Navbar.Brand style={{ color: "white" }}>
          Z Venn way - a resident journey
        </Navbar.Brand>
      </Container>
    </Navbar>
  );
}

export default Header;
