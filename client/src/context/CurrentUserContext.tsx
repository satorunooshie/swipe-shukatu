import firebase from "firebase";
import { FC, createContext, useEffect, useState } from "react";
import { useShowToast } from "../hooks/useShowToast";
import { User } from "../type/User";

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
  const showToast = useShowToast();

  useEffect(() => {
    const temp = localStorage.getItem('user')
    if(temp){
      const user = JSON.parse(temp)
      setCurrentUser(user);
      return
    }

    // マウント時にログインをチェック
    firebase.auth().onAuthStateChanged(async (user) => {
      if (user) {
        //@ts-ignore
        // console.log(user.Aa)
        showToast(`ログインしました`);
        //@ts-ignore
        setCurrentUser({ uid: user.uid, token: user.Aa });
        //@ts-ignore
        localStorage.setItem("user", JSON.stringify({ uid: user.uid, token: user.Aa }));
      } else {
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
