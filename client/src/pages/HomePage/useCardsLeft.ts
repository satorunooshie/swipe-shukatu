import { useState, useEffect } from "react";
import { Ltd } from "../../type/Ltd";

export const useCardsLeft = (alreadyRemoved: number[], ltdList: Ltd[]) => {
  const [cardsLeft, setCardsLeft] = useState<Ltd[]>(ltdList);

  useEffect(() => {
    // cardsLeftを算出
    setCardsLeft(ltdList.filter((ltd) => !alreadyRemoved.includes(ltd.id)));
  }, [alreadyRemoved]);
  console.log(cardsLeft)
  return cardsLeft;
};
