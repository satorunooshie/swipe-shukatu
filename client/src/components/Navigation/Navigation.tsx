import { VFC } from "react";
import {
  Flex,
  IconButton,
  Center,
  Icon,
  Divider,
  Wrap,
  Text,
} from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai";
import { HiChatAlt2 } from "react-icons/hi";
import { BsLightningFill } from "react-icons/bs";
import { NavLink, useLocation } from "react-router-dom";

const Navigation: VFC = () => {
  let location = useLocation();

  return (
    <Wrap pos="fixed" bottom="0" w="full">
      <Divider />
      <Flex px={10} pb={1} justify="space-around" w="full">
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
                color: "blue.100",
              }}
              _focus={{
                color: "blue.300",
              }}
              _active={{
                bg: "white",
                color: "blue.300",
              }}
            />
          </NavLink>
        </Center>
        <Center>
          <NavLink to="/recommend">
            <IconButton
              isActive={
                location.pathname === "/recommend" ||
                location.pathname === "/recommend/"
              }
              colorScheme="blackAlpha"
              aria-label="recommend"
              variant="ghost"
              size="lg"
              icon={<Icon as={BsLightningFill} w={7} h={7} />}
              _hover={{
                color: "blue.100",
              }}
              _focus={{
                color: "blue.300",
              }}
              _active={{
                bg: "white",
                color: "blue.300",
              }}
            />
          </NavLink>
        </Center>
        <Center>
          <NavLink to="/message">
            <IconButton
              isActive={
                location.pathname === "/message" ||
                location.pathname === "/message/"
              }
              colorScheme="blackAlpha"
              aria-label="message"
              variant="ghost"
              size="lg"
              icon={<Icon as={HiChatAlt2} w={7} h={7} />}
              _hover={{
                color: "blue.100",
              }}
              _focus={{
                color: "blue.300",
              }}
              _active={{
                color: "blue.300",
                bg: "white",
              }}
            />
          </NavLink>
        </Center>
      </Flex>
    </Wrap>
  );
};

export default Navigation;
