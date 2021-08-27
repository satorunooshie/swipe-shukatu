import { VFC } from "react";
import {
  Flex,
  IconButton,
  Center,
  Icon,
  Divider,
  Wrap,
} from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai";
import { HiChatAlt2 } from "react-icons/hi";
import { BsLightningFill } from "react-icons/bs";
import { NavLink, useLocation } from "react-router-dom";
import { MAIN_COLOR } from "../../constants/MainColor";

const BottomTabs: VFC = () => {
  let location = useLocation();

  return (
    <Wrap
      pos="fixed"
      bottom="0"
      w="full"
      display={{ md: "none" }}
      bg="white"
      zIndex="1"
    >
      <Divider />
      <Flex px={0} pb={1} justify="space-around" w="full">
        <Center>
          <NavLink to="/">
            <IconButton
              isActive={location.pathname === "/"}
              colorScheme="blackAlpha"
              aria-label="home"
              variant="ghost"
              size="lg"
              icon={<Icon as={AiFillFire} w={7} h={7} />}
              _hover={{
                color: `${MAIN_COLOR}.100`,
              }}
              _focus={{
                color: `${MAIN_COLOR}.300`,
              }}
              _active={{
                bg: "transparent",
                color: `${MAIN_COLOR}.300`,
              }}
            />
          </NavLink>
        </Center>
        <Center>
          <NavLink to="/recommend">
            <IconButton
              isActive={location.pathname.includes("recommend")}
              colorScheme="blackAlpha"
              aria-label="recommend"
              variant="ghost"
              size="lg"
              icon={<Icon as={BsLightningFill} w={7} h={7} />}
              _hover={{
                color: `${MAIN_COLOR}.100`,
              }}
              _focus={{
                color: `${MAIN_COLOR}.300`,
              }}
              _active={{
                bg: "transparent",
                color: `${MAIN_COLOR}.300`,
              }}
            />
          </NavLink>
        </Center>
        <Center>
          <NavLink to="/message">
            <IconButton
              isActive={location.pathname.includes("message")}
              colorScheme="blackAlpha"
              aria-label="message"
              variant="ghost"
              size="lg"
              icon={<Icon as={HiChatAlt2} w={7} h={7} />}
              _hover={{
                color: `${MAIN_COLOR}.100`,
              }}
              _focus={{
                color: `${MAIN_COLOR}.300`,
              }}
              _active={{
                bg: "transparent",
                color: `${MAIN_COLOR}.300`,
              }}
            />
          </NavLink>
        </Center>
      </Flex>
    </Wrap>
  );
};

export default BottomTabs;
