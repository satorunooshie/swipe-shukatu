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

export const useNewMatch = () => {
  const { data: ltds, error } = useSWR(
    "https://icanhazdadjoke.com/search",
    fetcher
  );

  return { ltds, error };
};
