import { VFC } from "react";
import { Text, Stack, Container } from "@chakra-ui/react";
import NewMatchList from "../../components/NewMatchList/NewMatchList";
import MessageList from "../../components/MessageList/MessageList";

const MessagePage: VFC = () => {
  return (
    <Container w="full">
      <NewMatchList />
      <MessageList />
    </Container>
  );
};

export default MessagePage;
