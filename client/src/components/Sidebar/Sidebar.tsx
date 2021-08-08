import { VFC } from "react";
import {
  Flex,
  IconButton,
  Center,
  Icon,
  Divider,
  Wrap,
  VStack,
  Text,
} from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai";
import { HiChatAlt2 } from "react-icons/hi";
import { BsLightningFill } from "react-icons/bs";
import { NavLink, useLocation } from "react-router-dom";

const Sidebar: VFC = () => {
  let location = useLocation();

  return (
    <VStack
      display={["none", "none", "block"]}
      pos="static"
      h="75vh"
      minW="250px"
      px="5"
      py="6"
      marginTop="2.5vh"
      boxShadow="0 4px 12px 0 rgba(0, 0, 0, 0.05)"
      borderRadius="15px"
    >
      <NavLink to="/">
        <Flex
          align="center"
          borderRadius="8"
          p="3"
          mb="5"
          bg={location.pathname === "/" ? "blue.100" : "transparent"}
          _hover={{ textDecor: "none", backgroundColor: "blue.100" }}
        >
          <Icon
            as={AiFillFire}
            w={7}
            h={7}
            color={location.pathname === "/" ? "blue.400" : "gray.500"}
          />
          <Text ml="4" color="gray.700">
            ホーム
          </Text>
        </Flex>
      </NavLink>
      <NavLink to="/recommend">
        <Flex
          align="center"
          borderRadius="8"
          p="3"
          mb="5"
          bg={
            location.pathname === "/recommend" ||
            location.pathname === "/recommend/"
              ? "blue.100"
              : "transparent"
          }
          _hover={{ textDecor: "none", backgroundColor: "blue.100" }}
        >
          <Icon
            as={BsLightningFill}
            w={7}
            h={7}
            color={
              location.pathname === "/recommend" ||
              location.pathname === "/recommend/"
                ? "blue.400"
                : "gray.500"
            }
          />
          <Text ml="4" color="gray.700">
            おすすめ企業
          </Text>
        </Flex>
      </NavLink>
      <NavLink to="/message">
        <Flex
          align="center"
          borderRadius="8"
          p="3"
          mb="5"
          bg={
            location.pathname === "/message" ||
            location.pathname === "/message/"
              ? "blue.100"
              : "transparent"
          }
          _hover={{ textDecor: "none", backgroundColor: "blue.100" }}
        >
          <Icon
            as={HiChatAlt2}
            w={7}
            h={7}
            color={
              location.pathname === "/message" ||
              location.pathname === "/message/"
                ? "blue.400"
                : "gray.500"
            }
          />
          <Text ml="4" color="gray.700">
            メッセージ
          </Text>
        </Flex>
      </NavLink>
    </VStack>
  );
};

export default Sidebar;
