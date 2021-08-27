import useSWR from "swr";

// TODO: API call
const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      Accept: "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => res.results);

export const useMessageList = () => {
  const { data: ltds, error } = useSWR(
    "https://icanhazdadjoke.com/search?page=2",
    fetcher
  );

  return { ltds, error };
};
