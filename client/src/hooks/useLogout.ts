import { useContext } from "react";
import firebase from "../lib/firebase";
import { CurrentUserContext } from "../context/CurrentUserContext";
import { useShowToast } from "./useShowToast"

export const useLogout = () => {
  const { setCurrentUser } = useContext(CurrentUserContext);
  const showToast = useShowToast("ログアウトしました")
  return () => {
    firebase
      .auth()
      .signOut()
      .then(() => {
        setCurrentUser(null);
        showToast()
      });
  };
};
