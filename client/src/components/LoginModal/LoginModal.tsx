import { VFC, useContext, useState } from "react";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Button,
} from "@chakra-ui/react";
import { LoginModalContext } from "../../context/LoginModalContext";
import { FcGoogle } from "react-icons/fc";
import { LoginWithGoogle } from "./LoginWithGoogle";

const LoginModal: VFC = () => {
  const { isOpen, onClose } = useContext(LoginModalContext);
  const [loading, setLoading] = useState(false);

  return (
    <>
      <Modal isOpen={isOpen} onClose={onClose} closeOnEsc id="LOGIN">
        <ModalOverlay />
        <ModalContent>
          <ModalHeader fontSize="2xl">Log In</ModalHeader>
          <ModalCloseButton />
          <ModalBody>現在、Googleアカウントでのみログイン可能です。</ModalBody>

          <ModalFooter>
            <Button
              isLoading={loading}
              loadingText="Redirecting"
              size="lg"
              colorScheme="gray"
              leftIcon={<FcGoogle />}
              onClick={() => {
                setLoading(true);
                LoginWithGoogle();
              }}
              mr="4"
            >
              Login with Google
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default LoginModal;
