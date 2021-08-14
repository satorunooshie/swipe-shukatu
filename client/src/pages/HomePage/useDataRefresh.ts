import React, { useState, useEffect, createRef } from "react";
import { Ltd } from "../../type/Ltd";
import { Ltds } from "./DB";

export const useDataRefresh = (cardsLeft: Ltd[], ltdList: Ltd[]) => {
  const [childRefs, setChildRefs] = useState<React.RefObject<any>[]>(
    Array(ltdList.length)
      .fill(0)
      .map((i) => createRef())
  );
  const [page, setPage] = useState(0);
  const [newData, setNewData] = useState<Ltd[]>([]);

  useEffect(() => {
    if (cardsLeft.length > 5) return;

    // TODO: ここでデータフェッチ, newDataに割り当てる
    setNewData(Ltds[page]);
    if (page >= 2) return;
    setPage(page + 1);
    setChildRefs([
      ...childRefs,
      ...Array(newData.length)
        .fill(0)
        .map((i) => createRef()),
    ]);
  }, [cardsLeft]);
  console.log(childRefs)
  return [childRefs, newData] as const;
};
