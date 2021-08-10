import { useContext } from "react";
import firebase from "../lib/firebase";
import { CurrentUserContext } from "../context/CurrentUserContext";

export const useLogout = () => {
  const { setCurrentUser } = useContext(CurrentUserContext);
  return () => {
    firebase
      .auth()
      .signOut()
      .then(() => {
        setCurrentUser(null);
        alert("logout");
      });
  };
};
