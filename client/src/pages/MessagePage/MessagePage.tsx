import { VFC, useContext, useEffect } from "react";
import { Container } from "@chakra-ui/react";
import NewMatchList from "../../components/NewMatchList/NewMatchList";
import MessageList from "../../components/MessageList/MessageList";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import { useHistory } from "react-router-dom";

const MessagePage: VFC = () => {
  const { currentUser } = useContext(CurrentUserContext);
  const history = useHistory();

  // ログインしていなければ/にリダイレクト
  useEffect(() => {
    if (currentUser) return;
    history.push("/");
  }, [currentUser]);

  return (
    <Container w="full">
      <NewMatchList />
      <MessageList />
    </Container>
  );
};

export default MessagePage;
