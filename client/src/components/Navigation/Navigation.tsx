import { VFC } from "react";
import {
  Flex,
  IconButton,
  Center,
  Icon,
  Divider,
  Wrap,
} from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai"
import { HiChatAlt2 } from "react-icons/hi"
import { BsLightningFill } from "react-icons/bs"

const Navigation: VFC = () => {
  return (
    <Wrap pos="fixed" bottom="0" w="full">
      <Divider />
      <Flex px={10} pb={1} justify="space-around" w="full">
        <Center>
          <IconButton
            colorScheme="blackAlpha"
            aria-label="home"
            variant="ghost"
            size="lg"
            icon={<Icon as={AiFillFire} w={7} h={7} />}
            _hover={{
              color: "blue.400"
            }}
            _focus={{
              color: "blue.400"
            }}
            _active={{
              bg: "white"
            }}
          />
        </Center>
        <Center>
          <IconButton
            colorScheme="blackAlpha"
            aria-label="recommend"
            variant="ghost"
            size="lg"
            icon={<Icon as={BsLightningFill} w={7} h={7} />}
            _hover={{
              color: "blue.400"
            }}
            _focus={{
              color: "blue.400"
            }}
            _active={{
              bg: "white"
            }}
          />
        </Center>
        <Center>
          <IconButton
            colorScheme="blackAlpha"
            aria-label="message"
            variant="ghost"
            size="lg"
            icon={<Icon as={HiChatAlt2} w={7} h={7} />}
            _hover={{
              color: "blue.400"
            }}
            _focus={{
              color: "blue.400"
            }}
            _active={{
              bg: "white"
            }}
          />
        </Center>
      </Flex>
    </Wrap>
  );
};

export default Navigation;
