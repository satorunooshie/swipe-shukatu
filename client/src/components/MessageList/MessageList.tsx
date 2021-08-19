import { VFC } from "react";
import { Flex, Stack, Text, Avatar } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { NavLink } from "react-router-dom";

const MessageList: VFC = () => {
  return (
    <Stack pb="5">
      <Text
        fontSize="2xl"
        fontFamily={"heading"}
        color={`${MAIN_COLOR}.500`}
        fontWeight="bold"
        mb="4"
      >
         メッセージ
      </Text>
      <Stack spacing="24px">
        {["A社", "B社", "C社", "D社", "E社", "F社", "G社"].map((ltd, i) => (
          <NavLink to={`/message/${i}`}>
          <Flex align="center" justify="space-between">
            <Avatar size="lg" />
            <Stack ml="4" flex="1">
              <Text
                fontSize="xl"
                fontFamily={"heading"}
                color={"gray.600"}
                fontWeight="bold"
              >
                {ltd}
              </Text>
              <Text color={"gray.500"}>20日までにエントリー忘れない</Text>
            </Stack>
            <Flex align="center" color={"gray.500"}>
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
