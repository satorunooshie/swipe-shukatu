import { useContext } from "react";
import useSWR from "swr";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import axios from "../../lib/axios";

// TODO: Check
const fetcher = (url: string, token: string) =>
  axios
    .get(url, {
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
    })
    .then((res) => res.data);

export const useMessageList = () => {
  const { currentUser } = useContext(CurrentUserContext);
  const { data: messages, error } = useSWR(
    ["/message/list", currentUser?.token],
    fetcher
  );

  return { messages, error };
};
