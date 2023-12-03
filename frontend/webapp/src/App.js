import { BrowserRouter } from "react-router-dom";
import { ReactQueryDevtools } from "react-query/devtools";
import { QueryClient, QueryClientProvider } from "react-query";
import Container from "react-bootstrap/Container";
import Header from "./components/layouts/Header";
import NeighborhoodPage from "./pages/Neighborhoods";

const queryClient = new QueryClient();

function App() {
  return (
    <Container fluid className="App">
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <ReactQueryDevtools initialIsOpen={false} />
          <Header />
          <NeighborhoodPage />
        </QueryClientProvider>
      </BrowserRouter>
    </Container>
  );
}

export default App;
