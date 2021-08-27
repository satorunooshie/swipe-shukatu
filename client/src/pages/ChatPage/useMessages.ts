import { useContext } from "react";
import useSWR from "swr";
import LtdDetailModal from "../../components/LtdDetailModal/LtdDetailModal";
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
  const { ltdId } = useParams();
  const { currentUser } = useContext(CurrentUserContext);
  const { data: res, error } = useSWR(
    [`/message/${ltdId}`, currentUser?.uid],
    fetcher
  );

  return { messages: res.messages, ltd: res.ltd, error };
};
