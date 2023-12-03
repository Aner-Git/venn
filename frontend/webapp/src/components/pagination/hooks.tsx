import { useState } from "react";

const useOnPage = (
  sartPage: number
): [number, () => void, () => void, () => void] => {
  const [page, setCurrentPage] = useState(sartPage);

  const handlePrev = () => {
    setCurrentPage((prev) => {
      if (prev === 1) {
        return 1;
      }
      return --prev;
    });
  };

  const handleNext = () => {
    setCurrentPage((prev) => {
      return ++prev;
    });
  };
  const reset = () => setCurrentPage(1);

  return [page, handlePrev, handleNext, reset];
};

export default useOnPage;
