import { useContext } from "react";
import useSWR from "swr";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import axios from "../../lib/axios";

// TODO: Check
const fetcher = (url: string, uid: string) =>
  axios
    .get(url, {
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${uid}`,
      },
    })
    .then((res) => res.data.results);

export const useMessageList = () => {
  const { currentUser } = useContext(CurrentUserContext);
  const { data: messages, error } = useSWR(
    ["/message/list", currentUser?.uid],
    fetcher
  );

  return { messages, error };
};
