type Props = { amount: number };

let US = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

const USDollars = ({ amount }: Props) => {
  return US.format(amount);
};

export default USDollars;
