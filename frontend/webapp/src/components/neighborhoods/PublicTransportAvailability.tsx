type Props = { value: string | null };

const PublicTransportAvailability = ({ value }: Props) => {
  if (value === null || value === undefined) {
    return <>-</>;
  }
  let clz = "";
  switch (value) {
    case "low":
    case "none":
      clz = "PublicTransportAvailabilityLowNone";
      break;
  }

  return <span className={`${clz}`}>{value.toLowerCase()}</span>;
};

export default PublicTransportAvailability;
