import { PropsWithChildren } from "react";
import Container from "react-bootstrap/Container";

type BodyProps = {};
export default function Body({ children }: PropsWithChildren<BodyProps>) {
  return <Container className="Content ">{children}</Container>;
}
