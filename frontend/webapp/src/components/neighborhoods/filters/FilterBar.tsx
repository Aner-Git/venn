import Stack from "react-bootstrap/Stack";
import type { PropsWithChildren } from "react";

type Props = {
  gap: number;
};

const FilterBar = ({ gap = 2, children }: PropsWithChildren<Props>) => {
  return (
    <Stack
      direction="horizontal"
      className="justify-content-start align-items-start"
      gap={gap}
    >
      {children}
    </Stack>
  );
};

export default FilterBar;
