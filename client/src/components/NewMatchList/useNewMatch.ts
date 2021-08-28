import { useContext } from "react";
import useSWR from "swr";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import axios from "../../lib/axios"

// TODO: Check
const fetcher = (url: string, uid: string) =>
  axios
    .get(url, {
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${uid}`
      },
    })
    .then((res) => res.data);

export const useNewMatch = () => {
  const { currentUser } = useContext(CurrentUserContext);
  const { data: matches, error } = useSWR(
    ["/match/list", currentUser?.uid],
    fetcher
  );

  return { matches, error };
};
