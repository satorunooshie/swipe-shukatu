import { VFC } from "react";
import { Flex, IconButton, Input, Button } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { AttachmentIcon } from "@chakra-ui/icons";
import { FaTelegramPlane } from "react-icons/fa";

const ChatForm: VFC = () => {
  // TODO: formの処理
  return (
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
        ml="2"
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
  );
};

export default ChatForm;
