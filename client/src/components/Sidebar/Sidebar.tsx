import { VFC } from "react";
import { Flex, Icon, VStack, Text, Button, Center } from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai";
import { HiChatAlt2 } from "react-icons/hi";
import { BsLightningFill } from "react-icons/bs";
import { NavLink, useLocation } from "react-router-dom";
import { MAIN_COLOR } from "../../constants/MainColor";

const Sidebar: VFC = () => {
  let location = useLocation();

  return (
    <VStack
      display={["none", "none", "block"]}
      pos="sticky"
      left="0"
      top="5"
      h="75vh"
      minW="250px"
      px="5"
      py="6"
      mr="4"
      marginTop="2.5vh"
      border="1px"
      borderColor="gray.100"
      borderRadius="15px"
    >
      <NavLink to="/">
        <Text
          h1
          fontSize="2xl"
          fontFamily={"heading"}
          color={"gray.800"}
          fontWeight="bold"
          mb="4"
          textAlign="center"
        >
          Swipe Shukatsu
        </Text>
      </NavLink>
      <NavLink to="/">
        <Flex
          align="center"
          borderRadius="8"
          p="3"
          mb="5"
          bg={location.pathname === "/" ? `${MAIN_COLOR}.100` : "transparent"}
          _hover={{ textDecor: "none", backgroundColor: `${MAIN_COLOR}.100` }}
        >
          <Icon
            as={AiFillFire}
            w={7}
            h={7}
            color={location.pathname === "/" ? `${MAIN_COLOR}.400` : "gray.500"}
          />
          <Text ml="4" color="gray.700" fontWeight="bold">
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
              ? `${MAIN_COLOR}.100`
              : "transparent"
          }
          _hover={{ textDecor: "none", backgroundColor: `${MAIN_COLOR}.100` }}
        >
          <Icon
            as={BsLightningFill}
            w={7}
            h={7}
            color={
              location.pathname === "/recommend" ||
              location.pathname === "/recommend/"
                ? `${MAIN_COLOR}.400`
                : "gray.500"
            }
          />
          <Text ml="4" color="gray.700" fontWeight="bold">
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
              ? `${MAIN_COLOR}.100`
              : "transparent"
          }
          _hover={{ textDecor: "none", backgroundColor: `${MAIN_COLOR}.100` }}
        >
          <Icon
            as={HiChatAlt2}
            w={7}
            h={7}
            color={
              location.pathname === "/message" ||
              location.pathname === "/message/"
                ? `${MAIN_COLOR}.400`
                : "gray.500"
            }
          />
          <Text ml="4" color="gray.700" fontWeight="bold">
            メッセージ
          </Text>
        </Flex>
      </NavLink>
      <Center>
        <Button
          fontSize={"lg"}
          size="lg"
          fontWeight={600}
          color={"white"}
          bg={`${MAIN_COLOR}.400`}
          w="100%"
          _hover={{
            bg: `${MAIN_COLOR}.300`,
          }}
        >
          Log In
        </Button>
      </Center>
    </VStack>
  );
};

export default Sidebar;
