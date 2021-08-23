import { useParams } from "react-router-dom";
import useSWR from "swr";

const fetcher = (url: string) =>
fetch(url, {
  headers: {
    Accept: "application/json",
  },
}).then((res) => res.json());

// LtdIdからMessageを取得
export const useMessages = () => {
  // @ts-ignore
  const { ltdId } = useParams();
  const { data: message, error } = useSWR(
    `https://icanhazdadjoke.com/j/${ltdId}`,
    fetcher
  );

  return {message, error}
};
