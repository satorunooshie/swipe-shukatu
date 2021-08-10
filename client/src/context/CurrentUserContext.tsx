import firebase from "firebase";
import { FC, createContext, useEffect, useState } from "react";
import { useShowToast } from "../hooks/useShowToast"

// 仮
type User = {
  readonly uid: string;
};

const CurrentUserContext = createContext(
  {} as {
    currentUser: User | null | undefined;
    setCurrentUser: React.Dispatch<
      React.SetStateAction<User | null | undefined>
    >;
  }
);

const CurrentUserProvider: FC = ({ children }) => {
  const [currentUser, setCurrentUser] = useState<User | null | undefined>(
    undefined
  );
  const showToast = useShowToast()

  useEffect(() => {
    // マウント時にログインをチェック
    firebase.auth().onAuthStateChanged(async (user) => {
      if (user) {
        showToast(`uid: ${user.uid} ログインしました`)

        // TODO： API叩く処理
        setCurrentUser({ uid: user.uid });
      } else {
        console.log("ログインしてない");
        setCurrentUser(null);
      }
    });
  }, []);

  /* 下階層のコンポーネントをラップする */
  return (
    <CurrentUserContext.Provider
      value={{ currentUser: currentUser, setCurrentUser: setCurrentUser }}
    >
      {children}
    </CurrentUserContext.Provider>
  );
};

export { CurrentUserContext, CurrentUserProvider };
