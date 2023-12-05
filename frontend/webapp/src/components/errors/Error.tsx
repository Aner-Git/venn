type Props = {
  msg: string;
};

const Error = ({ msg }: Props) => {
  return <div color="red">{msg}</div>;
};

export default Error;
