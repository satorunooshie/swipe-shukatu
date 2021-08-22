import { VFC } from "react";
import { Flex, Stack, Text, Avatar } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { NavLink } from "react-router-dom";
import useSWR from "swr";
import { Ltd } from "../../type/Ltd";
import LoadingMessageList from "./LoadingMessageList";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      Accept: "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => res.results);

const MessageList: VFC = () => {
  const { data: ltds, error } = useSWR(
    "https://icanhazdadjoke.com/search?page=2",
    fetcher
  );

  if (error) return <h1>An error has occurred.</h1>;
  if (!ltds) return <LoadingMessageList />;

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
        {ltds.map((ltd: Ltd, i: number) => (
          <NavLink to={`/message/${ltd.id}`} key={ltd.id}>
            <Flex align="center" justify="space-between">
              <Avatar
                size="md"
                src={`https://icanhazdadjoke.com/j/${ltd.id}.png`}
              />
              <Stack ml="4" flex="1">
                <Text
                  fontSize="xl"
                  fontFamily={"heading"}
                  color={"gray.600"}
                  fontWeight="bold"
                >
                  {ltd.joke.slice(0, 10)}社
                </Text>
                <Text color={"gray.500"}>{ltd.joke}</Text>
              </Stack>
              <Flex align="center" color={"gray.500"} ml="2">
                <Text>8月19日</Text>
              </Flex>
            </Flex>
          </NavLink>
        ))}
      </Stack>
    </Stack>
  );
};

export default MessageList;
