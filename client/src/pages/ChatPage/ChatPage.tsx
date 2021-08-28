import { VFC, useContext, useEffect } from "react";
import {
  Container,
  Stack,
  Flex,
  Avatar,
  Box,
  Text,
  Spinner,
  Center,
  useDisclosure,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import LtdDetailModal from "../../components/LtdDetailModal/LtdDetailModal";
import ChatForm from "../../components/ChatForm/ChatForm";
import GoBackBtn from "../../components/GoBackBtn/GoBackBtn";
import { useMessages } from "./useMessages";
import { Message } from "../../type/Message";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import { useHistory } from "react-router-dom";

const ChatPage: VFC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { messages, ltd, error } = useMessages();
  const { currentUser } = useContext(CurrentUserContext);
  const history = useHistory();

  // ログインしていなければ/にリダイレクト
  useEffect(() => {
    if (currentUser) return;
    history.push("/");
  }, [currentUser]);

  if (error) return <h1>An error has occurred.</h1>;
  if (!messages || !ltd)
    return (
      <Container w="full" minH="full">
        <Center w="full" h="full">
          <Spinner
            thickness="4px"
            speed="0.5s"
            emptyColor="gray.200"
            color="blue.300"
            size="xl"
          />
        </Center>
      </Container>
    );
  return (
    <Container w="full" minH="full">
      <LtdDetailModal ltd={ltd} isOpen={isOpen} onClose={onClose} />
      <Flex
        pos="fixed"
        top="5"
        zIndex="2"
        w="full"
        minH={"50px"}
        align={"center"}
        justify="space-between"
      >
        <GoBackBtn />
      </Flex>
      <Stack
        w="full"
        pt="50px"
        direction="column-reverse"
        minH="full"
        align="end"
        justifyContent="flex-end"
      >
        {/* 企業のメッセージ */}
        {/* <Flex align="center" w="full" mb="2">
          <Flex align="center" w="70%">
            <Avatar mr="4" onClick={() => onOpen()} />
            <Box bg="gray.200" borderRadius="xl" p="2" shadow="md">
              <Text color="gray.600" fontWeight="bold">
                {message.name}
              </Text>
            </Box>
          </Flex>
        </Flex> */}
        {/* 自身のメッセージ */}
        {messages.map((message: Message) => (
          <Flex align="center" w="full" justify="flex-end" mb="2">
            <Flex align="center" w="70%" direction="row-reverse" gap="2">
              <Avatar ml="4" />
              <Box bg={`${MAIN_COLOR}.400`} borderRadius="xl" p="2" shadow="md">
                <Text color="white" fontWeight="bold">
                  {message.content}
                </Text>
              </Box>
            </Flex>
          </Flex>
        ))}
      </Stack>
      {/* @ts-ignore */}
      <ChatForm currentUser={currentUser} recruit_id={ltd.recruit_id} />
    </Container>
  );
};

export default ChatPage;
