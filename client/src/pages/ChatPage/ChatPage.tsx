import { VFC } from "react";
import {
  Container,
  Stack,
  Flex,
  Avatar,
  Box,
  Text,
  IconButton,
  Input,
  Button,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { ArrowBackIcon, AttachmentIcon } from "@chakra-ui/icons";
import { FaTelegramPlane } from "react-icons/fa";
import { useHistory,useParams } from "react-router-dom";
import useSWR from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      Accept: "application/json",
    },
  })
    .then((res) => res.json())

const ChatPage: VFC = () => {
  const history = useHistory();
  // @ts-ignore
  const { ltdId } = useParams();
  const { data: message, error } = useSWR(
    `https://icanhazdadjoke.com/j/${ltdId}`,
    fetcher
  );

  if (error) return <h1>An error has occurred.</h1>;
  if (!message) return <h1>Loading...</h1>;
  return (
    <Container w="full" minH="full">
      <Flex
        pos="fixed"
        top="5"
        zIndex="2"
        w="full"
        minH={"50px"}
        align={"center"}
        justify="space-between"
      >
        <IconButton
          aria-label="back to page"
          size="lg"
          icon={<ArrowBackIcon />}
          onClick={() => history.goBack()}
        />
      </Flex>
      <Stack
        w="full"
        pt="50px"
        direction="column-reverse"
        minH="full"
        align="end"
        justifyContent="flex-end"
      >
        <Flex align="center" w="full" mb="2">
          <Flex align="center" w="70%">
            <Avatar mr="4" />
            <Box bg="gray.200" borderRadius="xl" p="2" shadow="md">
              <Text color="gray.600" fontWeight="bold">
                {message.joke}
              </Text>
            </Box>
          </Flex>
        </Flex>
        <Flex align="center" w="full" justify="flex-end" mb="2">
          <Flex align="center" w="70%" direction="row-reverse" gap="2">
            <Avatar ml="4" />
            <Box bg={`${MAIN_COLOR}.400`} borderRadius="xl" p="2" shadow="md">
              <Text color="white" fontWeight="bold">
                {message.joke}
              </Text>
            </Box>
          </Flex>
        </Flex>
      </Stack>
      <Flex
        pos="sticky"
        bottom="0"
        zIndex="2"
        w="full"
        minH={"50px"}
        align={"center"}
        justify="space-between"
        bg="white"
        py="4"
        borderTop="1px solid"
        borderColor="gray.200"
      >
        <IconButton
          variant="ghost"
          aria-label="attach image"
          icon={<AttachmentIcon />}
        />
        <Input
          variant="filled"
          placeholder="テキストを入力してください"
          aria-label="input message"
        />
        <Button
          colorScheme={MAIN_COLOR}
          loading="false"
          rightIcon={<FaTelegramPlane />}
          aria-label="send message"
          ml="2"
        >
          送信
        </Button>
      </Flex>
    </Container>
  );
};

export default ChatPage;
