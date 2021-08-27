import { VFC } from "react";
import { Flex, Stack, Text, Avatar } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { NavLink } from "react-router-dom";
import LoadingMessageList from "./LoadingMessageList";
import { useMessageList } from "./useMessageList";
import moment from "moment";

type MessageRes = {
  readonly ltd_id: number;
  readonly recruit_id: number;
  readonly name: string;
  readonly image: string;
  readonly last_message: string;
  readonly created_at: string;
};

const MessageList: VFC = () => {
  const { messages, error } = useMessageList();

  if (error) return <h1>An error has occurred.</h1>;
  if (!messages) return <LoadingMessageList />;

  return (
    <Stack pb="5">
      <Text
        fontSize="2xl"
        fontFamily={"heading"}
        color={`${MAIN_COLOR}.500`}
        fontWeight="bold"
      >
        メッセージ
      </Text>
      <Stack spacing="24px">
        {messages.map((message: MessageRes, i: number) => (
          <NavLink to={`/message/${message.ltd_id}`} key={i}>
            <Flex align="center" justify="space-between">
              <Avatar size="md" src={message.image} />
              <Stack ml="4" flex="1">
                <Text
                  fontSize="xl"
                  fontFamily={"heading"}
                  color={"gray.600"}
                  fontWeight="bold"
                >
                  {message.name}社
                </Text>
                <Text color={"gray.500"}>{message.last_message}</Text>
              </Stack>
              <Flex align="center" color={"gray.500"} ml="2">
                <Text>{moment(message.created_at).format('LL')}</Text>
              </Flex>
            </Flex>
          </NavLink>
        ))}
      </Stack>
    </Stack>
  );
};

export default MessageList;
