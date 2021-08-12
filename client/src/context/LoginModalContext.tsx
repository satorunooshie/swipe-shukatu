import { FC, createContext } from "react";
import { useDisclosure } from "@chakra-ui/react";

const LoginModalContext = createContext(
  {} as {
    isOpen: boolean;
    onOpen: () => void;
    onClose: () => void;
  }
);

const LoginModalProvider: FC = ({ children }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <LoginModalContext.Provider
      value={{ isOpen: isOpen, onOpen: onOpen, onClose: onClose }}
    >
      {children}
    </LoginModalContext.Provider>
  );
};

export { LoginModalContext, LoginModalProvider };
