import firebase from "firebase";
import { FC, createContext, useEffect, useState } from "react";

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

  useEffect(() => {
    // マウント時にログインをチェック
    firebase.auth().onAuthStateChanged(async (user) => {
      if (user) {
        console.log(`${user.uid}がログイン中`);

        // TODO：　API叩く処理をかく
        setCurrentUser({ uid: user.uid });
      } else {
        console.log("ログインしてない");
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
