import { useContext } from "react";
import useSWR from "swr";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import axios from "../../lib/axios";
import { useParams } from "react-router-dom"

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

// LtdIdからMessageを取得
// type Res = {
//   ltd: {
//     id: string;
//     name: string;
//     profile: string;
//     employee_number: number;
//     average_number: number;
//     industry_id: number;
//     created_at: string;
//     updated_at?: string;
//     deleted_at?: string;
//   },
//   messages: [
//     {
//       ltd_id: number;
//       content: string;
//       photo: string;
//       created_at: string;
//     }
//   ];
// };
export const useMessages = () => {
  // @ts-ignore
  const { recruit_id } = useParams();
  const { currentUser } = useContext(CurrentUserContext);
  const { data: res, error } = useSWR(
    [`/message/${recruit_id}`, currentUser?.token],
    fetcher
  );

  return { messages: res.messages, ltd: res.ltd, error };
};
